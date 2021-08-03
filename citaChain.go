package main

import(
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"context"
	"strconv"
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


func (citaChain *CitaChain)InitChain(request *RequestType) error{
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

func (citaChain *CitaChain) CreateChain(client *kubernetes.Clientset) error{
	if err := citaChain.createPersistentVolumeClaims(client);err != nil{
		return Err{Name:"CreateChain function error: ",Err:err}
	}

	if err := citaChain.createDeployments(client); err != nil{
		return Err{Name:"CreateChain function error: ",Err:err}
	}

	if err := citaChain.createServices(client); err != nil{
		return Err{Name:"CreateChain function error: ",Err:err}
	}

	return nil
}

func (citaChain *CitaChain) createDeployments(client *kubernetes.Clientset) error{
	_,err := client.AppsV1().Deployments("cita").Create(context.TODO(),&citaChain.Deployment,metav1.CreateOptions{})
	if err != nil {
		return Err{Name:"createDeployment function error: ",Err:err}
	}

	return nil
}

func (citaChain *CitaChain) createPersistentVolumeClaims(client *kubernetes.Clientset) error{
	_,err := client.CoreV1().PersistentVolumeClaims("cita").Create(context.TODO(),&citaChain.PersistentVolumeClaim,metav1.CreateOptions{})
	if err != nil {
		return Err{Name:"creatPersistentVolumeClaims function error: ",Err: err}
	}

	return nil
}

func (citaChain *CitaChain) createServices(client *kubernetes.Clientset) error{
	_,err := client.CoreV1().Services("cita").Create(context.TODO(),&citaChain.Service,metav1.CreateOptions{})

	//log.Println(service.Status.LoadBalancer.Ingress[0].IP)
	if err != nil {
		return Err{Name:"createServices function error: ",Err: err}
	}

	return nil
}

func (citaChain *CitaChain) DeleteChain(client *kubernetes.Clientset) error {
	if err := citaChain.deleteServices(client); err != nil{
		return Err{Name: "DeleteCitaChiain error:",Err: err}
	}

	if err := citaChain.deleteDeployments(client);err != nil{
		return Err{Name: "DeleteCitaChiain error:",Err: err}
	}

	if err := citaChain.deletePersistentVolumeClaims(client); err != nil{
		return Err{Name: "DeleteCitaChiain error:",Err: err}
	}
	return nil
}

func (citaChain *CitaChain) deleteServices(client *kubernetes.Clientset) error{
	err := client.CoreV1().Services("cita").Delete(context.TODO(),citaChain.Service.ObjectMeta.Name,metav1.DeleteOptions{})
	if err != nil{
		return Err{Name:"deleteServices function error: ",Err: err}
	}

	return nil
}

func (citaChain *CitaChain) deletePersistentVolumeClaims(client *kubernetes.Clientset) error {
	err := client.CoreV1().PersistentVolumeClaims("cita").Delete(context.TODO(),citaChain.PersistentVolumeClaim.ObjectMeta.Name,metav1.DeleteOptions{})
	if err != nil{
		return Err{Name:"deletePersistentVolumeClaims function error: ",Err: err}
	}

	return nil
}

func (citaChain *CitaChain) deleteDeployments(client *kubernetes.Clientset) error{
	err := client.AppsV1().Deployments("cita").Delete(context.TODO(),citaChain.Deployment.ObjectMeta.Name,metav1.DeleteOptions{})

	if err != nil{
		return Err{Name:"deleteDeployment function error: ",Err: err}
	}

	return nil
}

func (citaChain *CitaChain) GetService(client *kubernetes.Clientset) (*corev1.Service,error) {
	service, err := client.CoreV1().Services("cita").Get(context.TODO(),citaChain.Service.ObjectMeta.Name,metav1.GetOptions{})
	if err != nil{
		return nil,Err{Name:"GetService function error: ",Err:err}
	}
	return service,nil
}

func ListChain(client *kubernetes.Clientset) ([]string,error) {
	serviceList,err := client.CoreV1().Services("cita").List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return nil,Err{Name:"ListChain function error: ",Err: err}
	}

	var services []string
	for _,service := range serviceList.Items{
		serviceURL := service.ObjectMeta.Name + " " + service.Status.LoadBalancer.Ingress[0].IP + ":"+strconv.Itoa(int(service.Spec.Ports[0].Port))
		services = append(services,serviceURL)
	}

	if len(services) > 0 {
		return services,nil
	}else{
		return nil,Err{Name:"there no chain"}
	}
}
