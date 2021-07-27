package main

import(
	"k8s.io/api/core/v1"
	"fmt"
	"testing"
)

func TestGetPVCTemplate(t *testing.T) {
	file := "./template/cita_pvc.json"
	pvc := &v1.PersistentVolumeClaim{}

	if err := getPVCTemplate(file,pvc);err != nil{
		t.Fatal(err)
	}
	fmt.Println(pvc)
}

