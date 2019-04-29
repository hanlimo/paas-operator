package main

import (
	"flag"
	"time"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	"github.com/farmer-hutao/k6s/pkg/controller/signals"

	clientset "github.com/farmer-hutao/k6s/pkg/controller/client/clientset/versioned"
	informers "github.com/farmer-hutao/k6s/pkg/controller/client/informers/externalversions"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	blueprintClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building blueprint clientset: %s", err.Error())
	}

	blueprintInformerFactory := informers.NewSharedInformerFactory(blueprintClient, time.Second*30)
	controller := NewController(blueprintClient, blueprintInformerFactory.Blueprintcontroller().V1alpha1().Databases())
	blueprintInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}