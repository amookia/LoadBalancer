package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/amookia/loadbalancer/conf"
	"github.com/amookia/loadbalancer/internal/handler"
)

func main() {
	config := conf.Init()
	var nodes []*url.URL
	for _, j := range config.Nodes {
		node, err := url.Parse(j)
		if err != nil {
			log.Fatalln(err)
		}
		nodes = append(nodes, node)
	}

	handler := handler.NewHandler(nodes)
	proxy := http.HandlerFunc(handler.BalancerHandler)

	http.ListenAndServe(":8081", proxy)

}
