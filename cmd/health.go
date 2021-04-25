package cmd

import (
	"log"
	"net/http"
	"time"
)

func HealthCheck() error {
	log.Println("Starting http service on 9347...")

	// for consul
	var healthCheckCount int
	timeTickChan := time.NewTicker(time.Second * 30)
	http.HandleFunc("/-/healthy", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Health check from:", r.RemoteAddr)
		healthCheckCount += 1
		w.Write([]byte("Ping prober is healthy"))
	})

	// for k8s
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Re-register
	go func() {
		for {
			select {
			case <-timeTickChan.C:
				if healthCheckCount < 1 {
					log.Println("No health check within 30s, re-registering...")
					Deregister()
					if err := Register(); err != nil {
						log.Println("Register failed, re-register 30s later")
					}
				}
				healthCheckCount = 0
			}
		}
	}()

	if err := http.ListenAndServe(":9347", nil); err != nil {
		return err
	}
	return nil
}
