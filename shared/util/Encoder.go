package util

import (
	"bytes"
	"encoding/gob"
	"log"
)

func EncodeToBytes(structure interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(structure)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
