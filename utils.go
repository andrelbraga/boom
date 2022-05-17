package boom

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// FromYAML ...
func FromYAML(file string, dist interface{}) error {
	filename, _ := filepath.Abs(file)

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, dist)
}

// Swagger ...
func Swagger(file string) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := os.Open(file)

		if err != nil {
			logrus.WithError(err).
				Error("swagger file error")

			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		defer file.Close()

		content, _ := ioutil.ReadAll(file)

		c.Writer.Write(content)
		c.Writer.Header().Set("content-type", "application/json")
	}
}
