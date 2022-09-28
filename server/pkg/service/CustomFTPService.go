package service

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"custom-file-server/shared/service"
	"custom-file-server/shared/util"
	"net"
)

func RegisterChannel(request model.ClientRegistrationRequest, config *model.TopicConfig) {
	util.WriteMsgLog(constant.INFO, "New connection request detected!")
	for i := 0; i < len(config.Topics); i++ {
		if config.Topics[i].Name == request.ChannelName {
			config.Topics[i].Ports = append(config.Topics[i].Ports, request.Port)
		}
	}
	util.WriteMsgLog(constant.INFO, "New subscription successfully registered on topic "+request.ChannelName)
	util.UpdateConfig(config)
}

func SendFile(sendFileRequest model.SendFileRequest, headers model.Header, config *model.TopicConfig) {
	util.WriteMsgLog(constant.INFO, "Broadcasting file...")
	channelName := sendFileRequest.ChannelName
	for _, topic := range config.Topics {
		if topic.Name == channelName {
			for _, port := range topic.Ports {
				connection, err := net.Dial(constant.SERVER_TYPE, constant.CLIENT_HOST+":"+port)
				if err != nil {
					util.WriteMsgLog(constant.ERROR, err.Error())
				}
				request := service.CreateSendFileRequestFromServer(sendFileRequest, headers)
				_, err = connection.Write(request)
				err = connection.Close()
				if err != nil {
					util.WriteMsgLog(constant.ERROR, err.Error())
				}
			}
		}
	}
}

func ProcessRequest(conn net.Conn, config *model.TopicConfig) {
	buffer := make([]byte, constant.MAX_SIZE_MESSAGE)
	bufferLen, err := conn.Read(buffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	validBuffer := buffer[:bufferLen]
	headers := service.GetHeaders(validBuffer)
	if headers.Operation == constant.REGISTER_CLIENT {
		clientRegistrationRequest := service.GetClientRegistrationBody(validBuffer)
		RegisterChannel(clientRegistrationRequest, config)
	} else if headers.Operation == constant.SEND_FILE {
		sendFileRequest := service.GetSendFileBody(validBuffer)
		if util.Hash(sendFileRequest.Data) == headers.FingerPrint {
			util.WriteMsgLog(constant.INFO, "File successfully received in the server")
			SendFile(sendFileRequest, headers, config)
			util.WriteMsgLog(constant.INFO, "All files has been sended")
		} else {
			util.WriteMsgLog(constant.ERROR, "The file is corrupted")
		}
	}
}
