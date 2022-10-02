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

func SendFileToEveryClientOnChannel(ports []string, headers model.Header, sendFileRequest model.SendFileRequest) {
	for _, port := range ports {
		connection, err := net.Dial(constant.SERVER_TYPE, constant.CLIENT_HOST+":"+port)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
		request, err := service.CreateSendFileRequestFromServer(sendFileRequest, headers)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
		_, err = connection.Write(request)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
		err = connection.Close()
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
		}
	}
}

func SendFileToChannel(sendFileRequest model.SendFileRequest, headers model.Header, config *model.TopicConfig) {
	util.WriteMsgLog(constant.INFO, "Broadcasting file...")
	channelName := sendFileRequest.ChannelName
	for _, topic := range config.Topics {
		if topic.Name == channelName {
			SendFileToEveryClientOnChannel(topic.Ports, headers, sendFileRequest)
		}
	}
}

func ProcessRequest(conn net.Conn, config *model.TopicConfig) {
	buffer := make([]byte, constant.MAX_SIZE_MESSAGE)
	bufferLen, err := conn.Read(buffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		return
	}
	validBuffer := buffer[:bufferLen]
	headers, err := service.GetHeaders(validBuffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		return
	}
	if headers.Operation == constant.REGISTER_CLIENT {
		clientRegistrationRequest, err := service.GetClientRegistrationBody(validBuffer)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			return
		}
		RegisterChannel(clientRegistrationRequest, config)
	} else if headers.Operation == constant.SEND_FILE {
		sendFileRequest, err := service.GetSendFileBody(validBuffer)
		if err != nil {
			util.WriteMsgLog(constant.ERROR, err.Error())
			return
		}
		if util.Hash(sendFileRequest.Data) == headers.FingerPrint {
			util.WriteMsgLog(constant.INFO, "File successfully received in the server")
			SendFileToChannel(sendFileRequest, headers, config)
			util.WriteMsgLog(constant.INFO, "All files has been sended")
		} else {
			util.WriteMsgLog(constant.ERROR, "The file is corrupted")
		}
	}
}
