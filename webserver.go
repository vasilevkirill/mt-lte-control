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
	gin.SetMode(gin.DebugMode)
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

	res := gin.H{}
	res["lteInfo"] = R.readStatusMapKey("lteInfo")
	res["script"] = R.readStatusMapKey("script")
	res["simSlot"] = R.readStatusMapKey("simSlot")

	ctx.JSON(200, res)
}
func webGetHome(ctx *gin.Context) {
	res := gin.H{}
	res["control"] = Config.Get("control")
	ctx.HTML(200, "home.gohtml", res)
}
