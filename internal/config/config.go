package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

// type Config struct {
// 	Env string   // capital // parallel to config file
// 	storagePath string
// 	HTTPServer
// }

// serialize the struct in file // can do maually // or using package cleanenv
// with this we can fetch data from the yaml and env files

// struct tax --> `yaml:"env"  env:"ENV" env-required:"true" env-default: "production"`
// fetching "env" from yaml file, not present then env file , setting required makes it important to be present
// setting env-default to "production" will be safe guard so when something goes wrong and the dev config aren't fetched the production one will be used so logs etc aren't shown to user

type Config struct {
	// Env 		string  `yaml:"env"  env:"ENV" env-required:"true" env-default: "production"`
	Env         string `yaml:"env"  env:"ENV" env-required:"true"`
	storagePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// parsing login

func MustLoad() *Config{
	// any error in "Must" Load means the code ahead should not be used // do not return error from MustLoad because "Must" mus twork if error occurs just print error

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// check in flags
		// go run file_path -config-path xyz

		flags := flag.String("config", "", "path to the configuration file" )

		configPath = *flags // dereference flag pointer

		if configPath == ""{
			log.Fatalf("Config path is not set")
		}
 	}

	// check if file not available in path
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)	
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg) // put configpath in cfg
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg // return point


}
