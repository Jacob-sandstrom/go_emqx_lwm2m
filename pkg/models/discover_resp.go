// Model for lwm2m-emqx discover message response

package models

import (
	"encoding/json"
	"fmt"
)

func UnmarshalDiscoverResp(data []byte) (DiscoverResp, error) {
	var r DiscoverResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscoverResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DiscoverResp struct {
	ReqID   int64            `json:"reqID"`
	MsgType string           `json:"msgType"`
	Data    DiscoverRespData `json:"data"`
	Imei    int64            `json:"imei,omitempty"`
	Imsi    int64            `json:"imsi,omitempty"`
}

type DiscoverRespData struct {
	ReqPath string   `json:"reqPath"`
	Content []string `json:"content"`
	CodeMsg string   `json:"codeMsg"`
	Code    string   `json:"code"`
}

type IDiscoverResp interface {
	Print()
}

func (x *DiscoverResp) Print() {
	fmt.Printf("ReqID: %v\n", x.ReqID)
	fmt.Printf("MsgType: %v\n", x.MsgType)
	if x.Imei != 0 {
		fmt.Printf("Imei: %v\n", x.Imei)
	}
	if x.Imsi != 0 {
		fmt.Printf("Imsi: %v\n", x.Imsi)
	}
	fmt.Printf("Data:\n")
	fmt.Printf("	ReqPath: %v\n", x.Data.ReqPath)
	fmt.Printf("	CodeMsg: %v\n", x.Data.CodeMsg)
	fmt.Printf("	Code: %v\n", x.Data.Code)

	fmt.Print("	Content:\n")
	for _, v := range x.Data.Content {
		fmt.Printf("		%v\n", v)
	}
	// fmt.Printf("	Content: %v\n", x.DiscoverRespData.Content)
}
