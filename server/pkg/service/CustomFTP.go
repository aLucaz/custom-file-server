package service

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"custom-file-server/shared/util"
	"encoding/json"
	"net"
)

func RegisterChannel(conn net.Conn, config *model.TopicConfig) {
	util.WriteMsgLog(constant.INFO, "New connection request detected!")
	buffer := make([]byte, 1000000)
	_, err := conn.Read(buffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	request := model.ClientRegistrationRequest{}
	err = json.Unmarshal(buffer, &request)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	for i := 0; i < len(config.Topics); i++ {
		if config.Topics[i].Name == request.ChannelName {
			config.Topics[i].Ports = append(config.Topics[i].Ports, request.Port)
		}
	}
	util.WriteMsgLog(constant.INFO, "New subscription registered on topic "+request.ChannelName)
	util.UpdateConfig(config)
}
