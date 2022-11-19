package server

import (
	"net/http"
	"time"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	config *config.Config
}

func NewRouter(cfg *config.Config) *Router {
	router := gin.Default()

	return &Router{
		router: router,
		config: cfg,
	}
}
func (r *Router) Run() {
	s := &http.Server{
		Addr:           r.config.Addr(r.config.Server.ApiHost, r.config.Server.ApiPort),
		Handler:        r.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic("server could not run...")
	}
	log.Info("running with addr : %s", s.Addr)
}
