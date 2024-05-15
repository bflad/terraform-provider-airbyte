// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SapFieldglass string

const (
	SapFieldglassSapFieldglass SapFieldglass = "sap-fieldglass"
)

func (e SapFieldglass) ToPointer() *SapFieldglass {
	return &e
}
func (e *SapFieldglass) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sap-fieldglass":
		*e = SapFieldglass(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SapFieldglass: %v", v)
	}
}

type SourceSapFieldglass struct {
	// API Key
	APIKey     string        `json:"api_key"`
	sourceType SapFieldglass `const:"sap-fieldglass" json:"sourceType"`
}

func (s SourceSapFieldglass) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSapFieldglass) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSapFieldglass) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceSapFieldglass) GetSourceType() SapFieldglass {
	return SapFieldglassSapFieldglass
}
