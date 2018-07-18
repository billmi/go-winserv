package hbase64

import "encoding/base64"

func DeBase64ToString(data string) (string,bool) {
	res,err := base64.StdEncoding.DecodeString(data)
	if err != nil{
		return "",false
	}
	return string(res),true
}