package main

import (
	"github.com/Alica032/bitbucket/pkg/telegram"
	"gopkg.in/yaml.v2"
	"os"
)

var PATH = "config.yaml"

func main() {
	var config telegram.Config
	file, _ := os.ReadFile(PATH)
	_ = yaml.Unmarshal(file, &config)
	telegram.RunServer(&config.Server, &config.Bot, config.IsFirstStart)
}
