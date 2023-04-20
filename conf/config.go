package conf

import (
	"log"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Listen string   `toml:"listen"`
	Nodes  []string `toml:"nodes"`
}

type Nodes []*url.URL

func Init() (Config, Nodes) {
	file, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	var conf Config
	toml.Unmarshal(file, &conf)
	var nodes []*url.URL
	for _, j := range conf.Nodes {
		node, err := url.Parse(j)
		if err != nil {
			log.Fatalln(err)
		}
		nodes = append(nodes, node)
	}
	return conf, nodes
}
