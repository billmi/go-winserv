package config

import (
	"github.com/zpatrick/go-config"
)

var CurrPath string
var envHandle *config.Config
var configFileName = "config"

func GetGconfig(gName string, key string) string {
	_filepath := CurrPath + "/" + configFileName + ".ini"
	iniFile := config.NewINIFile(_filepath)
	envHandle = config.NewConfig([]config.Provider{iniFile})
	result, _ := envHandle.String(gName + "." + key)
	return result
}

func GetPort() string {
	return GetGconfig("uploadFile","port")
}
