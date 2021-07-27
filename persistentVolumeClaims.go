package main

import(
	"k8s.io/api/core/v1"
	"io/ioutil"
	"encoding/json"
)

var pvcTemplate = &v1.PersistentVolumeClaim{}

func getPVCTemplate(file string,pvcTemplate *v1.PersistentVolumeClaim) error{
	data,err := ioutil.ReadFile(file)
	if err != nil{
		return err
	}

	err = json.Unmarshal(data,pvcTemplate)
	if err != nil{
		return err
	}

	return nil
}
