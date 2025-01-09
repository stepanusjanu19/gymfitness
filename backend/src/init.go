package src

import (
	"backend/src/config"
	"github.com/gin-gonic/gin"
)

func Run()  {
	config.InitDB()
	defer config.CloseDB()

	route := gin.Default()

	route.Use()

	route.Run(":5001")
}