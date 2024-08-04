package router

import (
	"net/http"

	"github.com/loganrk/go-router/gin"
)

type Router interface {
	RegisterRoute(method, path string, handlerFunc http.HandlerFunc)
	StartServer(port string) error
	UseBefore(middlewares ...http.HandlerFunc)
}

func New() Router {
	return gin.New()
}
