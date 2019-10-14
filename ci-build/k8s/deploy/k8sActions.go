package deploy

import (
	"fmt"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"ntci/ci-build/k8s/dataBus"
	"ntci/ci-build/k8s/store"
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

func NewJob(b store.Build) (err error) {

	// Clear build job after 10mins.
	ttl := int32(60 * 10)

	job := v1.Job{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%d", b.Name, b.Id),
			Namespace: kc.namespace,
		},
		Spec: v1.JobSpec{
			TTLSecondsAfterFinished: &ttl,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%d", b.Name, b.Id),
					Namespace: kc.namespace,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  b.Name,
							Image: b.Image,
							Env: []apiv1.EnvVar{
								{
									Name:  "NTCI_BUILDER_JID",
									Value: b.Name,
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
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
		},
	}

	j, err := kc.client.BatchV1().Jobs(kc.namespace).Create(&job)
	if err != nil {
		return err
	}

	logrus.Debugf("Job Pod Num: %d", j.Status.Succeeded)

	return nil
}
