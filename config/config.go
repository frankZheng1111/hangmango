package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type GameSetting struct {
	Hp int `yaml: hp`
}

var GameSettingConfig GameSetting
var GO_ENV string

var VALID_ENV map[string]bool = map[string]bool{
	"dev": true,
}

func init() {
	GO_ENV := os.Getenv(GO_ENV)
	if _, ok := VALID_ENV[GO_ENV]; !ok {
		GO_ENV = "dev"
	}
	parseYmlConfigByEnv(GO_ENV, &GameSettingConfig)
}

// 项目根目录路径
//
func parseYmlConfigByEnv(env string, configPtr *GameSetting) {
	path := "config/" + env + ".yml"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic("missing config file")
	}
	err = yaml.Unmarshal(content, configPtr)
	if err != nil {
		panic("parse config file failed")
	}
	return
}
