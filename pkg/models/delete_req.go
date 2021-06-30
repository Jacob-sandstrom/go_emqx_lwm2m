package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalDeleteReq(data []byte) (DeleteReq, error) {
	var r DeleteReq
	err := json.Unmarshal(data, &r)
	return r, err
}

type DeleteReq base_models.ReadDiscoverDeleteReq
