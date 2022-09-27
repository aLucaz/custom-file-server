package model

type ClientRegistrationRequest struct {
	ChannelName string `json:"channelName"`
	Port        string `json:"port"`
}
