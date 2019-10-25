package deploy

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"ntci/ci-build/k8s/dataBus"
	"ntci/ci-build/k8s/store"
	build_rpc_v1 "ntci/ci-grpc/build"
)

/*
Deploy Single Job In Kubernetes Cluster
*/

var kc *k8sClient

type k8sClient struct {
	client    *kubernetes.Clientset
	namespace string
}

func InitK8sClient(bus *dataBus.DataBus) (err error) {

	config, err := clientcmd.BuildConfigFromFlags("", bus.K8S.Config)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	kc = new(k8sClient)

	kc.client = clientset
	kc.namespace = bus.K8S.Namespace

	info, err := kc.client.ServerVersion()
	if err != nil {
		return err
	}

	logrus.Infof("Kubernetes Version: %s", info.String())
	return
}

func DeleteJob(b store.Build) (err error) {
	job := fmt.Sprintf("%s-%d", b.Name, b.Id)
	logrus.Infof("Remove Job: %s", job)

	return kc.client.BatchV1().Jobs(kc.namespace).Delete(job, &metav1.DeleteOptions{})
}

// NewJob
// commenv is the common environment. Every user will use it.
func NewJob(b store.Build, commenv map[string]string) (err error) {

	// Clear build job after 10mins.
	// Try clean job before create

	err = DeleteJob(b)
	if err != nil {
		logrus.Errorf("Delete Error: %s . Maybe is normal", err.Error())
	}

	ttl := int32(60 * 10)
	bf := int32(1)

	var ev []apiv1.EnvVar

	for key, value := range commenv {
		ev = append(ev, apiv1.EnvVar{
			Name:  key,
			Value: value,
		})
	}

	job := v1.Job{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%d", b.Name, b.Id),
			Namespace: kc.namespace,
		},
		Spec: v1.JobSpec{
			BackoffLimit:            &bf,
			TTLSecondsAfterFinished: &ttl,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%d", b.Name, b.Id),
					Namespace: kc.namespace,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            b.Name,
							Image:           b.Image,
							ImagePullPolicy: apiv1.PullAlways,
							Env: []apiv1.EnvVar{
								{
									Name:  "NTCI_BUILDER_SHA",
									Value: b.Sha,
								},
								{
									Name:  "NTCI_BUILDER_USER",
									Value: b.User,
								},
								{
									Name:  "NTCI_BUILDER_JID",
									Value: b.Name,
								},
								{
									Name:  "NTCI_BUILDER_ID",
									Value: strconv.Itoa(b.Id),
								},
								{
									Name:  "NTCI_BUILDER_GIT",
									Value: b.Git,
								},
								{
									Name:  "NTCI_BUILDER_BRANCH",
									Value: b.Branch,
								},
								{
									Name:  "NTCI_BUILDER_TOKEN",
									Value: b.Token,
								},
								{
									Name:  "NTCI_BUILDER_ADDR",
									Value: b.Addr,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "dockersock",
									MountPath: "/var/run/docker.sock",
								},
							},
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
					Volumes: []apiv1.Volume{
						{
							Name: "dockersock",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									Path: "/var/run/docker.sock",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, e := range ev {
		job.Spec.Template.Spec.Containers[0].Env = append(job.Spec.Template.Spec.Containers[0].Env, e)
	}

	j, err := kc.client.BatchV1().Jobs(kc.namespace).Create(&job)
	if err != nil {
		return err
	}

	logrus.Debugf("Job Pod Num: %d", j.Status.Succeeded)

	return nil
}

func GetJobLog(jobname string, flowing bool, ls build_rpc_v1.BuildService_GetJobLogServer) (err error) {
	pod, err := getPodOfJob(jobname)
	if err != nil {
		return
	}

	logrus.Debugf("Find pod %s of job %s", pod, jobname)

	line := int64(1000)
	req := kc.client.CoreV1().Pods(kc.namespace).GetLogs(pod, &apiv1.PodLogOptions{
		TailLines: &line,
		Follow:    flowing,
	})

	podLogs, err := req.Stream()
	if err != nil {
		return errors.New("error in opening stream")
	}

	defer podLogs.Close()

	for {
		data := make([]byte, 1024)
		n, err := podLogs.Read(data)
		logrus.Errorf("%d, err: %v", n, err)
		if err != nil {
			fmt.Print(string(data[:n]))
			return ls.Send(&build_rpc_v1.Log{
				Message: string(data[:n]),
			})
		}

		fmt.Print(string(data[:n]))
		if ls.Send(&build_rpc_v1.Log{
			Message: string(data[:n]),
		}) != nil {
			break
		}
	}

	return
}

func getPodOfJob(jobname string) (podname string, err error) {
	selector := fmt.Sprintf("job-name=%s", jobname)

	logrus.Debugf("Select Pod via %s", selector)

	p, err := kc.client.CoreV1().Pods(kc.namespace).List(metav1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		return
	}

	if len(p.Items) == 0 {
		err = errors.New("Can not find the pod of this job. ")
		return
	}

	return p.Items[0].Name, nil
}
