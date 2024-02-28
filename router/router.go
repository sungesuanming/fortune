package router

import (
	"fortune/handler/color"
	"fortune/handler/sd"
	"fortune/router/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//Middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Option)
	g.Use(cors.Default())
	g.Use(middleware.Secure)
	g.Use(mw...)
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://www.prawards.com.cn/","http://prawards.com.cn/","http://chenge.duocaishenghuo.cn/","http://public.duocaishenghuo.cn/"}
	//404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	//The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	c := g.Group("/color")
	{
		c.POST("/test", color.ColorTest)
		c.POST("/today", color.TodayColor)
	}
	return g
}
