package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.ushareit.me/sgt/hawkeye/ping-monitor/cmd"
)

var cfgFile = flag.String("config", "/etc/config/config.prod.yaml", "config path")

func main() {
	flag.Parse()

	if err := cmd.InitConfig(*cfgFile); err != nil {
		log.Println("get config error:", err)
		os.Exit(1)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGTERM:
				log.Println("Exiting, sleep 5s")
				if err := cmd.Deregister(); err != nil {
					log.Println("Consul service deregister failed:", err)
				}
				time.Sleep(5 * time.Second)
				os.Exit(0)
			default:
				log.Println("Receive other signal")
			}
		}
	}()

	if err := cmd.Register(); err != nil {
		log.Println("Consul service register failed:", err)
		os.Exit(1)
	}

	if err := cmd.HealthCheck(); err != nil {
		log.Println("Start http service error:", err)
	}
}
