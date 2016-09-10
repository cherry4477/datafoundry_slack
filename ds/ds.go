package ds

import "time"

type OpenshiftClient struct {
	host string
	//authUrl string
	oapiUrl string
	kapiUrl string

	namespace   string
	username    string
	password    string
	bearerToken string
}

type Event struct {
	Type   string
	Object object
}

type object struct {
	Kind           string
	ApiVersion     string
	InvolvedObject involvedObject
	Reason         string
	Message        string
	Source         source
	FirstTimestamp time.Time
	LastTimestamp  time.Time
	Count          int
	Type           string
}

type involvedObject struct {
	Kind      string
	Namespace string
	Name      string
	Uid       string
}

type source struct {
	Component string
	Host      string
}

type SlackMsg struct {
	Channel    string `json:"channel"`
	Username   string `json:"username"`
	Text       string `json:"text"`
	Icon_emoji string `json:"icon_emoji"`
}
