package models

import "encoding/json"

func UnmarshalExecuteReq(data []byte) (ExecuteReq, error) {
	var r ExecuteReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecuteReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ExecuteReq struct {
	ReqID   int64          `json:"reqID"`
	MsgType string         `json:"msgType"`
	Data    ExecuteReqData `json:"data"`
}

type ExecuteReqData struct {
	Path string `json:"path"`
	Args string `json:"args"`
}
