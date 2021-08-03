package main

import(
	"testing"
	//"fmt"
	//"encoding/json"
	//"bytes"
)

//func showJson(v interface{}) {
//	data,_ := json.Marshal(v)
//	var out bytes.Buffer
//	json.Indent(&out,data,"","\t")
//	fmt.Println(out.String())
//}

//func TestInit(t *testing.T) {
//	citaChain := &CitaChain{}
//	err := citaChain.Init("sm2")
//	if err != nil{
//		t.Fatal(err)
//	}
//
//	//fmt.Println(citaChain)
//}
//
//func TestInitChain(t *testing.T) {
//	citaChain := &CitaChain{}
//	request := &RequestType {
//		ChainName: "test-chain-chainName",
//		ServicePort: 1000,
//		StorageSize: "50Gi",
//		ChainType: "secp256",
//	}
//
//	err := citaChain.InitChain(request)
//	if err != nil{
//		t.Error(err)
//	}
//
//	showJson(citaChain.Service)
//	showJson(citaChain.Deployment)
//	showJson(citaChain.PersistentVolumeClaim)
//}

func TestListChain(t *testing.T) {
	citaChain := &CitaChain{}
	//request := &RequestType {
	//	ChainName: "test-chain-chainName",
	//	ServicePort: 1000,
	//	StorageSize: "50Gi",
	//	ChainType: "secp256",
	//}
	_,err := citaChain.ListChain(InitClientset())
	if err != nil{
		t.Error(err)
	}

}
