package boom
import "github.com/gin-gonic/gin"

type Container interface {
	Close() error
	Controllers() []APIController
}

type APIController interface {
	RegisterRoutes(*gin.RouterGroup)
}