package models

import (
	"encoding/json"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
)

func UnmarshalWriteAttrResp(data []byte) (WriteAttrResp, error) {
	var r WriteAttrResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type WriteAttrResp base_models.StandardResp
