package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/util/logs"
	"github.com/spf13/pflag"
	"github.com/ebayinc/capcop/cmd/app"
	"github.com/ebayinc/capcop/pkg/kubernetes"
)

var cmd = &cobra.Command{
	Use:   "capcop",
	Short: "capcop.",
	Long:  `capcop`,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

var cop *app.CapCop

func init() {
	cop := &app.CapCop{}
	cop.AddFlags(pflag.CommandLine)
	defer logs.FlushLogs()
}


func Run() {
	log.Infof("Running capcop")
}

func main() {
	cop.ClientSet = kubernetes.DefaultClientSet(cop.KubeConfig)
	cmd.Execute()
	return
}
