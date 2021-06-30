package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalRegisterResp(data []byte) (RegisterResp, error) {
	var r RegisterResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RegisterResp base_models.RegisterUpdateResp
