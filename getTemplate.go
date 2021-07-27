package main

import(
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"io/ioutil"
	"encoding/json"
)

//var deployment = &v1.Deployment{}
type CitaChain struct{
	ChainType string
	Deployment *appv1.Deployment
	PersistentVolumeClaim *corev1.PersistentVolumeClaim
	Service *corev1.Service
}

func (c *CitaChain) Init() error{
	deploymentFile = ""
	if c.ChainType == "secp256" {
		deployment = "./template/cita-deployment-secp256.json"
	}else{
		deploymentFile = "./template/cita-deployment-sm2.json"
	}
	PVCFile := "./tempalate/cita-pvc.json"
	SVCFile := "./tempalate/svc.json"

	if err := getDeploymentTemplate(deploymentFile,c.Deployment); err != nil{
		return err
	}

	if err := getPVCTemplate(PVCFile,c.PersistentVolumeClaim); err != nil{
		return err
	}

	if err := getSVCTemplate(SVCFile,c.Service); err != nil{
		return err
	}

	return nil

}

func getDeploymentTemplate(file string,v *appsv1.Deployment) error {
	data,err := ioutil.ReadFile(file)
	if err != nil{
		return err
	}
	err = json.Unmarshal(data,v)
	if err != nil{
		return err
	}

	return nil
}

func getPVCTemplate(file string,v *corev1.PersistentVolumeClaim) error {
	data,err := ioutil.ReadFile(file)
	if err != nil{
		return err
	}
	err = json.Unmarshal(data,v)
	if err != nil{
		return err
	}

	return nil
}

func getSVCTemplate(file string,v *corev1.Service) error{
	data,err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data,v)
	if err != nil{
		return err
	}

	return nil
}
