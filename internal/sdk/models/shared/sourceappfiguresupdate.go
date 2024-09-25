// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

// GroupBy - Category term for grouping the search results
type GroupBy string

const (
	GroupByNetwork GroupBy = "network"
	GroupByProduct GroupBy = "product"
	GroupByCountry GroupBy = "country"
	GroupByDate    GroupBy = "date"
)

func (e GroupBy) ToPointer() *GroupBy {
	return &e
}
func (e *GroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "network":
		fallthrough
	case "product":
		fallthrough
	case "country":
		fallthrough
	case "date":
		*e = GroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupBy: %v", v)
	}
}

type SourceAppfiguresUpdate struct {
	APIKey string `json:"api_key"`
	// Category term for grouping the search results
	GroupBy *GroupBy `default:"product" json:"group_by"`
	// The store which needs to be searched in streams
	SearchStore *string   `default:"apple" json:"search_store"`
	StartDate   time.Time `json:"start_date"`
}

func (s SourceAppfiguresUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAppfiguresUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAppfiguresUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceAppfiguresUpdate) GetGroupBy() *GroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *SourceAppfiguresUpdate) GetSearchStore() *string {
	if o == nil {
		return nil
	}
	return o.SearchStore
}

func (o *SourceAppfiguresUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
