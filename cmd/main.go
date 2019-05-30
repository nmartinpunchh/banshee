package main

import (
	"log"

	"github.com/nmartinpunchh/banshee/internal/service"
)

func main() {
	log.Println("Running GRPC Server")
	if err := service.Run(); err != nil {
		panic(err)
	}
}
