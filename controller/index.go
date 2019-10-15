package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnli/go-base/controller/business/users"
	"github.com/jhonnli/go-base/controller/common"
	"net/http"
)

func Init(engine *gin.Engine) {
	engine.GET("/", index)
	engine.Any("/health", health)
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, common.Result{Code: "not found", Message: "Page not found"})
	})
	users.Init(engine)
}

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "I'm OK")
}

func health(ctx *gin.Context) {
	result := make(map[string]string)
	result["status"] = "service status is up"
	ctx.JSON(http.StatusOK, result)
}
