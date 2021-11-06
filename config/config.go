package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Zipkin struct {
		Url         string `yaml:"url"         env:"ZIPKIN_URL"          env-default:"http://localhost:9411"`
		Endpoint    string `yaml:"endpoint"    env:"ZIPKIN_ENDPOINT"     env-default:"/api/v2/spans"`
		HostPort    string `yaml:"hostPort"    env:"ZIPKIN_HOST_PORT"    env-default:"localhost:80"` // https://github.com/openzipkin-contrib/zipkin-go-opentracing/issues/45
		ServiceName string `yaml:"serviceName" env:"ZIPKIN_SERVICE_NAME" env-default:"TestService2"`
	} `yaml:"zipkin"`
}

var Cfg ConfigDatabase

func GetConfig(profilerConfigFile string) {
	err := cleanenv.ReadConfig(profilerConfigFile, &Cfg)
	if err != nil {
		log.Println("Configuration file was not found! Using default values")
		readEnvErr := cleanenv.ReadEnv(&Cfg)
		if readEnvErr != nil {
			log.Println("Error while reading config file")
		}
	}
}
