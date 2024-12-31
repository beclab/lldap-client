package main

import (
	"github.com/beclab/lldap-client/pkg/auth"
	"k8s.io/klog/v2"
	"os"
)

func main() {
	r, err := auth.Login("http://127.0.0.1:17170", "admin", "adminpassword")
	if err != nil {
		klog.Infof("Login err=%v", err)
		os.Exit(0)
	}
	klog.Infof("resp: %v", r)
}
