package main

import (
	"lensman/app/config"
	util "lensman/app/controller"
	"lensman/app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Init()
}

func main() {
	router := gin.Default()
	middleware.Intercept(router)
	util.Register(router)
	router.Run()
}
