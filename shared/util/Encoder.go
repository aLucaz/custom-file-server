package util

import (
	"bytes"
	"crypto/sha256"
	"custom-file-server/shared/constant"
	"encoding/gob"
	"fmt"
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
	dataByte := []byte(data)
	sum := sha256.Sum256(dataByte)
	return fmt.Sprintf("%x", string(sum[:]))
}
