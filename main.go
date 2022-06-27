package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matg94/go-url-shortener/config"
)

func LoadConfig(profile string) *config.AppConfig {
	data, err := config.ReadConfigFile(fmt.Sprintf("./config/%s.yaml", profile))
	if err != nil {
		log.Fatalf("Could not read config file for profile: %s", profile)
	}
	appConfig, err := config.ParseYamlConfig(data)
	if err != nil {
		log.Fatalf("Could not parse config for profile: %s", profile)
	}
	return appConfig
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Application profile is required as an Arg to run")
	}
	profile := os.Args[1]
	appConfig := LoadConfig(profile)
	fmt.Println(appConfig)
}
