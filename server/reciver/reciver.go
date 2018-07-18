package reciver

import (
	"github.com/gin-gonic/gin"
	"upfile/router"
	"upfile/config"
)

func RunServ() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.Run(r)
	r.Run(":" + config.GetPort())
}
