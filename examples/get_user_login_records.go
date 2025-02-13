package main

import (
	"context"
	"os"

	"github.com/beclab/lldap-client/pkg/cache/memory"
	"github.com/beclab/lldap-client/pkg/client"
	"github.com/beclab/lldap-client/pkg/config"
	"k8s.io/klog/v2"
)

func main() {
	cfg := config.Config{
		Host:       "http://127.0.0.1:17170",
		Username:   "admin",
		Password:   "adminpassword",
		TokenCache: memory.New(),
	}
	graphqlClient, err := client.New(&cfg)
	if err != nil {
		klog.Infof("create graphClient err=%v", err)
	}
	records, err := graphqlClient.Users().LoginRecords(context.TODO(), "admin")
	if err != nil {
		klog.Infof("get login records err=%v", err)
		os.Exit(0)
	}
	klog.Infof("records: %v", records)
	klog.Infof("len(records):", len(records))
}
