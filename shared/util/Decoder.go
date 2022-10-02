package util

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"encoding/base64"
	"encoding/json"
	"strings"
)

func DecodeBase64Str(str string) (string, error) {
	var decoded, err = base64.StdEncoding.DecodeString(str)
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
		return "", err
	}
	return string(decoded), nil
}

func DecodeBufferStrToRegistrationRequest(buffer []byte) model.ClientRegistrationRequest {
	index := strings.Index(string(buffer), "{")
	request := model.ClientRegistrationRequest{}
	err := json.Unmarshal(buffer[index:], &request)
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
	return request
}
