// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Buzzsprout string

const (
	BuzzsproutBuzzsprout Buzzsprout = "buzzsprout"
)

func (e Buzzsprout) ToPointer() *Buzzsprout {
	return &e
}
func (e *Buzzsprout) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "buzzsprout":
		*e = Buzzsprout(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Buzzsprout: %v", v)
	}
}

type SourceBuzzsprout struct {
	APIKey string `json:"api_key"`
	// Podcast ID found in `https://www.buzzsprout.com/my/profile/api`
	PodcastID  string     `json:"podcast_id"`
	sourceType Buzzsprout `const:"buzzsprout" json:"sourceType"`
	StartDate  time.Time  `json:"start_date"`
}

func (s SourceBuzzsprout) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBuzzsprout) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBuzzsprout) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceBuzzsprout) GetPodcastID() string {
	if o == nil {
		return ""
	}
	return o.PodcastID
}

func (o *SourceBuzzsprout) GetSourceType() Buzzsprout {
	return BuzzsproutBuzzsprout
}

func (o *SourceBuzzsprout) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
