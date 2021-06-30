package models

import "encoding/json"

func UnmarshalObserveReq(data []byte) (ObserveReq, error) {
	var r ObserveReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ObserveReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ObserveReq struct {
	ReqID   int64          `json:"reqID"`
	MsgType string         `json:"msgType"`
	Data    ObserveReqData `json:"data"`
}

type ObserveReqData struct {
	Path string `json:"path"`
}
