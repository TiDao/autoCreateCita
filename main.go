package main

import (
	//"context"
	//"flag"
	"fmt"
	//"os"
	//"path/filepath"
	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/rest"
	"encoding/json"
	"bytes"
	"log"
	"net/http"
)

func showJson(v interface{}) {
	data,_:= json.Marshal(v)
	var out bytes.Buffer
	json.Indent(&out,data,"","\t")
	fmt.Printf("%v\n",out.String())
}

var templateDir = "./template"
//var citaChain = &CitaChain{}
//var request = &RequestType{
//	ChainName: "test-chain-name",
//	ServicePort: 1010,
//	StorageSize: "10Gi",
//	ChainType: "secp256",
//}

func InitClientset() *kubernetes.Clientset{

	//config,err := rest.InClusterConfig()
	config, err := clientcmd.BuildConfigFromFlags("", "/home/heyue/kube/config")
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

	http.HandleFunc("/create",HttpCreateChain)
	http.HandleFunc("/delete",HttpDeleteChain)
	http.HandleFunc("/",HttpListChain)
	err := http.ListenAndServe("0.0.0.0:10000",nil)
	if err != nil{
		log.Println(err)
	}
}
