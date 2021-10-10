package proto

import "encoding/json"

type Request struct {
	Command string           `json:"command"`
	Data    *json.RawMessage `json:"data"`
}
type Response struct {
	Status string           `json:"status"`
	Data   *json.RawMessage `json:"data"`
}

type Events struct {
	Eventmessage string `json:"mes"`
	Time         string `json:"time"`
}
