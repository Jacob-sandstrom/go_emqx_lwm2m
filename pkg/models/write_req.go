// Model for lwm2m-emqx write messages

package models

import "encoding/json"

func UnmarshalWriteReq(data []byte) (WriteReq, error) {
	var r WriteReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WriteReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WriteReq struct {
	ReqID   int64     `json:"reqID"`
	MsgType string    `json:"msgType"`
	Data    WriteData `json:"data"`
}

type WriteData struct {
	Path  string      `json:"path"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
