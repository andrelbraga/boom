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
)

func init() {
	gin.SetMode("release")
}

// ServerHttp ...
func (server *API) ServerHttp() {
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

	logrus.Infof("start server at http://localhost:%s", server.Settings.Http.Host)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithField("error", err).Info("startup error")
	}
}
