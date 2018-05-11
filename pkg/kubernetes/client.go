package kubernetes

import (
	"path"

	log "github.com/Sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/util/homedir"
)

func DefaultClientSet(kubeconfig string) *kubernetes.Clientset {

	var clientset *kubernetes.Clientset

	if len(kubeconfig) == 0 {
		kubeconfig = path.Join(homedir.HomeDir(), ".kube", "config")
	}

	outClusterConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Errorf("Unable to build out of cluster config, continuing to try in cluster config")
	} else if outClusterConfig != nil {
		clientset, err = kubernetes.NewForConfig(outClusterConfig)
		if err != nil {
			log.Errorf("Unable to build out of cluster client, continuing to try in cluster client")
		}
		if clientset != nil {
			log.Infof("Successfully created an out of cluster client")
			return clientset
		}
	}

	inClusterConfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	} else {
		clientset, err = kubernetes.NewForConfig(inClusterConfig)
		if err != nil {
			panic(err.Error())
		}
	}

	return clientset

}
