package main

import (
	"log"

	"github.com/Mycunycu/carsales/gateway/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal("error in initialize application: ", err)
	}
}
