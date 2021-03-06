package main

import(
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
	"log"
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


func HttpCreateChain(w http.ResponseWriter,r *http.Request) {
	defer func(){
		if r := recover(); r!=nil{
			log.Println("recover...",r)
		}
	}()

	request := &RequestType{}
	citaChain := &CitaChain{}
	response := &ResponseType{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil{
		log.Println(err)
		fmt.Fprintf(w,err.Error())
	}

	if err := citaChain.InitChain(request);err != nil{
		fmt.Fprintf(w,err.Error())
	}

	clientset := InitClientset()
	log.Printf("start init k8s client\n")

	//check chainName has same one in the cluster
	chains,_:= ListChain(clientset)
	if len(chains) > 0{
		for _,chain := range chains{
			if chain == request.ChainName{
				fmt.Fprintf(w,"the chainName %s had same one in the cluster\n",chain)
				log.Printf("%s has be used in cluster\n",chain)
				panic(err)
			}
		}
	}


	if err := citaChain.CreateChain(clientset);err != nil{
		log.Println(err)
	}else{
		log.Printf("start create %s chain\n",request.ChainName)
	}

	//wait the service get LoadBalancer IP
	time.Sleep(2 * time.Second)


	//get service LoadBalancer IP and return to user
	service,err := citaChain.GetService(clientset)
	if err != nil{
		log.Println(err)
		fmt.Fprintf(w,"get servcie error : %s\n",err.Error())
	}else{
		response.URL = service.Status.LoadBalancer.Ingress[0].IP+":"+strconv.Itoa(
			int(request.ServicePort))
		log.Printf("create %s chain success",service.ObjectMeta.Name)
		fmt.Fprintf(w,"finish to create chain,the URL is: %s\n",response.URL)
	}

}

func HttpDeleteChain(w  http.ResponseWriter,r *http.Request) {
	defer func(){
		if r := recover(); r!=nil{
			log.Println("recover...",r)
		}
	}()
	request := &RequestType{}
	citaChain := &CitaChain{}
	//response := &ResponseType{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil{
		fmt.Fprintf(w,err.Error())
	}
	log.Println("request Unmarshal success")

	if err := citaChain.InitChain(request); err != nil{
		fmt.Fprintf(w,err.Error())
	}
	log.Println("init chain config success")

	clientset := InitClientset()
	log.Println("init k8s clientset success")

	if err := citaChain.DeleteChain(clientset);err != nil{
		panic(err)
	}else{
		fmt.Fprintf(w,"delete chain %s success\n",request.ChainName)
	}
}

func HttpListChain(w http.ResponseWriter,r *http.Request) {
	clientset := InitClientset()
	log.Println("init k8s client success")

	chains,err := ListChain(clientset)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w,err.Error())
	}else{
		for _,chainName := range chains{
			fmt.Fprintf(w,"chainName: %s\n",chainName)
		}
	}
}
