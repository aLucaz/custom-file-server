package main

import (
	"custom-file-server/client/pkg/service"
	"custom-file-server/shared/constant"
	sharedService "custom-file-server/shared/service"
	"custom-file-server/shared/util"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial(constant.SERVER_TYPE, constant.SERVER_HOST+":"+constant.SERVER_PORT)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
	}(connection)
	decoded := util.DecodeBase64Str(constant.CLIENT_BANNER)
	util.WriteBanner(decoded)
	util.WriteMsgLog(constant.INFO, "Starting client...")
	arguments := os.Args
	if arguments[1] == constant.RECEIVE_MODE {
		util.WriteMsgLog(constant.INFO, "Receiver mode activated")
		server, err := net.Listen(constant.SERVER_TYPE, constant.CLIENT_HOST+":")
		request := sharedService.CreateClientRegistrationRequest(server.Addr().String(), arguments[2])
		_, err = connection.Write(request)
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
		util.WriteMsgLog(constant.INFO, fmt.Sprintf("Starting watcher..."))
		for true {
			_, err := server.Accept()
			if err != nil {
				util.WriteMsgLog(constant.ERROR, err.Error())
				os.Exit(1)
			}
			service.ProcessMessage(connection)
		}
	} else if arguments[1] == constant.SEND_MODE {
		util.WriteMsgLog(constant.INFO, "Sender mode activated")
		request := sharedService.CreateSendFileRequest(arguments[2], arguments[3])
		util.WriteMsgLog(constant.INFO, fmt.Sprintf("Sending file to topic named: %s", arguments[3]))
		_, err = connection.Write(request)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			os.Exit(1)
		}
	}
}
