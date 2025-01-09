package main

import (
	"context"

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
		klog.Infof("create graphqlClient err=%v", err)
	}

	err = graphqlClient.Users().ResetPassword(context.TODO(), "test002", "test002test002")

	if err != nil {
		panic(err)
	}

}
