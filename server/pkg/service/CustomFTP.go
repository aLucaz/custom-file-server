package service

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"custom-file-server/shared/util"
	"net"
)

func RegisterChannel(conn net.Conn, config *model.TopicConfig) {
	util.WriteMsgLog(constant.INFO, "New connection request detected!")
	buffer := make([]byte, 1000000)
	bufferLen, err := conn.Read(buffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	request := util.DecodeBufferStrToRegistrationRequest(buffer[:bufferLen])
	for i := 0; i < len(config.Topics); i++ {
		if config.Topics[i].Name == request.ChannelName {
			config.Topics[i].Ports = append(config.Topics[i].Ports, request.Port)
		}
	}
	util.WriteMsgLog(constant.INFO, "New subscription successfully registered on topic "+request.ChannelName)
	util.UpdateConfig(config)
}
