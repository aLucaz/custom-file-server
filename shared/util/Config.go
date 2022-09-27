package util

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
	return path
}

func GetConfigPath() string {
	return GetCurrentDir() + constant.CONFIG_PATH
}

func GetConfig() (*model.TopicConfig, error) {
	path := GetConfigPath()
	file, err := ioutil.ReadFile(path)
	conf := &model.TopicConfig{}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func UpdateConfig(config *model.TopicConfig) {
	path := GetConfigPath()
	file, _ := json.MarshalIndent(config, "", "  ")
	err := ioutil.WriteFile(path, file, 0644)
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
}
