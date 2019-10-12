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

	job := v1.Job{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%d", b.Name, b.Id),
			Namespace: kc.namespace,
		},
		Spec: v1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%d", b.Name, b.Id),
					Namespace: kc.namespace,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  b.Name,
							Image: "nginx",
						},
					},
				},
			},
		},
	}

	j, err := kc.client.BatchV1().Jobs(kc.namespace).Create(&job)
	if err != nil {
		return err
	}

	logrus.Debugf("Job: %s", j.String())
	//job := &batchv1.Job{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      "demo-job",
	//		Namespace: "gitlab",
	//	},
	//	Spec: batchv1.JobSpec{
	//		Template: apiv1.PodTemplateSpec{
	//			Spec: apiv1.PodSpec{
	//				Containers: []apiv1.Container{
	//					{
	//						Name:  "demo",
	//						Image: "myimage",
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	return nil
}
