package server

import (
	"net/http"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/interfaces/restapi/handler"
	"github.com/finnpn/workout-tracker/interfaces/restapi/middleware"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	config *config.Config
}

func NewRouter(cfg *config.Config, handler *handler.Handler) *Router {
	router := gin.New()
	router.Use(middleware.HandleErrors())
	router.Use(gin.Recovery())
	routerNotAuth := router.Group("/auth")
	handler.ConfigureNotAuth(routerNotAuth)

	return &Router{
		router: router,
		config: cfg,
	}
}
func (r *Router) Run() {
	s := &http.Server{
		Addr:    r.config.Addr(r.config.Server.ApiHost, r.config.Server.ApiPort),
		Handler: r.router,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic("server could not run...")
	}
	log.Info("running with addr : %s", s.Addr)
}
