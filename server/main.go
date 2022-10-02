package main

import (
	"custom-file-server/server/pkg/service"
	"custom-file-server/shared/constant"
	"custom-file-server/shared/util"
	"fmt"
	"net"
)

func main() {
	util.WriteBanner(constant.SERVER_BANNER)
	util.WriteMsgLog(constant.INFO, "Loading configuration...")
	config, err := util.GetConfig()
	if err != nil {
		util.WriteMsgLog(constant.ERROR, "Error trying to load configurations")
		util.WriteMsgLog(constant.ERROR, err.Error())
		panic(err)
	}
	server, err := net.Listen(constant.SERVER_TYPE, constant.SERVER_HOST+":"+constant.SERVER_PORT)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		panic(err)
	}
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
	}(server)
	util.WriteMsgLog(constant.INFO, fmt.Sprintf("Starting server on port %s...", constant.SERVER_PORT))
	for true {
		conn, err := server.Accept()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			panic(err)
		}
		go service.ProcessRequest(conn, config)
	}
}
