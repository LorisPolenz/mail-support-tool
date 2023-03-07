package helper

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Configuration = load_config()

type Config struct {
	Email struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		From     string `yaml:"from"`
		Bcc      string `yaml:"bcc"`
	} `yaml:"mail"`

	Filesystem struct {
		Path string `yaml:"path"`
		Year string `yaml:"year"`
	} `yaml:"filesystem"`
}

func load_config() *Config {
	log.Println("Loading config file...")

	var config Config

	config_file, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	defer config_file.Close()

	decoder := yaml.NewDecoder(config_file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
