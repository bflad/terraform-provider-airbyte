// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Bitly string

const (
	BitlyBitly Bitly = "bitly"
)

func (e Bitly) ToPointer() *Bitly {
	return &e
}
func (e *Bitly) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bitly":
		*e = Bitly(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Bitly: %v", v)
	}
}

type SourceBitly struct {
	APIKey     string    `json:"api_key"`
	EndDate    time.Time `json:"end_date"`
	sourceType Bitly     `const:"bitly" json:"sourceType"`
	StartDate  time.Time `json:"start_date"`
}

func (s SourceBitly) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBitly) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBitly) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceBitly) GetEndDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndDate
}

func (o *SourceBitly) GetSourceType() Bitly {
	return BitlyBitly
}

func (o *SourceBitly) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
