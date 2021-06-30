// Model for lwm2m-emqx register and update messages

package base_models

import "encoding/json"

func UnmarshalRegisterUpdateResp(data []byte) (RegisterUpdateResp, error) {
	var r RegisterUpdateResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RegisterUpdateResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RegisterUpdateResp struct {
	MsgType string             `json:"msgType"`
	Data    RegisterUpdateData `json:"data"`
}

type RegisterUpdateData struct {
	Ep            string   `json:"ep"`
	Lt            int64    `json:"lt"`
	SMS           string   `json:"sms,omitempty"`
	Lwm2M         string   `json:"lwm2m"`
	B             string   `json:"b,omitempty"`
	AlternatePath string   `json:"alternatePath"`
	ObjectList    []string `json:"objectList"`
}
