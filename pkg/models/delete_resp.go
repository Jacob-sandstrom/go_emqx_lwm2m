package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalDeleteResp(data []byte) (DeleteResp, error) {
	var r DeleteResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DeleteResp base_models.StandardResp
