// Model for lwm2m-emqx "write", "write-attr", "execute", "create", and "delete" message responses

package base_models

import "encoding/json"

func UnmarshalStandardResp(data []byte) (StandardResp, error) {
	var r StandardResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StandardResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StandardResp struct {
	ReqID   int64            `json:"reqID"`
	MsgType string           `json:"msgType"`
	Data    StandardRespData `json:"data"`
	Imei    int64            `json:"imei,omitempty"`
	Imsi    int64            `json:"imsi,omitempty"`
}

type StandardRespData struct {
	ReqPath string `json:"reqPath"`
	CodeMsg string `json:"codeMsg"`
	Code    string `json:"code"`
}
