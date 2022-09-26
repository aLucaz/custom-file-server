package service

import (
	"custom-file-server/server/pkg/constant"
	"custom-file-server/server/pkg/util"
	"fmt"
	"net"
	"os"
)

func RegisterChannel(conn net.Conn, channelRegistry map[string][]string) {
	channelRegistry["key"] = append(channelRegistry["key"], "value")
	buffer := make([]byte, 1024)
	messageLen, err := conn.Read(buffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	fmt.Println("Received: ", string(buffer[:messageLen]))
	_, err = conn.Write([]byte("Thanks! Got your message:" + string(buffer[:messageLen])))
	err = conn.Close()
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		os.Exit(1)
	}
}
