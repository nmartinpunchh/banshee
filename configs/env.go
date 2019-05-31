package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Env represents the environment variables
type Env struct {
	ServiceName string `required:"true" split_words:"true"`
	Address     string `required:"true" split_words:"true"`
	//TODO: Write a custom string parser
	// https://github.com/kelseyhightower/envconfig#custom-decoders
	DbConnectionString string `required:"true" split_words:"true"`
}

// Load reads the environment config and returns a Env struct
func Load() *Env {
	var e Env
	if err := envconfig.Process("", &e); err != nil {
		log.Fatal(err.Error())
	}
	return &e
}
