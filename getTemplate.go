package main

import(
	appsv1 "k8s.io/api/apps/v1"
	//"k8s.io/apimachinery/pkg/api/resource"
	corev1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/api/resource"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Err struct{
	Name string
	Err error
}

func (e Err) Error() string{
	return fmt.Sprintf("%s:\n%v",e.Name,e.Err)
}

func getDeploymentTemplate(file string,v *appsv1.Deployment) error {
	data,err := ioutil.ReadFile(file)
	if err != nil{
		return Err{Name: "getDeploymentTemplate function error",Err: err}
	}
	err = json.Unmarshal(data,v)
	if err != nil{
		return Err{Name: "getDeploymentTemplate function error",Err: err}
	}

	return nil
}

func getPVCTemplate(file string,v *corev1.PersistentVolumeClaim) error {
	data,err := ioutil.ReadFile(file)
	if err != nil{
		return Err{Name: "getPVCTemplate function error",Err: err}
	}
	err = json.Unmarshal(data,v)
	if err != nil{
		return Err{Name: "getPVCTemplate function error",Err: err}
	}

	return nil
}

func getSVCTemplate(file string,v *corev1.Service) error{
	data,err := ioutil.ReadFile(file)
	if err != nil {
		return Err{Name: "getSVCTemplate function error",Err: err}
	}

	err = json.Unmarshal(data,v)
	if err != nil{
		return Err{Name: "getSVCTemplate function error",Err: err}
	}

	return nil
}


