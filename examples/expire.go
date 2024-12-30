package main

import (
	"bytetrade.io/web3os/lldap-client/pkg/cache/memory"
	"bytetrade.io/web3os/lldap-client/pkg/client"
	"bytetrade.io/web3os/lldap-client/pkg/config"
	"context"
	"k8s.io/klog/v2"
	"time"
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
	for {
		u, err := graphqlClient.Users().List(context.Background())
		if err != nil {
			panic(err)
		}

		klog.Infof("Get User: %#v", u)
		time.Sleep(5 * time.Second)

	}

}
