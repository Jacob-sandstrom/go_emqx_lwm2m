// Model for lwm2m-emqx read and discover messages

package base_models

import "encoding/json"

func UnmarshalReadDiscoverDeleteReq(data []byte) (ReadDiscoverDeleteReq, error) {
	var r ReadDiscoverDeleteReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReadDiscoverDeleteReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReadDiscoverDeleteReq struct {
	ReqID   int64                  `json:"reqID"`
	MsgType string                 `json:"msgType"`
	Data    ReadDiscoverDeleteData `json:"data"`
}

type ReadDiscoverDeleteData struct {
	Path string `json:"path"`
}
