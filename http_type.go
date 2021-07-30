package main


type RequestType struct{
	ChainName string `json: "chainName"`
	ServicePort int32 `json: "servicePort"`
	StorageSize string `json: "storageSize"`
	ChainType string `json: "chainType"`
}

type ResponseType struct{
	URL string `json: "URL"`
}



