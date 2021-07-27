package main

import(
	"testing"
	"k8s.io/api/apps/v1"
	"fmt"
)

func TestGetTemplate(t *testing.T) {
	file := "./template/cita-deployment-sm2.json"
	deployment := &v1.Deployment{}

	err := getTemplate(file,deployment)
	fmt.Println(deployment)
	if err != nil{
		t.Fatal(err)
	}
}
