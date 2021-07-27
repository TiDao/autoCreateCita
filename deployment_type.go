package main

import(
	"k8s.io/api/apps/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"io/ioutil"
	"encoding/json"
)

var deployment = &v1.Deployment{}

func getTemplate(file string,deployment *v1.Deployment) error {
	data,err := ioutil.ReadFile(file)
	if err != nil{
		return err
	}
	err = json.Unmarshal(data,deployment)
	if err != nil{
		return err
	}

	return nil
}


