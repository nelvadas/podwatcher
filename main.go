package main

import (
	"os"
	"path/filepath"
	"flag"
	"log"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	 "k8s.io/apimachinery/pkg/watch"

)

func main() {
	var _namespace,
	 	_labelSelector,
		_fieldSelector string

	//Define kubeconfig file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")


	//Load kubernetes config
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	//Create a client set to use k8s apis
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	//create an api object
	api := clientset.CoreV1()

	// Read namesapce from command line
	flag.StringVar(&_namespace, "n", "default", " namespace")
	flag.StringVar(&_labelSelector, "l", "", "Label selector ")
	flag.StringVar(&_fieldSelector, "f", "", "Field selector ")
	flag.Parse()

	opts := metav1.ListOptions{
		LabelSelector: _labelSelector,
		FieldSelector: _fieldSelector,
	}

	//Create a watcher on pods
	fmt.Printf("Starting a Pod watcher in namespace [%s]\n",_namespace)

	podWatcher, err := api.Pods(_namespace).Watch(opts)
	if err != nil {
		log.Fatal(err)
	}


    //Watch loop 

	podChannel := podWatcher.ResultChan()

	for event := range podChannel {
		pod,ok  :=event.Object.(*v1.Pod)
		if (! ok ) { log. Fatal(err )
		}

		switch event.Type{

		case watch.Added:
						log.Printf(" Pod %s added \n",pod.Name)

		case watch.Deleted:
						log.Printf(" Pod %s deleted \n",pod.Name)

		}
	}



}
