package util

import (
	"bytes"
	"crypto/sha256"
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

func Hash(data string) string {
	hashEncoder := sha256.New()
	hashEncoder.Write([]byte(data))
	hashed := hashEncoder.Sum(nil)
	return string(hashed)
}
