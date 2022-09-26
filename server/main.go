package main

import (
	"custom-file-server/server/pkg/constant"
	"custom-file-server/server/pkg/service"
	"custom-file-server/server/pkg/util"
	"net"
	"os"
)

func main() {
	server, err := net.Listen(constant.SERVER_TYPE, constant.SERVER_HOST+":"+constant.SERVER_PORT)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		os.Exit(1)
	}
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			os.Exit(1)
		}
	}(server)
	decoded := util.DecodeStr(constant.BANNER)
	util.WriteBanner(decoded)
	util.WriteMsgLog(constant.INFO, "Starting server...")
	channelRegistry := make(map[string][]string)
	for true {
		conn, err := server.Accept()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			os.Exit(1)
		}
		go service.RegisterChannel(conn, channelRegistry)
	}
}
