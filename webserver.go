package main

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func runWebServer() {
	client := gin.Default()
	client.Static("/assets", "./resources/assets")
	gv := goview.Config{
		DisableCache: true,
		Extension:    ".gohtml",
		Root:         "./resources/templates/",
		Master:       "layouts/master",
		Funcs:        nil,
	}
	client.HTMLRender = ginview.New(gv)

	webHandlers(client)
	runStr := fmt.Sprintf("%s:%d", Config.GetString("web.address"), Config.GetInt("web.port"))
	err := client.Run(runStr)
	if err != nil {
		Log.Panicf("fatal error config file: %s", err.Error())
	}
}

func webHandlers(router *gin.Engine) {
	router.GET("/", webGetHome)
	router.GET("/getinfo", webGetInfo)
}

func webGetInfo(ctx *gin.Context) {
	res, err := getInfoOverApi()
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}
func webGetHome(ctx *gin.Context) {
	ctx.HTML(200, "home.gohtml", nil)
}
