package cmd

import (
	"fmt"
	"os"
)

func Deregister() error {
	conf := GetConfig()
	podName := os.Getenv("POD_NAME")
	podIP := os.Getenv("POD_IP")
	svcName := fmt.Sprintf("%v/%v", podName, podIP)

	consulClient, err := getConsulClient(conf)
	if err != nil {
		return err
	}

	if err := consulClient.Agent().ServiceDeregister(svcName); err != nil {
		return err
	}

	return nil
}
