// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Emailoctopus string

const (
	EmailoctopusEmailoctopus Emailoctopus = "emailoctopus"
)

func (e Emailoctopus) ToPointer() *Emailoctopus {
	return &e
}

func (e *Emailoctopus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "emailoctopus":
		*e = Emailoctopus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Emailoctopus: %v", v)
	}
}

type SourceEmailoctopus struct {
	// EmailOctopus API Key. See the <a href="https://help.emailoctopus.com/article/165-how-to-create-and-delete-api-keys">docs</a> for information on how to generate this key.
	APIKey     string       `json:"api_key"`
	sourceType Emailoctopus `const:"emailoctopus" json:"sourceType"`
}

func (s SourceEmailoctopus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceEmailoctopus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceEmailoctopus) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceEmailoctopus) GetSourceType() Emailoctopus {
	return EmailoctopusEmailoctopus
}
