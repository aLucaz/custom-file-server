package util

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetCurrentDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetConfigPath() (string, error) {
	currentDir, err := GetCurrentDir()
	if err != nil {
		return "", err
	}
	return currentDir + constant.CONFIG_PATH, nil
}

func GetConfig() (*model.TopicConfig, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := &model.TopicConfig{}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func UpdateConfig(config *model.TopicConfig) {
	path, err := GetConfigPath()
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
}
