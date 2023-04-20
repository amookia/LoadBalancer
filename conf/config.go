package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Listen string   `toml:"listen"`
	Nodes  []string `toml:"nodes"`
}

func Init() Config {
	file, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	var conf Config
	toml.Unmarshal(file, &conf)
	return conf
}
