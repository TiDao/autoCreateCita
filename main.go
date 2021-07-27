package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"encoding/json"
	"bytes"
)

func showJson(v interface{}) {
	data,_:= json.Marshal(v)
	var out bytes.Buffer
	json.Indent(&out,data,"","\t")
	fmt.Printf("%v\n",out.String())
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	_, err = clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(pods.Items[0])


	deployment,err := clientset.AppsV1().Deployments("default").Get(context.TODO(),"cita-master-sm2",metav1.GetOptions{})

	fmt.Println(deployment.TypeMeta)
	showJson(deployment.Spec.Template)
}
