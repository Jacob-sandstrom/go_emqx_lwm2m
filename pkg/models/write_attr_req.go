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
	Pmin int64  `json:"pmin,omitempty"`
	Pmax int64  `json:"pmax,omitempty"`
	Gt   int64  `json:"gt,omitempty"`
	Lt   int64  `json:"lt,omitempty"`
	St   int64  `json:"st,omitempty"`
}
