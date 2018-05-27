package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type GameSetting struct {
	DictionaryName string `yaml:"dictionary_name"`
	Hp             int    `yaml:"hp"`
}

const PORJECT_NAME string = "hangmango"
const FILE_NAME string = "config.go"

var FILE_NAME2 = 3
var FILE_NAME3 int = 3

var GameSettingConfig GameSetting
var GO_ENV string
var CONFIG_FOLDER_PATH string

var VALID_ENV map[string]bool = map[string]bool{
	"dev":  true,
	"test": true,
}

func init() {
	GO_ENV := os.Getenv("GO_ENV")
	if _, ok := VALID_ENV[GO_ENV]; !ok {
		GO_ENV = "dev"
	}
	initConfigFileFolderPath()
	parseYmlConfigByEnv(GO_ENV, &GameSettingConfig)
}

func initConfigFileFolderPath() {
	configPaths := []string{}
	goPathStr := os.Getenv("GOPATH")
	cmdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPaths = append(configPaths, filepath.Join(cmdPath, "config"))
	for _, goPath := range strings.Split(goPathStr, ":") {
		configPaths = append(configPaths, filepath.Join(goPath, "src", PORJECT_NAME, "config"))
	}
	for _, path := range configPaths {
		_, err = ioutil.ReadFile(filepath.Join(path, FILE_NAME))
		if err == nil {
			CONFIG_FOLDER_PATH = path
			break
		}
	}
	if CONFIG_FOLDER_PATH == "" {
		panic(fmt.Sprintf("missing config folder in these paths: %s", configPaths))
	}
}

// 项目根目录路径
//
func parseYmlConfigByEnv(env string, configPtr *GameSetting) {
	filename := env + ".yml"
	content, err := ioutil.ReadFile(filepath.Join(CONFIG_FOLDER_PATH, filename))
	if err != nil {
		panic(fmt.Sprintf("missing config file in %s", CONFIG_FOLDER_PATH))
	}
	err = yaml.Unmarshal(content, configPtr)
	if err != nil {
		panic("parse config file failed")
	}
	return
}
