package boom

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	gin.SetMode("release")
}

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
	if server.Healthz != nil {
		server.Router.GET("/healthz", server.Healthz)
	}

	// If existe setup Swagger
	if server.Swagger != nil {
		server.Router.GET("/swagger", Swagger(server.Settings.Http.Swagger))
		swaggerPath := fmt.Sprintf("%s*any", server.Swagger.Path)
		logrus.Infof("Swagger at http://localhost:%s%sindex.html", server.Settings.Http.Host, server.Swagger.Path)
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

// Run starts the server.
func (server *API) Run() {
	defer server.Container.Close()

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", server.Settings.Http.Host),
		Handler: server.Engine,
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		sig := <-sigs

		logrus.Infof("caught sig: %+v", sig)
		logrus.Info("waiting 5 seconds to finish processing")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logrus.WithField("error", err).Error("shutdown error")
		}
	}()

	logrus.Infof("start api %s", server.Settings.Http.Host)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithField("error", err).Info("startup error")
	}
}
