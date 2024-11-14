package main

import (
	"context"
	"fmt"
	"gitlab.zhf.cn/myhpa_controller/pkg/apis/mygroup.example.com/v1alpha1"
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
	var myresource v1alpha1.MyResource = v1alpha1.MyResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sss1",
			Namespace: "default",
		},
		Spec: v1alpha1.MyResourceSpec{
			Image: "sss",
		},
	}
	_, err = clientset.MygroupV1alpha1().MyResources("default").Create(context.TODO(),
		&myresource, metav1.CreateOptions{})
	if err != nil {
		println("ssad")
		panic(err)
	}
}
