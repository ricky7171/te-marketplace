package router

import (
	"github.com/gin-gonic/gin"
	accountpresent "github.com/ricky7171/te-marketplace/internal/modules/account/presentation"
)

type Router struct {
	r              *gin.Engine
	accountHandler *accountpresent.Handler
}

const (
	PostMethod  = "POST"
	GetMethod   = "GET"
	PatchMethod = "PATCH"
)

func NewRouter(r *gin.Engine, accountHandler *accountpresent.Handler) *Router {
	return &Router{
		r:              r,
		accountHandler: accountHandler,
	}
}

func (router *Router) initApi() {
	router.r.POST("/auth/login", func(ctx *gin.Context) {
		router.accountHandler.HandleLogin(ctx)
	})
}

func (router *Router) Run() {
	router.initApi()
	router.r.Run()

}
