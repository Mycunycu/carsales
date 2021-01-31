package main

import "carsales/internal/app"

const configPath = "config/main"

func main() {
	app.Run(configPath)
}
