// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle string

const (
	SourceMixpanelAuthenticationWildcardProjectSecretOptionTitleProjectSecret SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle = "Project Secret"
)

func (e SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle) ToPointer() *SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle {
	return &e
}

func (e *SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Project Secret":
		*e = SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle: %v", v)
	}
}

// SourceMixpanelAuthenticationWildcardProjectSecret - Choose how to authenticate to Mixpanel
type SourceMixpanelAuthenticationWildcardProjectSecret struct {
	// Mixpanel project secret. See the <a href="https://developer.mixpanel.com/reference/project-secret#managing-a-projects-secret">docs</a> for more information on how to obtain this.
	APISecret   string                                                        `json:"api_secret"`
	OptionTitle *SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle `json:"option_title,omitempty"`
}

type SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle string

const (
	SourceMixpanelAuthenticationWildcardServiceAccountOptionTitleServiceAccount SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle = "Service Account"
)

func (e SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle) ToPointer() *SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle {
	return &e
}

func (e *SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service Account":
		*e = SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle: %v", v)
	}
}

// SourceMixpanelAuthenticationWildcardServiceAccount - Choose how to authenticate to Mixpanel
type SourceMixpanelAuthenticationWildcardServiceAccount struct {
	OptionTitle *SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle `json:"option_title,omitempty"`
	// Mixpanel Service Account Secret. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.
	Secret string `json:"secret"`
	// Mixpanel Service Account Username. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.
	Username string `json:"username"`
}

type SourceMixpanelAuthenticationWildcardType string

const (
	SourceMixpanelAuthenticationWildcardTypeSourceMixpanelAuthenticationWildcardServiceAccount SourceMixpanelAuthenticationWildcardType = "source-mixpanel_Authentication *_Service Account"
	SourceMixpanelAuthenticationWildcardTypeSourceMixpanelAuthenticationWildcardProjectSecret  SourceMixpanelAuthenticationWildcardType = "source-mixpanel_Authentication *_Project Secret"
)

type SourceMixpanelAuthenticationWildcard struct {
	SourceMixpanelAuthenticationWildcardServiceAccount *SourceMixpanelAuthenticationWildcardServiceAccount
	SourceMixpanelAuthenticationWildcardProjectSecret  *SourceMixpanelAuthenticationWildcardProjectSecret

	Type SourceMixpanelAuthenticationWildcardType
}

func CreateSourceMixpanelAuthenticationWildcardSourceMixpanelAuthenticationWildcardServiceAccount(sourceMixpanelAuthenticationWildcardServiceAccount SourceMixpanelAuthenticationWildcardServiceAccount) SourceMixpanelAuthenticationWildcard {
	typ := SourceMixpanelAuthenticationWildcardTypeSourceMixpanelAuthenticationWildcardServiceAccount

	return SourceMixpanelAuthenticationWildcard{
		SourceMixpanelAuthenticationWildcardServiceAccount: &sourceMixpanelAuthenticationWildcardServiceAccount,
		Type: typ,
	}
}

func CreateSourceMixpanelAuthenticationWildcardSourceMixpanelAuthenticationWildcardProjectSecret(sourceMixpanelAuthenticationWildcardProjectSecret SourceMixpanelAuthenticationWildcardProjectSecret) SourceMixpanelAuthenticationWildcard {
	typ := SourceMixpanelAuthenticationWildcardTypeSourceMixpanelAuthenticationWildcardProjectSecret

	return SourceMixpanelAuthenticationWildcard{
		SourceMixpanelAuthenticationWildcardProjectSecret: &sourceMixpanelAuthenticationWildcardProjectSecret,
		Type: typ,
	}
}

func (u *SourceMixpanelAuthenticationWildcard) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMixpanelAuthenticationWildcardServiceAccount := new(SourceMixpanelAuthenticationWildcardServiceAccount)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMixpanelAuthenticationWildcardServiceAccount); err == nil {
		u.SourceMixpanelAuthenticationWildcardServiceAccount = sourceMixpanelAuthenticationWildcardServiceAccount
		u.Type = SourceMixpanelAuthenticationWildcardTypeSourceMixpanelAuthenticationWildcardServiceAccount
		return nil
	}

	sourceMixpanelAuthenticationWildcardProjectSecret := new(SourceMixpanelAuthenticationWildcardProjectSecret)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMixpanelAuthenticationWildcardProjectSecret); err == nil {
		u.SourceMixpanelAuthenticationWildcardProjectSecret = sourceMixpanelAuthenticationWildcardProjectSecret
		u.Type = SourceMixpanelAuthenticationWildcardTypeSourceMixpanelAuthenticationWildcardProjectSecret
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMixpanelAuthenticationWildcard) MarshalJSON() ([]byte, error) {
	if u.SourceMixpanelAuthenticationWildcardServiceAccount != nil {
		return json.Marshal(u.SourceMixpanelAuthenticationWildcardServiceAccount)
	}

	if u.SourceMixpanelAuthenticationWildcardProjectSecret != nil {
		return json.Marshal(u.SourceMixpanelAuthenticationWildcardProjectSecret)
	}

	return nil, nil
}

// SourceMixpanelRegion - The region of mixpanel domain instance either US or EU.
type SourceMixpanelRegion string

const (
	SourceMixpanelRegionUs SourceMixpanelRegion = "US"
	SourceMixpanelRegionEu SourceMixpanelRegion = "EU"
)

func (e SourceMixpanelRegion) ToPointer() *SourceMixpanelRegion {
	return &e
}

func (e *SourceMixpanelRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = SourceMixpanelRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelRegion: %v", v)
	}
}

type SourceMixpanelMixpanel string

const (
	SourceMixpanelMixpanelMixpanel SourceMixpanelMixpanel = "mixpanel"
)

func (e SourceMixpanelMixpanel) ToPointer() *SourceMixpanelMixpanel {
	return &e
}

func (e *SourceMixpanelMixpanel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mixpanel":
		*e = SourceMixpanelMixpanel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelMixpanel: %v", v)
	}
}

type SourceMixpanel struct {
	//  A period of time for attributing results to ads and the lookback period after those actions occur during which ad results are counted. Default attribution window is 5 days.
	AttributionWindow *int64 `json:"attribution_window,omitempty"`
	// Choose how to authenticate to Mixpanel
	Credentials *SourceMixpanelAuthenticationWildcard `json:"credentials,omitempty"`
	// Defines window size in days, that used to slice through data. You can reduce it, if amount of data in each window is too big for your environment.
	DateWindowSize *int64 `json:"date_window_size,omitempty"`
	// The date in the format YYYY-MM-DD. Any data after this date will not be replicated. Left empty to always sync to most recent date
	EndDate *types.Date `json:"end_date,omitempty"`
	// Your project ID number. See the <a href="https://help.mixpanel.com/hc/en-us/articles/115004490503-Project-Settings#project-id">docs</a> for more information on how to obtain this.
	ProjectID *int64 `json:"project_id,omitempty"`
	// Time zone in which integer date times are stored. The project timezone may be found in the project settings in the <a href="https://help.mixpanel.com/hc/en-us/articles/115004547203-Manage-Timezones-for-Projects-in-Mixpanel">Mixpanel console</a>.
	ProjectTimezone *string `json:"project_timezone,omitempty"`
	// The region of mixpanel domain instance either US or EU.
	Region *SourceMixpanelRegion `json:"region,omitempty"`
	// Setting this config parameter to TRUE ensures that new properties on events and engage records are captured. Otherwise new properties will be ignored.
	SelectPropertiesByDefault *bool                   `json:"select_properties_by_default,omitempty"`
	SourceType                *SourceMixpanelMixpanel `json:"sourceType,omitempty"`
	// The date in the format YYYY-MM-DD. Any data before this date will not be replicated. If this option is not set, the connector will replicate data from up to one year ago by default.
	StartDate *types.Date `json:"start_date,omitempty"`
}
