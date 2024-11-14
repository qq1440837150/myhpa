package main

import (
	"context"
	"gitlab.zhf.cn/myhpa_controller/pkg/apis/mygroup.example.com/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type Func func(context.Context, reconcile.Request) (reconcile.Result, error)

func (r Func) Reconcile(ctx context.Context, o reconcile.Request) (reconcile.Result, error) {
	return r(ctx, o)
}
func main() {
	println("sss")
	scheme := runtime.NewScheme()
	err := v1alpha1.AddToScheme(scheme)
	if err != nil {
		return
	}
	mgr, err := manager.New(
		config.GetConfigOrDie(),
		manager.Options{
			Scheme: scheme,
		},
	)
	if err != nil {
		log.Fatalf("unable to set up overall controller manager: %v", err)
	}
	var myFunc Func = func(c context.Context, r reconcile.Request) (reconcile.Result, error) {
		println("接受deployment")

		return reconcile.Result{}, nil
	}
	hpaController, err := controller.New(
		"my-operator", mgr,
		controller.Options{
			Reconciler: myFunc,
		})
	if err != nil {
		log.Fatalf("unable to set up individual controller: %v", err)
	}
	// 监控 Deployment 资源
	err = hpaController.Watch(
		&source.Kind{
			Type: &v1alpha1.MyResource{},
		},
		&handler.EnqueueRequestForObject{},
	)
	err = mgr.Start(context.Background())
	if err != nil {
		log.Fatalf("unable to start manager: %v", err)
	}
}
