package main

import (
	"demo1/pkg/informer"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
)

func main() {
	stopChan := make(chan struct{})
	err := informer.NewSharedInformerFactory(stopChan)
	if err != nil {
		panic(err.Error())
	}
	// 由informer负责 list and watch， 这边使用list方法
	items, err := informer.GetSharedInformerFactory().Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range items {
		fmt.Println(pod.Name, " ", pod.Namespace)
	}
}
