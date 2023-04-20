package main

import (
	"fmt"
	"net/http"

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
	err = http.ListenAndServe(
		fmt.Sprintf(":%s", config.Listen), // listen
		proxy,                             // handler
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
}