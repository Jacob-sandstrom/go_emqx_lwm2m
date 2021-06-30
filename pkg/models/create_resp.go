package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalCreateResp(data []byte) (CreateResp, error) {
	var r CreateResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateResp base_models.StandardResp
