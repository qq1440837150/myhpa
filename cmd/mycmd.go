package main

import (
	"context"
	"fmt"
	"gitlab.zhf.cn/myhpa_controller/pkg/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	list, err := clientset.MygroupV1alpha1().
		MyResources("default").
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, res := range list.Items {
		fmt.Printf("%s\n", res.GetName())
	}
}
