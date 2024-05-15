// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Aha string

const (
	AhaAha Aha = "aha"
)

func (e Aha) ToPointer() *Aha {
	return &e
}
func (e *Aha) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aha":
		*e = Aha(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Aha: %v", v)
	}
}

type SourceAha struct {
	// API Key
	APIKey     string `json:"api_key"`
	sourceType Aha    `const:"aha" json:"sourceType"`
	// URL
	URL string `json:"url"`
}

func (s SourceAha) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAha) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAha) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceAha) GetSourceType() Aha {
	return AhaAha
}

func (o *SourceAha) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
