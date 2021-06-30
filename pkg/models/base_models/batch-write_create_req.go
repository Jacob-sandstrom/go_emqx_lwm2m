// Model for lwm2m-emqx create and write messages when writing multiple values at once

package base_models

import "encoding/json"

func UnmarshalBatchWriteCreateReq(data []byte) (BatchWriteCreateReq, error) {
	var r BatchWriteCreateReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BatchWriteCreateReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BatchWriteCreateReq struct {
	ReqID   int64                `json:"reqID"`
	MsgType string               `json:"msgType"`
	Data    BatchWriteCreateData `json:"data"`
}

type BatchWriteCreateData struct {
	BasePath string                    `json:"basePath"`
	Content  []BatchWriteCreateContent `json:"content"`
}

type BatchWriteCreateContent struct {
	Path  string      `json:"path"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
