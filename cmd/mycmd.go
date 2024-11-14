package main

import (
	"fmt"
	"gitlab.zhf.cn/myhpa_controller/pkg/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err :=
		clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			nil,
		).ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := clientset.NewForConfig(config)

	if err != nil {
		panic(err)
	}
	list, err := clientset.MygroupV1alpha1()
	if err != nil {
		panic(err)
	}
	for _, res := range list.Items {
		fmt.Printf("%s\n", res.GetName())
	}
}
