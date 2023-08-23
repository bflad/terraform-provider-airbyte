// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceKyveKyve string

const (
	SourceKyveKyveKyve SourceKyveKyve = "kyve"
)

func (e SourceKyveKyve) ToPointer() *SourceKyveKyve {
	return &e
}

func (e *SourceKyveKyve) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kyve":
		*e = SourceKyveKyve(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceKyveKyve: %v", v)
	}
}

type SourceKyve struct {
	// The maximum amount of pages to go trough. Set to 'null' for all pages.
	MaxPages *int64 `json:"max_pages,omitempty"`
	// The pagesize for pagination, smaller numbers are used in integration tests.
	PageSize *int64 `json:"page_size,omitempty"`
	// The IDs of the KYVE storage pool you want to archive. (Comma separated)
	PoolIds    string         `json:"pool_ids"`
	SourceType SourceKyveKyve `json:"sourceType"`
	// The start-id defines, from which bundle id the pipeline should start to extract the data (Comma separated)
	StartIds string `json:"start_ids"`
	// URL to the KYVE Chain API.
	URLBase *string `json:"url_base,omitempty"`
}
