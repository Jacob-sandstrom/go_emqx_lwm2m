package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalDiscoverReq(data []byte) (DiscoverReq, error) {
	var r DiscoverReq
	err := json.Unmarshal(data, &r)
	return r, err
}

type DiscoverReq struct {
	base_models.ReadDiscoverDeleteReq
}

type DiscoverReqData base_models.ReadDiscoverDeleteData
