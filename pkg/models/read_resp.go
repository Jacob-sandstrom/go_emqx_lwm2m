// Model for lwm2m-emqx read message response

package models

import "encoding/json"

func UnmarshalReadResp(data []byte) (ReadResp, error) {
	var r ReadResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReadResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReadResp struct {
	ReqID   int64        `json:"reqID"`
	MsgType string       `json:"msgType"`
	Data    ReadRespData `json:"data"`
	Imei    int64        `json:"imei,omitempty"`
	Imsi    int64        `json:"imsi,omitempty"`
}

type ReadRespData struct {
	ReqPath string            `json:"reqPath"`
	Content []ReadRespContent `json:"content,omitempty"`
	CodeMsg string            `json:"codeMsg"`
	Code    string            `json:"code"`
}

type ReadRespContent struct {
	Value int64  `json:"value"`
	Path  string `json:"path"`
}
