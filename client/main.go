package main

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"custom-file-server/shared/util"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
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
	//arguments := os.Args
	arguments := [2]string{"receive", "channel-1"}
	if arguments[0] == constant.RECEIVE_MODE {
		util.WriteMsgLog(constant.INFO, "Reveiver mode activated")
		// initializating receiver
		server, err := net.Listen(constant.SERVER_TYPE, constant.CLIENT_HOST+":")
		request := model.ClientRegistrationRequest{}
		request.ChannelName = arguments[1]
		address := strings.Split(server.Addr().String(), ":")
		request.Port = address[1]
		jsonRequest, err := json.Marshal(request)
		encodedRequest := util.EncodeToBytes(jsonRequest)
		//util.WriteMsgLog(constant.INFO, string(encodedRequest))
		_, err = connection.Write(encodedRequest)
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
		util.WriteMsgLog(constant.INFO, fmt.Sprintf("Starting client receiver on port %s...", address[1]))
		for true {
			_, err := server.Accept()
			if err != nil {
				util.WriteMsgLog(constant.ERROR, err.Error())
				os.Exit(1)
			}
		}
	}
}
