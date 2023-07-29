package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Clients struct {
	clientSet kubernetes.Interface
}

func NewClients() (clients Clients) {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/max.local/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	clients.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
		return Clients{}
	}
	return clients
}

func (c *Clients) ClientSet() kubernetes.Interface {
	return c.clientSet
}
