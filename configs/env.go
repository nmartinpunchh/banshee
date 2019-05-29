package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Env represents the environment variables
type Env struct {
	ServiceName string
	Address     string
	//TODO: Write a custom string parser
	// https://github.com/kelseyhightower/envconfig#custom-decoders
	DbConnectionString string
}

// Load reads the environment config and returns a Env struct
func Load(params) *Env {
	var e Env
	if err := envconfig.Process("", &e); err != nil {
		log.Fatal(err.Error())
	}
}
