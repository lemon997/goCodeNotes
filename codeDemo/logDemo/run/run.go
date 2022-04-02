package run

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"logDemo/logtest"
)

func b() error {
	for i := 0; i < 10; i++ {
		if i == 5 {
			return errors.New("run.b err")
		}
	}
	return nil
}

func a() {
	err := b()
	if err != nil {
		logtest.Instance().Error(err)
	}
	panic("123")
}

func Run() {
	defer func() {
		if p := recover(); p != nil {
			logtest.Instance().Logf(logrus.FatalLevel, "RUn panic = %v", p)
		}
		fmt.Println("Run完成")
	}()
	a()
}
