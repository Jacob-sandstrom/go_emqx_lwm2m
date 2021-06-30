package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalWriteResp(data []byte) (WriteResp, error) {
	var r WriteResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type WriteResp base_models.StandardResp
