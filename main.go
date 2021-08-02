package main

import (
	//"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"encoding/json"
	"bytes"
	"log"
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

var citaChain = &CitaChain{}
var request = &RequestType{
	ChainName: "test-chain-name",
	ServicePort: 1010,
	StorageSize: "10Gi",
	ChainType: "secp256",
}

func InitClientset() *kubernetes.Clientset{
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Println(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err.Error())
	}

	return clientset
}

func main() {

	if err := citaChain.InitChain(request); err != nil{
		log.Println(err)
	}

	//citaChain.CreateChain(clientset)
	citaChain.DeleteCitaChain(clientset)
}
