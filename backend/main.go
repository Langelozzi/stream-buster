package main

import (
	"fmt"

	"github.com/STREAM-BUSTER/stream-buster/docs"
	"github.com/STREAM-BUSTER/stream-buster/routes"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	"github.com/STREAM-BUSTER/stream-buster/utils/database/post_deployment_functions"

	// "github.com/STREAM-BUSTER/stream-buster/utils/database"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           StreamBuster API
// @version         1.0
// @description     A Streaming app inspired by the OG's
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      api.streambuster.xyz
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI Specification for StreamBuster
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Uncomment the following line to run the db initialization for updates
	//database.InitializeDb()

	// Initialize the router
	router := routes.InitRouter()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Set the route for accessing the Swagger UI
	router.GET("api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	post_deployment_functions.CreateEndpointRecords(database.GetInstance(), router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start the server: ", err)
		return
	}
}
