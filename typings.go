package boom

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"google.golang.org/grpc"
)

// APIOption wrapps all server configurations
type APIOption func(server *API)

// HealtzOption ...
type HealtzOption func(*Healthz)

// SwaggerSettings ...
type SwaggerSettings struct {
	Spec *swag.Spec
	Path string
}

// Healthz ...
type Healthz struct {
	Settings *Settings
	Checks   []func(*Healthz) error
}

// API ...
type API struct {
	Engine     *gin.Engine
	Router     *gin.RouterGroup
	Cors       bool
	Settings   *Settings
	Healthz    gin.HandlerFunc
	Handlers   []gin.HandlerFunc
	Swagger    *SwaggerSettings
	GrpcServer *grpc.Server
	Container  Container
}
