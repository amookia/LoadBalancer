package conf

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

type Base struct {
	Host string
	Port string
}

type Config struct {
	Listen  Base
	Service string
	Hosts   []Base
}

type Nodes []*url.URL

func Init() (Config, Nodes) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	// implement config
	var conf Config
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}
	// implement nodes
	var nodes []*url.URL
	for _, h := range conf.Hosts {
		node, _ := url.Parse(
			fmt.Sprintf("http://%s:%s", h.Host, h.Port),
		)
		nodes = append(nodes, node)
	}
	return conf, nodes
}
