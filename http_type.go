package main

import(
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"context"
)


type RequestType struct{
	ChainName string `json: "chainName,omitempty"`
	ServicePort int32 `json: "servicePort,omitempty"`
	StorageSize string `json: "storageSize,omitempty"`
	ChainType string `json: "chainType,omitempty"`
}

type ResponseType struct{
	URL string `json: "URL,omiempty"`
	ChainList []string `json: "ChainList,omitempty"`

}


func CreateChain(w http.ResponseWriter,r *http.Request) {
	request := &RequestType{}
	citaChain := &CitaChain{}
	err := json.Unmarshal(r.Body,request)
	if err != nil{
		fmt.Sprintf(w,err.Error())
	}

	if err := citaChain.InitChain(request);err != nil{
		fmt.Sprintf(w,err.Error())
	}

	clientset := InitClientset()
	citaChain.CreateChain(clientset)

	time.Sleep(2 * time.Second)

	service,err := citaChain.GetService(clientset)
	if err != nil{
		log.Println(err)
		fmt.Sprintf(w,err.Error())
	}
}


