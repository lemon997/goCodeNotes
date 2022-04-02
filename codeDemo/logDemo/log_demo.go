package main

import (
	"fmt"
	"github.com/kardianos/service"
	"logDemo/logtest"
	"logDemo/run"
	"os"
	"time"
)

//windows使用步骤：logDemo.exe install --> logDemo.exe start --> logDemo.exe stop --> logDemo.exe uninstall

func main() {
	//DisplayName是显示名称
	//Name是服务器名

	srvConfig := &service.Config{
		Name:        "logServer",
		DisplayName: "logServer123456",
		Description: "logServer运行",
	}
	prg := &program{}
	s, err := service.New(prg, srvConfig)
	if err != nil {
		fmt.Println(err)
	}
	if len(os.Args) > 1 {
		serviceAction := os.Args[1]
		switch serviceAction {
		case "install":
			err := s.Install()
			if err != nil {
				fmt.Println("安装服务失败: ", err.Error())
			} else {
				fmt.Println("安装服务成功")
			}
			return
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				fmt.Println("卸载服务失败: ", err.Error())
			} else {
				fmt.Println("卸载服务成功")
			}
			return
		case "start":
			err := s.Start()
			if err != nil {
				fmt.Println("运行服务失败: ", err.Error())
			} else {
				fmt.Println("运行服务成功")
			}
			return
		case "stop":
			err := s.Stop()
			if err != nil {
				fmt.Println("停止服务失败: ", err.Error())
			} else {
				fmt.Println("停止服务成功")
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
}

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *program) run() error {
	// 具体的服务实现
	log := logtest.Instance()
	ticker := time.NewTicker(5 * time.Second)
	go run.Run()
	for {
		select {
		case tm := <-ticker.C:
			log.Infof("日志正常输出, tm = %v", tm)
		case <-p.exit:
			ticker.Stop()
			return nil
		}
	}
	return nil
}
func (p *program) Stop(s service.Service) error {
	log := logtest.Instance()
	log.Info("I'm Stopping!")
	close(p.exit)
	return nil
}
