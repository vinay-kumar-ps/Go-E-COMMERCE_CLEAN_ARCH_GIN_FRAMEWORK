package main

import (
	config "ecommerce/pkg/config"
	"ecommerce/pkg/di"
	"log"

	"ecommerce/cmd/api/docs"

	"github.com/joho/godotenv"
)

// @SecurityDefinition BearerAuth
// @TokenUrl /auth/token

// @securityDefinitions.Bearer		type apiKey
// @securityDefinitions.Bearer		name Authorization
// @securityDefinitions.Bearer		in header
// @securityDefinitions.BasicAuth	type basic
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = ""
	docs.SwaggerInfo.Description = "Here passion meets the fashion,This is an online store for purchasing high quality Dress of your favorite Anime.."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.BASE_URL
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
