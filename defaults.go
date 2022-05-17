package boom

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"google.golang.org/grpc"
)

// APIOption wrapps all server configurations
type APIOption func(server *API)

// SwaggerSettings ...
type SwaggerSettings struct {
	Spec *swag.Spec
	Path string
}

// API ...
type API struct {
	Container Container
	Cors      bool
	Engine    *gin.Engine
	Healthz   bool
	Handlers  []gin.HandlerFunc
	Settings  *Settings
	Swagger   *SwaggerSettings
	Router    *gin.RouterGroup
	Grpc      *grpc.Server
}
