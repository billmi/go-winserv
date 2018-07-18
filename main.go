package main

import (
	"upfile/config"
	"runtime"
	"path/filepath"
	"os/exec"
	"os"
	"github.com/kardianos/service"
	"upfile/server/reciver"
	"github.com/jander/golog/logger"
)

func getLinuxConfiPath() string {
	file, _ := exec.LookPath(os.Args[0])
	ApplicationPath, _ := filepath.Abs(file)
	ApplicationDir, _ := filepath.Split(ApplicationPath)
	return ApplicationDir + "\\"
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

//业务服务
func (p *program) run() {
	reciver.RunServ()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func init() {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		config.CurrPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		config.CurrPath = getLinuxConfiPath()
	}
}

func main() {
	svcConfig := &service.Config{
		Name:        "AIReciver",
		DisplayName: "AIID AIReciver",
		Description: "智影医疗高速影响图片上传服务组件", //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Fatal(err)
	}

	if err != nil {
		logger.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "-i" {
			s.Install()
			logger.Println("服务安装成功!")
			return
		}

		if os.Args[1] == "remove" {
			s.Stop()
			s.Uninstall()
			logger.Println("服务卸载成功!")
			return
		}

		if os.Args[1] == "start" {
			s.Start()
			logger.Println("服务启动成功!")
			return
		}

		if os.Args[1] == "restart" {
			s.Restart()
			logger.Println("服务重启成功!")
			return
		}
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
