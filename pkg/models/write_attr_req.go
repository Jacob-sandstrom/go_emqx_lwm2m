// Model for lwm2m-emqx write-attr messages

package models

import "encoding/json"

func UnmarshalWriteAttrReq(data []byte) (WriteAttrReq, error) {
	var r WriteAttrReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WriteAttrReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WriteAttrReq struct {
	ReqID   int64         `json:"reqID"`
	MsgType string        `json:"msgType"`
	Data    WriteAttrData `json:"data"`
}

type WriteAttrData struct {
	Path string `json:"path"`
	Pmin string `json:"pmin"`
	Pmax string `json:"pmax"`
	Gt   string `json:"gt"`
	Lt   string `json:"lt"`
	St   string `json:"st"`
}
