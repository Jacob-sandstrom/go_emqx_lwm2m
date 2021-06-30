package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalExecuteResp(data []byte) (ExecuteResp, error) {
	var r ExecuteResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ExecuteResp base_models.StandardResp
