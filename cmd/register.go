package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

const (
	consulService = "ping-probe"
	port          = 9346
)

func Register() error {
	conf := GetConfig()
	podName := os.Getenv("POD_NAME")
	podIP := os.Getenv("POD_IP")
	region := os.Getenv("REGION")
	vpc := os.Getenv("VPC")
	account := os.Getenv("ACCOUNT")

	tags := []string{
		fmt.Sprintf("region=%v", region),
		fmt.Sprintf("vpc=%v", vpc),
		fmt.Sprintf("account=%v", account),
	}

	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%s/-/healthy", podIP, conf.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}

	pod := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v/%v", podName, podIP),
		Name:    consulService,
		Tags:    tags,
		Port:    port,
		Address: podIP,
		Check:   check,
	}

	consulClient, err := getConsulClient(conf)
	if err != nil {
		return err
	}

	if err := consulClient.Agent().ServiceRegister(pod); err != nil {
		return err
	}

	log.Println("register successfully")

	return nil
}

func getConsulClient(cfg Conf) (*api.Client, error) {
	cCfg := &api.Config{
		Address: cfg.Consul.Address,
		Scheme:  cfg.Consul.Scheme,
	}

	clt, err := api.NewClient(cCfg)
	if err != nil {
		return nil, err
	}

	return clt, nil
}
