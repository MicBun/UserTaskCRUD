package main

import (
	"UserSimpleCRUD/config"
	"UserSimpleCRUD/docs"
	"UserSimpleCRUD/routes"
	"UserSimpleCRUD/utils"
	"github.com/joho/godotenv"
	"log"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	docs.SwaggerInfo.Title = "Swagger Rest API CRUD User"
	docs.SwaggerInfo.Description = "This is a CRUD User and User Login API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
