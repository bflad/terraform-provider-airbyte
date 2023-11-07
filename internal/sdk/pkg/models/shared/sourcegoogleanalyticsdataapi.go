// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceGoogleAnalyticsDataAPISchemasAuthType string

const (
	SourceGoogleAnalyticsDataAPISchemasAuthTypeService SourceGoogleAnalyticsDataAPISchemasAuthType = "Service"
)

func (e SourceGoogleAnalyticsDataAPISchemasAuthType) ToPointer() *SourceGoogleAnalyticsDataAPISchemasAuthType {
	return &e
}

func (e *SourceGoogleAnalyticsDataAPISchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service":
		*e = SourceGoogleAnalyticsDataAPISchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleAnalyticsDataAPISchemasAuthType: %v", v)
	}
}

// SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication - Credentials for the service
type SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication struct {
	authType *SourceGoogleAnalyticsDataAPISchemasAuthType `const:"Service" json:"auth_type,omitempty"`
	// The JSON key linked to the service account used for authorization. For steps on obtaining this key, refer to <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#setup-guide">the setup guide</a>.
	CredentialsJSON string `json:"credentials_json"`
}

func (s SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication) GetAuthType() *SourceGoogleAnalyticsDataAPISchemasAuthType {
	return SourceGoogleAnalyticsDataAPISchemasAuthTypeService.ToPointer()
}

func (o *SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication) GetCredentialsJSON() string {
	if o == nil {
		return ""
	}
	return o.CredentialsJSON
}

type SourceGoogleAnalyticsDataAPIAuthType string

const (
	SourceGoogleAnalyticsDataAPIAuthTypeClient SourceGoogleAnalyticsDataAPIAuthType = "Client"
)

func (e SourceGoogleAnalyticsDataAPIAuthType) ToPointer() *SourceGoogleAnalyticsDataAPIAuthType {
	return &e
}

func (e *SourceGoogleAnalyticsDataAPIAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceGoogleAnalyticsDataAPIAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleAnalyticsDataAPIAuthType: %v", v)
	}
}

// SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth - Credentials for the service
type SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth struct {
	// Access Token for making authenticated requests.
	AccessToken *string                               `json:"access_token,omitempty"`
	authType    *SourceGoogleAnalyticsDataAPIAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Google Analytics developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Google Analytics developer application.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining a new access token.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) GetAuthType() *SourceGoogleAnalyticsDataAPIAuthType {
	return SourceGoogleAnalyticsDataAPIAuthTypeClient.ToPointer()
}

func (o *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceGoogleAnalyticsDataAPICredentialsType string

const (
	SourceGoogleAnalyticsDataAPICredentialsTypeAuthenticateViaGoogleOauth      SourceGoogleAnalyticsDataAPICredentialsType = "AuthenticateViaGoogleOauth"
	SourceGoogleAnalyticsDataAPICredentialsTypeServiceAccountKeyAuthentication SourceGoogleAnalyticsDataAPICredentialsType = "ServiceAccountKeyAuthentication"
)

type SourceGoogleAnalyticsDataAPICredentials struct {
	AuthenticateViaGoogleOauth      *SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth
	ServiceAccountKeyAuthentication *SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication

	Type SourceGoogleAnalyticsDataAPICredentialsType
}

func CreateSourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth(authenticateViaGoogleOauth SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth) SourceGoogleAnalyticsDataAPICredentials {
	typ := SourceGoogleAnalyticsDataAPICredentialsTypeAuthenticateViaGoogleOauth

	return SourceGoogleAnalyticsDataAPICredentials{
		AuthenticateViaGoogleOauth: &authenticateViaGoogleOauth,
		Type:                       typ,
	}
}

func CreateSourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication(serviceAccountKeyAuthentication SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication) SourceGoogleAnalyticsDataAPICredentials {
	typ := SourceGoogleAnalyticsDataAPICredentialsTypeServiceAccountKeyAuthentication

	return SourceGoogleAnalyticsDataAPICredentials{
		ServiceAccountKeyAuthentication: &serviceAccountKeyAuthentication,
		Type:                            typ,
	}
}

func (u *SourceGoogleAnalyticsDataAPICredentials) UnmarshalJSON(data []byte) error {

	serviceAccountKeyAuthentication := new(SourceGoogleAnalyticsDataAPIServiceAccountKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &serviceAccountKeyAuthentication, "", true, true); err == nil {
		u.ServiceAccountKeyAuthentication = serviceAccountKeyAuthentication
		u.Type = SourceGoogleAnalyticsDataAPICredentialsTypeServiceAccountKeyAuthentication
		return nil
	}

	authenticateViaGoogleOauth := new(SourceGoogleAnalyticsDataAPIAuthenticateViaGoogleOauth)
	if err := utils.UnmarshalJSON(data, &authenticateViaGoogleOauth, "", true, true); err == nil {
		u.AuthenticateViaGoogleOauth = authenticateViaGoogleOauth
		u.Type = SourceGoogleAnalyticsDataAPICredentialsTypeAuthenticateViaGoogleOauth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGoogleAnalyticsDataAPICredentials) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaGoogleOauth != nil {
		return utils.MarshalJSON(u.AuthenticateViaGoogleOauth, "", true)
	}

	if u.ServiceAccountKeyAuthentication != nil {
		return utils.MarshalJSON(u.ServiceAccountKeyAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GoogleAnalyticsDataAPI string

const (
	GoogleAnalyticsDataAPIGoogleAnalyticsDataAPI GoogleAnalyticsDataAPI = "google-analytics-data-api"
)

func (e GoogleAnalyticsDataAPI) ToPointer() *GoogleAnalyticsDataAPI {
	return &e
}

func (e *GoogleAnalyticsDataAPI) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-analytics-data-api":
		*e = GoogleAnalyticsDataAPI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GoogleAnalyticsDataAPI: %v", v)
	}
}

type SourceGoogleAnalyticsDataAPI struct {
	// Credentials for the service
	Credentials *SourceGoogleAnalyticsDataAPICredentials `json:"credentials,omitempty"`
	// A JSON array describing the custom reports you want to sync from Google Analytics. See <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#custom-reports">the documentation</a> for more information about the exact format you can use to fill out this field.
	CustomReports *string `json:"custom_reports,omitempty"`
	// The start date from which to replicate report data in the format YYYY-MM-DD. Data generated before this date will not be included in the report. Not applied to custom Cohort reports.
	DateRangesStartDate types.Date `json:"date_ranges_start_date"`
	// The Property ID is a unique number assigned to each property in Google Analytics, found in your GA4 property URL. This ID allows the connector to track the specific events associated with your property. Refer to the <a href='https://developers.google.com/analytics/devguides/reporting/data/v1/property-id#what_is_my_property_id'>Google Analytics documentation</a> to locate your property ID.
	PropertyID string                 `json:"property_id"`
	sourceType GoogleAnalyticsDataAPI `const:"google-analytics-data-api" json:"sourceType"`
	// The interval in days for each data request made to the Google Analytics API. A larger value speeds up data sync, but increases the chance of data sampling, which may result in inaccuracies. We recommend a value of 1 to minimize sampling, unless speed is an absolute priority over accuracy. Acceptable values range from 1 to 364. Does not apply to custom Cohort reports. More information is available in <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api">the documentation</a>.
	WindowInDays *int64 `default:"1" json:"window_in_days"`
}

func (s SourceGoogleAnalyticsDataAPI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleAnalyticsDataAPI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleAnalyticsDataAPI) GetCredentials() *SourceGoogleAnalyticsDataAPICredentials {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceGoogleAnalyticsDataAPI) GetCustomReports() *string {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourceGoogleAnalyticsDataAPI) GetDateRangesStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.DateRangesStartDate
}

func (o *SourceGoogleAnalyticsDataAPI) GetPropertyID() string {
	if o == nil {
		return ""
	}
	return o.PropertyID
}

func (o *SourceGoogleAnalyticsDataAPI) GetSourceType() GoogleAnalyticsDataAPI {
	return GoogleAnalyticsDataAPIGoogleAnalyticsDataAPI
}

func (o *SourceGoogleAnalyticsDataAPI) GetWindowInDays() *int64 {
	if o == nil {
		return nil
	}
	return o.WindowInDays
}
