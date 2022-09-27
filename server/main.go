package main

import (
	"custom-file-server/server/pkg/service"
	"custom-file-server/shared/constant"
	"custom-file-server/shared/util"
	"fmt"
	"net"
	"os"
)

func main() {
	decoded := util.DecodeStr(constant.SERVER_BANNER)
	util.WriteBanner(decoded)
	util.WriteMsgLog(constant.INFO, "Loading configuration...")
	config, err := util.GetConfig()
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
	util.WriteMsgLog(constant.INFO, fmt.Sprintf("Starting server on port %s...", constant.SERVER_PORT))
	for true {
		conn, err := server.Accept()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			os.Exit(1)
		}
		go service.RegisterChannel(conn, config)
	}
}
