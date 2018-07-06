package router

import (
	"github.com/gin-gonic/gin"
	"github.com/NickDu2009/apiserver/router/middleware"
	"net/http"
	"github.com/NickDu2009/apiserver/hanlder/sd"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

		g.Use(gin.Recovery())
		g.Use(middleware.NoCache)
		g.Use(middleware.Options)
		g.Use(middleware.Secure)
		g.Use(mw...)

		g.NoRoute(func(c *gin.Context) {
			c.String(http.StatusNotFound, "The incorrect API route")
		})

		svcd := g.Group("/sd")
		{
			svcd.GET("/health", sd.HealthCheck)
			svcd.GET("/disk", sd.DiskCheck)
			svcd.GET("/cpu", sd.CPUChecke)
			svcd.GET("/ram", sd.RAMCheck)
		}
		return g
}