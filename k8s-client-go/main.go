package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var config *rest.Config
	var err error

	// Out-of-cluster configuration
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		panic("KUBECONFIG environment variable is not set")
	}
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	fmt.Printf("NAME\n")
	for _, pod := range pods.Items {
		fmt.Printf("%s\n",pod.Name)
	}	
}