package main

import (
	"github.com/gin-gonic/gin"
	"wm-take-out/global"
	"wm-take-out/initialize"
)

func main() {
	router := initialize.GlobalInit()
	gin.SetMode(global.Config.Server.Level)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
