package main

import (
	"github.com/nmartinpunchh/banshee/internal/service"
)

func main() {
	if err := service.Run(); err != nil {
		panic(err)
	}
}
