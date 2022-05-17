package boom

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New Inicia as principais configurações da engine Gin Framework
func New(opts ...APIOption) *API {

	// Instance API type
	server := &API{}

	// Handlers gin framework
	server.Handlers = []gin.HandlerFunc{}

	for _, opt := range opts {
		opt(server)
	}

	// Instance engine gin
	server.Engine = gin.New()

	// Recovery if panic
	server.Engine.Use(gin.Recovery())

	// Router initial
	server.Router = server.Engine.Group("")

	// If CORS
	if server.Cors {
		server.Engine.Use(CORS())
	}

	// If not exist route, response status not found
	server.Engine.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Healthz if exist handler function
	if server.Healthz {
		logrus.Infof("start heatlth check at http://localhost:%s/healthz", server.Settings.Http.Host)
		server.Router.GET("/healthz", Healthz())
	}

	// If existe setup Swagger
	if server.Swagger != nil {
		server.Router.GET("/swagger", Swagger(server.Settings.Http.Host))
		swaggerPath := fmt.Sprintf("/%s/*any", server.Swagger.Path)
		logrus.Infof("start swagger at http://localhost:%s/%s/index.html", server.Settings.Http.Host, server.Swagger.Path)
		server.Router.GET(swaggerPath, ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//If Handlers
	if server.Handlers != nil {
		server.Router.Use(server.Handlers...)
	}

	for _, ctrl := range server.Container.Controllers() {
		ctrl.RegisterRoutes(server.Router)
	}
	return server
}
