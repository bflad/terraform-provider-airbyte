// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Guru string

const (
	GuruGuru Guru = "guru"
)

func (e Guru) ToPointer() *Guru {
	return &e
}
func (e *Guru) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "guru":
		*e = Guru(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Guru: %v", v)
	}
}

type SourceGuru struct {
	Password *string `json:"password,omitempty"`
	// Query for searching cards
	SearchCardsQuery *string   `json:"search_cards_query,omitempty"`
	sourceType       Guru      `const:"guru" json:"sourceType"`
	StartDate        time.Time `json:"start_date"`
	// Team ID received through response of /teams streams, make sure about access to the team
	TeamID   *string `json:"team_id,omitempty"`
	Username string  `json:"username"`
}

func (s SourceGuru) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGuru) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGuru) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceGuru) GetSearchCardsQuery() *string {
	if o == nil {
		return nil
	}
	return o.SearchCardsQuery
}

func (o *SourceGuru) GetSourceType() Guru {
	return GuruGuru
}

func (o *SourceGuru) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceGuru) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *SourceGuru) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
