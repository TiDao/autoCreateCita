package main

import(
	"k8s.io/apimachinery/pkg/api/resource"
)

type RequestType struct{
	ChainName string `json: "chainName"`
	ServicePort int32 `json: "servicePort"`
	StorageSize string `json: "storageSize"`
	ChainType string `json: "chainType"`
}

type ResponseType struct{
	URL string `json: "URL"`
}

func (request *RequestType)Createchain() error{
	citaChain := &CitaChain{}

	//set chainType(sm2 or secp256) to choose template
	if err := citaChain.Init(request.ChainType); err != nil{
		return err
	}

	//config core v1 PersistentVolumeClaim Object
	citaChain.PersistentVolumeClaim.ObjectMeta.Name = request.ChainName
	citaChain.PersistentVolumeClaim.Spec.Resources.Requests["storage"] = resource.MustParse(request.StorageSize)
	citaChain.PersistentVolumeClaim.Spec.Resources.Limits["storage"] = resource.MustParse(request.StorageSize)

	//config core v1 Service Object
	citaChain.Service.ObjectMeta.Name = request.ChainName
	citaChain.Service.Spec.Ports[0].Port = ServicePort
	citaChain.Service.Spec.Selector["cita"] = request.ChainName

	//config apps v1 Deployment Object

}


