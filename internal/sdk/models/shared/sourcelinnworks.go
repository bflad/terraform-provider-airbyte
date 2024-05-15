// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Linnworks string

const (
	LinnworksLinnworks Linnworks = "linnworks"
)

func (e Linnworks) ToPointer() *Linnworks {
	return &e
}
func (e *Linnworks) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linnworks":
		*e = Linnworks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Linnworks: %v", v)
	}
}

type SourceLinnworks struct {
	// Linnworks Application ID
	ApplicationID string `json:"application_id"`
	// Linnworks Application Secret
	ApplicationSecret string    `json:"application_secret"`
	sourceType        Linnworks `const:"linnworks" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
	Token     string    `json:"token"`
}

func (s SourceLinnworks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinnworks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinnworks) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

func (o *SourceLinnworks) GetApplicationSecret() string {
	if o == nil {
		return ""
	}
	return o.ApplicationSecret
}

func (o *SourceLinnworks) GetSourceType() Linnworks {
	return LinnworksLinnworks
}

func (o *SourceLinnworks) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceLinnworks) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
