package api

import (
	"user/api/handler"
	"user/pkg/logger"
	"user/service"
	"user/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func New(store storage.IStorage, service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(store, service, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/v1/user", h.CreateUser)
	r.POST("/api/v1/users", h.CreateMany)
	r.PUT("/api/v1/user", h.Update)
	r.PUT("/api/v1/users", h.UpdateMany)
	r.DELETE("/api/v1/user", h.Delete)
	r.DELETE("/api/v1/users", h.DeleteMany)
	r.GET("/api/v1/user/:id", h.GetById)
	r.GET("/api/v1/users", h.GetList)

	return r
}
