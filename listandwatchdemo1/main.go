package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/max.local/.kube/config")
	clientSet, _ := kubernetes.NewForConfig(config)
	// list all ns
	list, _ := clientSet.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})

	for _, item := range list.Items {
		log.Info(item.Name)
		podList, _ := clientSet.CoreV1().Pods(item.Name).List(context.Background(), v1.ListOptions{})
		for _, pod := range podList.Items {
			log.Info(pod.Name)
		}
	}
	// list all pod

}
