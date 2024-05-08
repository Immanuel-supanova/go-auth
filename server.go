package goauth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/go-auth/routes"
)

func userRouter() http.Handler {
	debug := os.Getenv("DEBUG")

	if debug == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routes.UserRoutes(r)

	return r
}

var UserServer = &http.Server{
	Addr:         ":3000",
	Handler:      userRouter(),
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}
