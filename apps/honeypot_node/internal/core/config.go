package core

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"honeypot_node/internal/config"
	"honeypot_node/internal/flags"
	"honeypot_node/internal/global"
	"os"
)

func ReadConfig() *config.Config {
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		logrus.Fatalf("配置文件读取错误 %s", err)
		return nil
	}
	var c = new(config.Config)
	err = yaml.Unmarshal(byteData, &c)
	if err != nil {
		logrus.Fatalf("配置文件配置错误 %s", err)
		return nil
	}
	return c
}
func SetConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		logrus.Errorf("配置序列化失败 %s", err)
		return
	}
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("配置文件写入错误 %s", err)
		return
	}
	logrus.Infof("%s 配置文件更新成功", flags.Options.File)
}
