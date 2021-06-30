package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalUpdateResp(data []byte) (UpdateResp, error) {
	var r UpdateResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type UpdateResp base_models.RegisterUpdateResp
