package app

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	clientset "k8s.io/client-go/kubernetes"
)

type CapCop struct {
	KubeConfig           string
	ClientSet            clientset.Interface
}

func (cop *CapCop) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&cop.KubeConfig, "kubeconfig", cop.KubeConfig, "Kubeconfig used to contact the Kubernetes cluster.")
	viper.BindPFlag("kubeconfig", fs.Lookup("kubeconfig"))
}

