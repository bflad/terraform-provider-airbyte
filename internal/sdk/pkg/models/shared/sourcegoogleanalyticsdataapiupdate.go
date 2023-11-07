// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType string

const (
	SourceGoogleAnalyticsDataAPIUpdateSchemasAuthTypeService SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType = "Service"
)

func (e SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType) ToPointer() *SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType {
	return &e
}

func (e *SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service":
		*e = SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType: %v", v)
	}
}

// ServiceAccountKeyAuthentication - Credentials for the service
type ServiceAccountKeyAuthentication struct {
	authType *SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType `const:"Service" json:"auth_type,omitempty"`
	// The JSON key linked to the service account used for authorization. For steps on obtaining this key, refer to <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#setup-guide">the setup guide</a>.
	CredentialsJSON string `json:"credentials_json"`
}

func (s ServiceAccountKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceAccountKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ServiceAccountKeyAuthentication) GetAuthType() *SourceGoogleAnalyticsDataAPIUpdateSchemasAuthType {
	return SourceGoogleAnalyticsDataAPIUpdateSchemasAuthTypeService.ToPointer()
}

func (o *ServiceAccountKeyAuthentication) GetCredentialsJSON() string {
	if o == nil {
		return ""
	}
	return o.CredentialsJSON
}

type SourceGoogleAnalyticsDataAPIUpdateAuthType string

const (
	SourceGoogleAnalyticsDataAPIUpdateAuthTypeClient SourceGoogleAnalyticsDataAPIUpdateAuthType = "Client"
)

func (e SourceGoogleAnalyticsDataAPIUpdateAuthType) ToPointer() *SourceGoogleAnalyticsDataAPIUpdateAuthType {
	return &e
}

func (e *SourceGoogleAnalyticsDataAPIUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceGoogleAnalyticsDataAPIUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleAnalyticsDataAPIUpdateAuthType: %v", v)
	}
}

// AuthenticateViaGoogleOauth - Credentials for the service
type AuthenticateViaGoogleOauth struct {
	// Access Token for making authenticated requests.
	AccessToken *string                                     `json:"access_token,omitempty"`
	authType    *SourceGoogleAnalyticsDataAPIUpdateAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Google Analytics developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Google Analytics developer application.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining a new access token.
	RefreshToken string `json:"refresh_token"`
}

func (a AuthenticateViaGoogleOauth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaGoogleOauth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaGoogleOauth) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *AuthenticateViaGoogleOauth) GetAuthType() *SourceGoogleAnalyticsDataAPIUpdateAuthType {
	return SourceGoogleAnalyticsDataAPIUpdateAuthTypeClient.ToPointer()
}

func (o *AuthenticateViaGoogleOauth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *AuthenticateViaGoogleOauth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *AuthenticateViaGoogleOauth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type CredentialsType string

const (
	CredentialsTypeAuthenticateViaGoogleOauth      CredentialsType = "AuthenticateViaGoogleOauth"
	CredentialsTypeServiceAccountKeyAuthentication CredentialsType = "ServiceAccountKeyAuthentication"
)

type Credentials struct {
	AuthenticateViaGoogleOauth      *AuthenticateViaGoogleOauth
	ServiceAccountKeyAuthentication *ServiceAccountKeyAuthentication

	Type CredentialsType
}

func CreateCredentialsAuthenticateViaGoogleOauth(authenticateViaGoogleOauth AuthenticateViaGoogleOauth) Credentials {
	typ := CredentialsTypeAuthenticateViaGoogleOauth

	return Credentials{
		AuthenticateViaGoogleOauth: &authenticateViaGoogleOauth,
		Type:                       typ,
	}
}

func CreateCredentialsServiceAccountKeyAuthentication(serviceAccountKeyAuthentication ServiceAccountKeyAuthentication) Credentials {
	typ := CredentialsTypeServiceAccountKeyAuthentication

	return Credentials{
		ServiceAccountKeyAuthentication: &serviceAccountKeyAuthentication,
		Type:                            typ,
	}
}

func (u *Credentials) UnmarshalJSON(data []byte) error {

	serviceAccountKeyAuthentication := new(ServiceAccountKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &serviceAccountKeyAuthentication, "", true, true); err == nil {
		u.ServiceAccountKeyAuthentication = serviceAccountKeyAuthentication
		u.Type = CredentialsTypeServiceAccountKeyAuthentication
		return nil
	}

	authenticateViaGoogleOauth := new(AuthenticateViaGoogleOauth)
	if err := utils.UnmarshalJSON(data, &authenticateViaGoogleOauth, "", true, true); err == nil {
		u.AuthenticateViaGoogleOauth = authenticateViaGoogleOauth
		u.Type = CredentialsTypeAuthenticateViaGoogleOauth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Credentials) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaGoogleOauth != nil {
		return utils.MarshalJSON(u.AuthenticateViaGoogleOauth, "", true)
	}

	if u.ServiceAccountKeyAuthentication != nil {
		return utils.MarshalJSON(u.ServiceAccountKeyAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceGoogleAnalyticsDataAPIUpdate struct {
	// Credentials for the service
	Credentials *Credentials `json:"credentials,omitempty"`
	// A JSON array describing the custom reports you want to sync from Google Analytics. See <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#custom-reports">the documentation</a> for more information about the exact format you can use to fill out this field.
	CustomReports *string `json:"custom_reports,omitempty"`
	// The start date from which to replicate report data in the format YYYY-MM-DD. Data generated before this date will not be included in the report. Not applied to custom Cohort reports.
	DateRangesStartDate types.Date `json:"date_ranges_start_date"`
	// The Property ID is a unique number assigned to each property in Google Analytics, found in your GA4 property URL. This ID allows the connector to track the specific events associated with your property. Refer to the <a href='https://developers.google.com/analytics/devguides/reporting/data/v1/property-id#what_is_my_property_id'>Google Analytics documentation</a> to locate your property ID.
	PropertyID string `json:"property_id"`
	// The interval in days for each data request made to the Google Analytics API. A larger value speeds up data sync, but increases the chance of data sampling, which may result in inaccuracies. We recommend a value of 1 to minimize sampling, unless speed is an absolute priority over accuracy. Acceptable values range from 1 to 364. Does not apply to custom Cohort reports. More information is available in <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api">the documentation</a>.
	WindowInDays *int64 `default:"1" json:"window_in_days"`
}

func (s SourceGoogleAnalyticsDataAPIUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleAnalyticsDataAPIUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleAnalyticsDataAPIUpdate) GetCredentials() *Credentials {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceGoogleAnalyticsDataAPIUpdate) GetCustomReports() *string {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourceGoogleAnalyticsDataAPIUpdate) GetDateRangesStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.DateRangesStartDate
}

func (o *SourceGoogleAnalyticsDataAPIUpdate) GetPropertyID() string {
	if o == nil {
		return ""
	}
	return o.PropertyID
}

func (o *SourceGoogleAnalyticsDataAPIUpdate) GetWindowInDays() *int64 {
	if o == nil {
		return nil
	}
	return o.WindowInDays
}
