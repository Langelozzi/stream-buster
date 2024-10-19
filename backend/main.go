package main

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/routes"
	"github.com/STREAM-BUSTER/stream-buster/utils/database"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           DOCSWAP API
// @version         1.0
// @description     DOCSWAP is a platform designed for realtors to efficiently share, manage, and analyze documents on a neighborhood level. It enables realtors to make better use of the content created for each listing by organizing it into a searchable database that supports both uploads and downloads of documents. These documents are categorized to facilitate market analysis, putting the power of data directly into the hands of realtors and market analysts interested in housing information in the lower mainland.
// @termsOfService  http://swagger.io/terms/

// @contact.name   DOC-SWAP Support
// @contact.url    http://www.doc-swap.com/support
// @contact.email  support@doc-swap.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI Specification for DOCSWAP
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Uncomment the following line to run the db initialization for updates
	database.InitializeDb()

	// Initialize the router
	router := routes.InitRouter()
	// Set the route for accessing the Swagger UI
	router.GET("api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start the server: ", err)
		return
	}
}
