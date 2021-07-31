package main

import(
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"context"
)

//var deployment = &v1.Deployment{}
type CitaChain struct{
	//ChainType string
	Deployment appsv1.Deployment
	PersistentVolumeClaim corev1.PersistentVolumeClaim
	Service corev1.Service
}

func (c *CitaChain) Init(chainType string) error{
	deploymentFile := ""
	if chainType == "secp256" {
		deploymentFile = "./template/cita-deployment-secp256.json"
	}else{
		deploymentFile = "./template/cita-deployment-sm2.json"
	}
	PVCFile := "./template/cita-pvc.json"
	SVCFile := "./template/svc.json"

	if err := getDeploymentTemplate(deploymentFile,&c.Deployment); err != nil{
		return Err{Name: "init function error",Err: err}
	}

	if err := getPVCTemplate(PVCFile,&c.PersistentVolumeClaim); err != nil{
		return Err{Name: "init function error",Err: err}
	}

	if err := getSVCTemplate(SVCFile,&c.Service); err != nil{
		return Err{Name: "init function error",Err: err}
	}

	//c.PersistentVolumeClaim.Spec.Resources.Requests["storage"] = resource.MustParse("10Gi")

	return nil

}


func (citaChain *CitaChain)CreateChain(request *RequestType) error{

	//set chainType(sm2 or secp256) to choose template
	if err := citaChain.Init(request.ChainType); err != nil{
		return Err{Name: "CreateChain function error",Err: err}
	}

	//config core v1 PersistentVolumeClaim Object
	citaChain.PersistentVolumeClaim.ObjectMeta.Name = request.ChainName
	citaChain.PersistentVolumeClaim.Spec.Resources.Requests["storage"] = resource.MustParse(request.StorageSize)
	citaChain.PersistentVolumeClaim.Spec.Resources.Limits["storage"] = resource.MustParse(request.StorageSize)

	//config core v1 Service Object
	citaChain.Service.ObjectMeta.Name = request.ChainName
	citaChain.Service.Spec.Ports[0].Port = request.ServicePort
	citaChain.Service.Spec.Selector["cita"] = request.ChainName

	//config apps v1 Deployment Object
	citaChain.Deployment.ObjectMeta.Name = request.ChainName
	citaChain.Deployment.Spec.Selector.MatchLabels["cita"] = request.ChainName
	citaChain.Deployment.Spec.Template.ObjectMeta.Name = request.ChainName
	citaChain.Deployment.Spec.Template.ObjectMeta.Labels["cita"] = request.ChainName
	for i,_ := range citaChain.Deployment.Spec.Template.Spec.Volumes {
		switch citaChain.Deployment.Spec.Template.Spec.Volumes[i].Name{
		case "data-pvc":
			citaChain.Deployment.Spec.Template.Spec.Volumes[i].PersistentVolumeClaim.ClaimName = request.ChainName
		default:
			continue
		}
	}

	return nil
}


func (citaChain *CitaChain) CreateDeployment(client *kubernetes.Clientset) error{
	_,err := client.AppsV1().Deployment("cita").Create(context.TODO(),citaChain.Deployment,metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (citaChain *CitaChain) CreatePersistentVolumeClaim(client *kubernetes.Clientset) error{
	_,err := client.CoreV1().PersistentVolumeClaim("cita").Create(context.TODO(),citaChain.PersistentVolumeClaim,metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (citaChain *CitaChain) CreateService(client *kubernetes.Clientset) error{
	_,err := client.CoreV1().Service("cita").Create(context.TODO(),citaChain.Service,metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

