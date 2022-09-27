package model

type Topic struct {
	Name  string   `json:"name"`
	Ports []string `json:"port"`
}

type TopicConfig struct {
	Topics []Topic `json:"topics"`
}
