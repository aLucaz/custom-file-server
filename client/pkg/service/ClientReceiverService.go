package service

import (
	"custom-file-server/shared/constant"
	"custom-file-server/shared/service"
	"custom-file-server/shared/util"
	"net"
)

func ProcessMessage(conn net.Conn) {
	buffer := make([]byte, constant.MAX_SIZE_MESSAGE)
	bufferLen, err := conn.Read(buffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		return
	}
	validBuffer := buffer[:bufferLen]
	util.WriteMsgLog(constant.INFO, string(validBuffer))
	headers, err := service.GetHeaders(validBuffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		return
	}
	sendFileRequest, err := service.GetSendFileBody(validBuffer)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		return
	}
	if util.Hash(sendFileRequest.Data) == headers.FingerPrint {
		util.WriteMsgLog(constant.INFO, "File successfully received in the client")
	} else {
		util.WriteMsgLog(constant.ERROR, "The file is corrupted")
	}
}
