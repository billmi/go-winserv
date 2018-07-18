package helper

import (
	"upfile/utils/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	userToken = "u:token:%d"
	userLogin = "u:login:%d"
	token     = "token"
	memberId  = "member_id"
)

func GetRequestHeader(c *gin.Context) map[string]interface{} {
	headers := strings.TrimSpace(c.DefaultPostForm("headers", ""))
	ResultVo := map[string]interface{}{}
	if headers == "" {
		return ResultVo
	}
	return json.StringToJson(headers)
}

func GetUserHeader(userId int, tok string) map[string]interface{} {
	headers := map[string]interface{}{}
	headers[token] = tok
	headers[memberId] = strconv.Itoa(userId)
	return headers
}
