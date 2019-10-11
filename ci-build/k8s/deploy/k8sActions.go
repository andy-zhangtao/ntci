package deploy

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"ntci/ci-build/k8s/dataBus"
)

/*
Deploy Single Job In Kubernetes Cluster
*/

var kc *k8sClient

type k8sClient struct {
	client *kubernetes.Clientset
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

	return
}
