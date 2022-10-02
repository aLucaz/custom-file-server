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
		util.WriteMsgLog(constant.ERROR, "Error trying to connect to server")
		util.WriteMsgLog(constant.ERROR, err.Error())
		return
	}
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, "Error trying to close Dial connection with server")
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
	}(connection)
	util.WriteBanner(constant.CLIENT_BANNER)
	util.WriteMsgLog(constant.INFO, "Starting client...")
	arguments := os.Args
	if arguments[1] == constant.RECEIVE_MODE {
		util.WriteMsgLog(constant.INFO, "Receiver mode activated")
		server, err := net.Listen(constant.SERVER_TYPE, constant.CLIENT_HOST+":")
		if err != nil {
			util.WriteMsgLog(constant.ERROR, "Couldn't stablish a connection")
			util.WriteMsgLog(constant.ERROR, err.Error())
			panic(err)
		}
		defer func(server net.Listener) {
			err := server.Close()
			if err != nil {
				util.WriteMsgLog(constant.ERROR, err.Error())
			}
		}(server)
		request, err := sharedService.CreateClientRegistrationRequest(server.Addr().String(), arguments[2])
		if err != nil {
			util.WriteMsgLog(constant.ERROR, "Error creating the request, try again")
			util.WriteMsgLog(constant.ERROR, err.Error())
			// TODO unsubscribe before panic
			panic(err)
		}
		_, err = connection.Write(request)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			panic(err)
		}
		util.WriteMsgLog(constant.INFO, fmt.Sprintf("Starting watcher..."))
		for true {
			conn, err := server.Accept()
			if err != nil {
				util.WriteMsgLog(constant.ERROR, err.Error())
				panic(err)
			}
			go service.ProcessMessage(conn)
		}
	} else if arguments[1] == constant.SEND_MODE {
		util.WriteMsgLog(constant.INFO, "Sender mode activated")
		request := sharedService.CreateSendFileRequest(arguments[2], arguments[3])
		util.WriteMsgLog(constant.INFO, fmt.Sprintf("File sent on topic named: %s", arguments[3]))
		_, err = connection.Write(request)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			os.Exit(1)
		}
	}
}
