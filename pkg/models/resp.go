// Model for lwm2m-emqx read message response

package models

import "encoding/json"

func UnmarshalResp(data []byte) (Resp, error) {
	var r Resp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Resp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Resp struct {
	ReqID   int64    `json:"reqID"`
	MsgType string   `json:"msgType"`
	Data    RespData `json:"data"`
	Imei    int64    `json:"imei,omitempty"`
	Imsi    int64    `json:"imsi,omitempty"`
}

type RespData struct {
	ReqPath string      `json:"reqPath,omitempty"`
	Content interface{} `json:"content,omitempty"`
	CodeMsg string      `json:"codeMsg"`
	Code    string      `json:"code"`
}
