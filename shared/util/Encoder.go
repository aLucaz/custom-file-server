package util

import (
	"bytes"
	"custom-file-server/shared/constant"
	"encoding/gob"
)

func EncodeToBytes(structure interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(structure)
	if err != nil {
		WriteMsgLog(constant.ERROR, err.Error())
	}
	return buf.Bytes()
}
