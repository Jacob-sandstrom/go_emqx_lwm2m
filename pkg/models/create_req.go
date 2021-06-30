package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalCreateReq(data []byte) (CreateReq, error) {
	var r CreateReq
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateReq base_models.BatchWriteCreateReq
