// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Front string

const (
	FrontFront Front = "front"
)

func (e Front) ToPointer() *Front {
	return &e
}
func (e *Front) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "front":
		*e = Front(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Front: %v", v)
	}
}

type SourceFront struct {
	APIKey string `json:"api_key"`
	// Page limit for the responses
	PageLimit  *string   `default:"50" json:"page_limit"`
	sourceType Front     `const:"front" json:"sourceType"`
	StartDate  time.Time `json:"start_date"`
}

func (s SourceFront) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFront) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFront) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceFront) GetPageLimit() *string {
	if o == nil {
		return nil
	}
	return o.PageLimit
}

func (o *SourceFront) GetSourceType() Front {
	return FrontFront
}

func (o *SourceFront) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
