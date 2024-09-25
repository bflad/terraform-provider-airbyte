// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Instagram string

const (
	InstagramInstagram Instagram = "instagram"
)

func (e Instagram) ToPointer() *Instagram {
	return &e
}
func (e *Instagram) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "instagram":
		*e = Instagram(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Instagram: %v", v)
	}
}

type SourceInstagram struct {
	// The value of the access token generated with <b>instagram_basic, instagram_manage_insights, pages_show_list, pages_read_engagement, Instagram Public Content Access</b> permissions. See the <a href="https://docs.airbyte.com/integrations/sources/instagram/#step-1-set-up-instagram">docs</a> for more information
	AccessToken string    `json:"access_token"`
	sourceType  Instagram `const:"instagram" json:"sourceType"`
	// The date from which you'd like to replicate data for User Insights, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated. If left blank, the start date will be set to 2 years before the present date.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceInstagram) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceInstagram) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceInstagram) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceInstagram) GetSourceType() Instagram {
	return InstagramInstagram
}

func (o *SourceInstagram) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
