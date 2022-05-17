package boom

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
)

// WithContainer adds a container to the server
func WithContainer(c Container) APIOption {
	return func(server *API) {
		server.Container = c
	}
}

// WithSettings sets server configurations
func WithSettings(settings *Settings) APIOption {
	return func(server *API) {
		server.Settings = settings
	}
}

// WithCORS enables CORS
func WithCORS() APIOption {
	return func(server *API) {
		server.Cors = true
	}
}

// WithBaseHandler add a base handler
func WithBaseHandler(h gin.HandlerFunc) APIOption {
	return func(server *API) {
		server.Handlers = append(server.Handlers, h)
	}
}

// WithHealthz default response time.now()
func WithHealthz() APIOption {
	return func(server *API) {
		server.Healthz = true
	}
}

// WithSwagger ...
func WithSwagger(spec *swag.Spec, path string) APIOption {
	return func(server *API) {
		server.Swagger = &SwaggerSettings{
			Spec: spec,
			Path: path,
		}
	}
}
