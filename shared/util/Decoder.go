package util

import (
	"custom-file-server/shared/constant"
	"encoding/base64"
)

func DecodeStr(str string) string {
	var decoded, err = base64.StdEncoding.DecodeString(str)
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
		return ""
	}
	return string(decoded)
}
