package model

type SendFileRequest struct {
	Data        string `json:"data"`
	ChannelName string `json:"channelName"`
}
