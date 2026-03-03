package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

// env-default:"production"
type config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true"`
	Storagepath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

func MustLoad() *config{ // in the functions named as must then do not return error just run fatal
	
	var configPath string

	if configPath == ""{
		configPath = os.Getenv("CONFIG_PATH")

		flags := flag.String("config","","path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _,err := os.Stat(configPath) ; os.IsNotExist(err){
		log.Fatalf("config file does not exist : %s" , configPath) // fatalF is used because we are formatting
	} 
	
	var cfg config // we are making an inititalization of the struct 

	err := cleanenv.ReadConfig(configPath , &cfg ) // installed package will be imported by this 
	// this will return error 

	if err != nil {
		log.Fatalf("can not read config file :%s" , err.Error())
	}

	return &cfg 
}