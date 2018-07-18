package router

import (
	"github.com/gin-gonic/gin"
	"upfile/handlers/upload"
)

func Run(r *gin.Engine) {
	go r.POST("/upload", upload.UploadFile)
}
