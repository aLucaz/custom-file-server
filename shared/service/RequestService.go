package service

import (
	"bufio"
	"custom-file-server/shared/constant"
	"custom-file-server/shared/model"
	"custom-file-server/shared/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateClientRegistrationRequest(address string, channelName string) []byte {
	body := model.ClientRegistrationRequest{}
	body.ChannelName = channelName
	addressList := strings.Split(address, ":")
	body.Port = addressList[1]
	headers := model.Header{}
	headers.Operation = constant.REGISTER_CLIENT
	jsonBody, err := json.Marshal(body)
	jsonHeaders, err := json.Marshal(headers)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		os.Exit(1)
	}
	requestStr := constant.HEADER_TITLE + string(jsonHeaders[:]) + constant.REQ_SEP + constant.BODY_TITLE + string(jsonBody[:])
	util.WriteMsgLog(constant.INFO, fmt.Sprintf("Request created on port %s...", body.Port))
	return []byte(requestStr)
}

func CreateSendFileRequest(fileName string, channelName string) []byte {
	fileSeparator := fmt.Sprintf("%c", filepath.Separator)
	path := constant.TEST_FILES_DIRECTORY + fileSeparator + fileName
	file, err := os.Open(path)
	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		os.Exit(1)
	}
	encoded := base64.StdEncoding.EncodeToString(content)
	hashed := util.Hash(encoded)
	headers := model.Header{}
	headers.Operation = constant.SEND_FILE
	headers.FingerPrint = hashed
	body := model.SendFileRequest{}
	body.Data = encoded
	body.ChannelName = channelName
	jsonBody, err := json.Marshal(body)
	jsonHeaders, err := json.Marshal(headers)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		os.Exit(1)
	}
	requestStr := constant.HEADER_TITLE + string(jsonHeaders[:]) + constant.REQ_SEP + constant.BODY_TITLE + string(jsonBody[:])
	return []byte(requestStr)
}

func SendFileFromServer(body model.SendFileRequest, headers model.Header) []byte {
	jsonBody, err := json.Marshal(body)
	jsonHeaders, err := json.Marshal(headers)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
		os.Exit(1)
	}
	requestStr := constant.HEADER_TITLE + string(jsonHeaders[:]) + constant.REQ_SEP + constant.BODY_TITLE + string(jsonBody[:])
	return []byte(requestStr)
}

func GetHeaders(buffer []byte) model.Header {
	bufferStr := string(buffer)
	request := strings.Split(bufferStr, constant.REQ_SEP)
	requestHeader := strings.Split(request[0], constant.HEADER_TITLE)[1]
	headers := model.Header{}
	err := json.Unmarshal([]byte(requestHeader), &headers)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	return headers
}

func GetClientRegistrationBody(buffer []byte) model.ClientRegistrationRequest {
	bufferStr := string(buffer)
	request := strings.Split(bufferStr, constant.REQ_SEP)
	requestBody := strings.Split(request[1], constant.BODY_TITLE)[1]
	clientRegistrationRequest := model.ClientRegistrationRequest{}
	err := json.Unmarshal([]byte(requestBody), &clientRegistrationRequest)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	return clientRegistrationRequest
}

func GetSendFileBody(buffer []byte) model.SendFileRequest {
	bufferStr := string(buffer)
	request := strings.Split(bufferStr, constant.REQ_SEP)
	requestBody := strings.Split(request[1], constant.BODY_TITLE)[1]
	sendFileRequest := model.SendFileRequest{}
	err := json.Unmarshal([]byte(requestBody), &sendFileRequest)
	if err != nil {
		util.WriteMsgLog(constant.ERROR, err.Error())
	}
	return sendFileRequest
}
