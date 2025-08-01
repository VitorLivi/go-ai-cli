package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vitorlivi/go-ai-cli/internal/pipelines"
	pipelinecontext "github.com/vitorlivi/go-ai-cli/internal/pipelines/contexts"
)

func loadEnvs() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("No .env file found, using system environment variables")
	}
}

func main() {
	loadEnvs()
	_, err := pipelines.NewMainPipeline().Start(pipelinecontext.NewMainPipelineContext())

	if err != nil {
		log.Fatal(err)
	}
}
