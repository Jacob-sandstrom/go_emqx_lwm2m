package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalBatchWriteReq(data []byte) (BatchWriteReq, error) {
	var r BatchWriteReq
	err := json.Unmarshal(data, &r)
	return r, err
}

type BatchWriteReq base_models.BatchWriteCreateReq
