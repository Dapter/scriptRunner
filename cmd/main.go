package main

import (
	myEntryPoint "github.com/Dapter/scriptRunner/internal/adapters/left/desktop"
	myYaml "github.com/Dapter/scriptRunner/internal/adapters/right/yaml"
	myApp "github.com/Dapter/scriptRunner/internal/app"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	yamlParser := myYaml.New()
	app := myApp.New(yamlParser)
	entrypoint := myEntryPoint.New(app)

	entrypoint.Start()

	os.Exit(0)
}
