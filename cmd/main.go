package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amookia/loadbalancer/conf"
	"github.com/amookia/loadbalancer/internal/handler"
	"github.com/amookia/loadbalancer/pkg/logger"
)

func main() {
	config, nodes := conf.Init()
	logger, err := logger.New(config.Service)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// register handlers
	handler := handler.NewHandler(nodes, logger)
	proxy := http.HandlerFunc(handler.BalancerHandler)

	// listening
	logger.Infof("service started at port : %s", config.Listen)

	srv := http.Server{
		Addr: fmt.Sprintf("%s:%s",
			config.Listen.Host,  // host
			config.Listen.Port), // port
		WriteTimeout: 1 * time.Second, // write timeout
		ReadTimeout:  1 * time.Second, // read timeout
		Handler:      http.TimeoutHandler(proxy, 1*time.Second, "<b>500</b>\n"),
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
