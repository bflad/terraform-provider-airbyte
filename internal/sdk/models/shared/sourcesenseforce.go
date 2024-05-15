// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
)

type Senseforce string

const (
	SenseforceSenseforce Senseforce = "senseforce"
)

func (e Senseforce) ToPointer() *Senseforce {
	return &e
}
func (e *Senseforce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "senseforce":
		*e = Senseforce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Senseforce: %v", v)
	}
}

type SourceSenseforce struct {
	// Your API access token. See <a href="https://manual.senseforce.io/manual/sf-platform/public-api/get-your-access-token/">here</a>. The toke is case sensitive.
	AccessToken string `json:"access_token"`
	// Your Senseforce API backend URL. This is the URL shown during the Login screen. See <a href="https://manual.senseforce.io/manual/sf-platform/public-api/get-your-access-token/">here</a> for more details. (Note: Most Senseforce backend APIs have the term 'galaxy' in their ULR)
	BackendURL string `json:"backend_url"`
	// The ID of the dataset you want to synchronize. The ID can be found in the URL when opening the dataset. See <a href="https://manual.senseforce.io/manual/sf-platform/public-api/get-your-access-token/">here</a> for more details. (Note: As the Senseforce API only allows to synchronize a specific dataset, each dataset you  want to synchronize needs to be implemented as a separate airbyte source).
	DatasetID string `json:"dataset_id"`
	// The time increment used by the connector when requesting data from the Senseforce API. The bigger the value is, the less requests will be made and faster the sync will be. On the other hand, the more seldom the state is persisted and the more likely one could run into rate limites.  Furthermore, consider that large chunks of time might take a long time for the Senseforce query to return data - meaning it could take in effect longer than with more smaller time slices. If there are a lot of data per day, set this setting to 1. If there is only very little data per day, you might change the setting to 10 or more.
	SliceRange *int64     `default:"10" json:"slice_range"`
	sourceType Senseforce `const:"senseforce" json:"sourceType"`
	// UTC date and time in the format 2017-01-25. Only data with "Timestamp" after this date will be replicated. Important note: This start date must be set to the first day of where your dataset provides data.  If your dataset has data from 2020-10-10 10:21:10, set the start_date to 2020-10-10 or later
	StartDate types.Date `json:"start_date"`
}

func (s SourceSenseforce) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSenseforce) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSenseforce) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceSenseforce) GetBackendURL() string {
	if o == nil {
		return ""
	}
	return o.BackendURL
}

func (o *SourceSenseforce) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

func (o *SourceSenseforce) GetSliceRange() *int64 {
	if o == nil {
		return nil
	}
	return o.SliceRange
}

func (o *SourceSenseforce) GetSourceType() Senseforce {
	return SenseforceSenseforce
}

func (o *SourceSenseforce) GetStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.StartDate
}
