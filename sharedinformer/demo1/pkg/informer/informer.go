package informer

import (
	"demo1/pkg/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"time"
)

// 声明本地局部变量
var shareInformerFactory informers.SharedInformerFactory

func NewSharedInformerFactory(stopChan chan struct{}) (err error) {
	// 加载客户端
	var clients client.Clients
	clients = client.NewClients()

	// 实例化sharedinformerfactory
	shareInformerFactory = informers.NewSharedInformerFactory(clients.ClientSet(), time.Second*60)
	// gvrs  group version resources
	gvrs := []schema.GroupVersionResource{
		{Group: "", Resource: "pods", Version: "v1"},
		{Group: "", Resource: "services", Version: "v1"},
		{Group: "", Resource: "namespaces", Version: "v1"},
		{Group: "apps", Resource: "deployments", Version: "v1"},
		{Group: "apps", Resource: "statefulsets", Version: "v1"},
		{Group: "apps", Resource: "daemonsets", Version: "v1"},
	}
	for _, v := range gvrs {
		_, err := shareInformerFactory.ForResource(v)
		if err != nil {
			return err
		}
	}
	// start stf
	shareInformerFactory.Start(stopChan)
	shareInformerFactory.WaitForCacheSync(stopChan)
	return nil
}

func GetSharedInformerFactory() informers.SharedInformerFactory {
	return shareInformerFactory
}
