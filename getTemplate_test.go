package main

import(
	"testing"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	"fmt"
)

func TestGetTemplate(t *testing.T) {
	file := "./template/cita-deployment-sm2.json"
	deployment := &appsv1.Deployment{}

	err := getTemplate(file,deployment)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(deployment)
}

func TestGetPVCTemplate(t *testing.T) {
	file := "./template/cita-pvc.json"
	pvc := &corev1.PersistentVolumeClaim{}

	err := getPVCTemplate(file,pvc)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pvc)
}

func TestGetSVCTemplate(t *testing.T) {
	file := "./template/svc.json"
	svc := &corev1.Service{}

	err := getSVCTemplate(file,svc)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(svc)
}
