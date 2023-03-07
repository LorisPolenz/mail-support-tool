package helper

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Email struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		From     string `yaml:"from"`
		Bcc      string `yaml:"bcc"`
	} `yaml:"mail"`
}

func Load_config() *Config {
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

	if config.Email.Username == "" {
		log.Println("Loading Username from environment variables.")
		config.Email.Username = os.Getenv("OMA_TOOL_EMAIL_USERNAME")
	}

	if config.Email.Password == "" {
		log.Println("Loading Password from environment variables.")
		config.Email.Password = os.Getenv("OMA_TOOL_EMAIL_PASSWORD")
	}

	return &config
}
