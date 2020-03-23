package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	podName := "example-memcached-7c4df9b7b4-h9fmd"
	for {
		pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		_, err = clientSet.CoreV1().Pods("default").Get(context.TODO(), podName, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Sprintf("Pod %s not found in default namespace\n", podName)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Sprintf("Found %s pod in default namespace\n", podName)
		}

		time.Sleep(10 * time.Second)
	}

}
