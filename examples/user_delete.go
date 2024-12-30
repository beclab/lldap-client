package main

import (
	"bytetrade.io/web3os/lldap-client/pkg/cache/memory"
	"bytetrade.io/web3os/lldap-client/pkg/client"
	"bytetrade.io/web3os/lldap-client/pkg/config"
	"context"
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

	err = graphqlClient.Users().Delete(context.TODO(), "test001")

	if err != nil {
		panic(err)
	}

}
