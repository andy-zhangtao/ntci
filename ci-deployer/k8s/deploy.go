package main

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func restartService(name, namespace string) (err error) {

	logrus.Infof("name: %s namespace: %s", name, namespace)
	r, err := client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	oldNum := r.Spec.Replicas

	logrus.Infof("old replicas: %d", *oldNum)
	zero := int32(0)
	r.Spec.Replicas = &zero
	_, err = client.AppsV1().Deployments(namespace).Update(r)
	if err != nil {
		return
	}

	retry := 0
	for {
		if retry == 60 {
			return errors.New("Operation Timeout! ")
		}

		time.Sleep(1 * time.Second)
		r, err = client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return
		}

		if r.Status.Replicas == 0 {
			r.Spec.Replicas = oldNum
			_, err = client.AppsV1().Deployments(namespace).Update(r)
			return err
		}

		retry++
	}

	return
}
