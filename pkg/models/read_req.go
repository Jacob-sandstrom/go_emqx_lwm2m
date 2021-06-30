package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalReadReq(data []byte) (ReadReq, error) {
	var r ReadReq
	err := json.Unmarshal(data, &r)
	return r, err
}

type ReadReq base_models.ReadDiscoverDeleteReq
