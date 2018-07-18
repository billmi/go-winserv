package upload

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"upfile/utils/respo"
	"strings"
	"mime/multipart"
	"upfile/config"
	"upfile/utils/filecheck"
)

func UploadFile(c *gin.Context) {
	filename := strings.TrimSpace(c.DefaultPostForm("filename", ""))
	key := strings.TrimSpace(c.DefaultPostForm("key", ""))
	if filename == "" {
		respo.HttpErr(c, "Filename is empty!", nil, "")
		return
	}
	var skey = config.GetGconfig("uploadFile", "key")
	var _saveDir = config.GetGconfig("uploadFile", "destdir")
	if key == "" || key != skey {
		respo.HttpErr(c, "Key is err!", nil, "")
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		respo.HttpErr(c, "File is err!", nil, "")
		return
	}
	savePath := _saveDir + "\\" + filename
	if res1, _ := filecheck.PathExists(_saveDir); res1 == false {
		respo.HttpErr(c, "The dir is not exist ! Please create a directory for storage !", nil, "")
		return
	}
	_write := make(chan int64)
	go func(file multipart.File, savePath string, write chan int64) {
		out, _ := os.Create(savePath)
		defer out.Close()
		writen, _ := io.Copy(out, file)
		_write <- writen
	}(file, savePath, _write)
	var res = map[string]interface{}{}
	res["filename"] = filename
	res["savePath"] = savePath
	res["fileSize"] = <-_write
	respo.HttpHSucc(c, "Success !", nil, res, "")
	return
}
