// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceGoogleSearchConsoleUpdateSchemasAuthType string

const (
	SourceGoogleSearchConsoleUpdateSchemasAuthTypeService SourceGoogleSearchConsoleUpdateSchemasAuthType = "Service"
)

func (e SourceGoogleSearchConsoleUpdateSchemasAuthType) ToPointer() *SourceGoogleSearchConsoleUpdateSchemasAuthType {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service":
		*e = SourceGoogleSearchConsoleUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateSchemasAuthType: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication struct {
	authType SourceGoogleSearchConsoleUpdateSchemasAuthType `const:"Service" json:"auth_type"`
	// The email of the user which has permissions to access the Google Workspace Admin APIs.
	Email string `json:"email"`
	// The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.
	ServiceAccountInfo string `json:"service_account_info"`
}

func (s SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication) GetAuthType() SourceGoogleSearchConsoleUpdateSchemasAuthType {
	return SourceGoogleSearchConsoleUpdateSchemasAuthTypeService
}

func (o *SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication) GetServiceAccountInfo() string {
	if o == nil {
		return ""
	}
	return o.ServiceAccountInfo
}

type SourceGoogleSearchConsoleUpdateAuthType string

const (
	SourceGoogleSearchConsoleUpdateAuthTypeClient SourceGoogleSearchConsoleUpdateAuthType = "Client"
)

func (e SourceGoogleSearchConsoleUpdateAuthType) ToPointer() *SourceGoogleSearchConsoleUpdateAuthType {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceGoogleSearchConsoleUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateAuthType: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdateOAuth struct {
	// Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	AccessToken *string                                 `json:"access_token,omitempty"`
	authType    SourceGoogleSearchConsoleUpdateAuthType `const:"Client" json:"auth_type"`
	// The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	ClientID string `json:"client_id"`
	// The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceGoogleSearchConsoleUpdateOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleUpdateOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleUpdateOAuth) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceGoogleSearchConsoleUpdateOAuth) GetAuthType() SourceGoogleSearchConsoleUpdateAuthType {
	return SourceGoogleSearchConsoleUpdateAuthTypeClient
}

func (o *SourceGoogleSearchConsoleUpdateOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceGoogleSearchConsoleUpdateOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceGoogleSearchConsoleUpdateOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type AuthenticationTypeType string

const (
	AuthenticationTypeTypeOAuth                           AuthenticationTypeType = "OAuth"
	AuthenticationTypeTypeServiceAccountKeyAuthentication AuthenticationTypeType = "ServiceAccountKeyAuthentication"
)

type AuthenticationType struct {
	OAuth                           *SourceGoogleSearchConsoleUpdateOAuth
	ServiceAccountKeyAuthentication *SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication

	Type AuthenticationTypeType
}

func CreateAuthenticationTypeOAuth(oAuth SourceGoogleSearchConsoleUpdateOAuth) AuthenticationType {
	typ := AuthenticationTypeTypeOAuth

	return AuthenticationType{
		OAuth: &oAuth,
		Type:  typ,
	}
}

func CreateAuthenticationTypeServiceAccountKeyAuthentication(serviceAccountKeyAuthentication SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication) AuthenticationType {
	typ := AuthenticationTypeTypeServiceAccountKeyAuthentication

	return AuthenticationType{
		ServiceAccountKeyAuthentication: &serviceAccountKeyAuthentication,
		Type:                            typ,
	}
}

func (u *AuthenticationType) UnmarshalJSON(data []byte) error {

	serviceAccountKeyAuthentication := new(SourceGoogleSearchConsoleUpdateServiceAccountKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &serviceAccountKeyAuthentication, "", true, true); err == nil {
		u.ServiceAccountKeyAuthentication = serviceAccountKeyAuthentication
		u.Type = AuthenticationTypeTypeServiceAccountKeyAuthentication
		return nil
	}

	oAuth := new(SourceGoogleSearchConsoleUpdateOAuth)
	if err := utils.UnmarshalJSON(data, &oAuth, "", true, true); err == nil {
		u.OAuth = oAuth
		u.Type = AuthenticationTypeTypeOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AuthenticationType) MarshalJSON() ([]byte, error) {
	if u.OAuth != nil {
		return utils.MarshalJSON(u.OAuth, "", true)
	}

	if u.ServiceAccountKeyAuthentication != nil {
		return utils.MarshalJSON(u.ServiceAccountKeyAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceGoogleSearchConsoleUpdateValidEnums - An enumeration of dimensions.
type SourceGoogleSearchConsoleUpdateValidEnums string

const (
	SourceGoogleSearchConsoleUpdateValidEnumsCountry SourceGoogleSearchConsoleUpdateValidEnums = "country"
	SourceGoogleSearchConsoleUpdateValidEnumsDate    SourceGoogleSearchConsoleUpdateValidEnums = "date"
	SourceGoogleSearchConsoleUpdateValidEnumsDevice  SourceGoogleSearchConsoleUpdateValidEnums = "device"
	SourceGoogleSearchConsoleUpdateValidEnumsPage    SourceGoogleSearchConsoleUpdateValidEnums = "page"
	SourceGoogleSearchConsoleUpdateValidEnumsQuery   SourceGoogleSearchConsoleUpdateValidEnums = "query"
)

func (e SourceGoogleSearchConsoleUpdateValidEnums) ToPointer() *SourceGoogleSearchConsoleUpdateValidEnums {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "country":
		fallthrough
	case "date":
		fallthrough
	case "device":
		fallthrough
	case "page":
		fallthrough
	case "query":
		*e = SourceGoogleSearchConsoleUpdateValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateValidEnums: %v", v)
	}
}

type CustomReportConfig struct {
	// A list of dimensions (country, date, device, page, query)
	Dimensions []SourceGoogleSearchConsoleUpdateValidEnums `json:"dimensions"`
	// The name of the custom report, this name would be used as stream name
	Name string `json:"name"`
}

func (o *CustomReportConfig) GetDimensions() []SourceGoogleSearchConsoleUpdateValidEnums {
	if o == nil {
		return []SourceGoogleSearchConsoleUpdateValidEnums{}
	}
	return o.Dimensions
}

func (o *CustomReportConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// DataFreshness - If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
type DataFreshness string

const (
	DataFreshnessFinal DataFreshness = "final"
	DataFreshnessAll   DataFreshness = "all"
)

func (e DataFreshness) ToPointer() *DataFreshness {
	return &e
}

func (e *DataFreshness) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "final":
		fallthrough
	case "all":
		*e = DataFreshness(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DataFreshness: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdate struct {
	Authorization AuthenticationType `json:"authorization"`
	// (DEPRCATED) A JSON array describing the custom reports you want to sync from Google Search Console. See our <a href='https://docs.airbyte.com/integrations/sources/google-search-console'>documentation</a> for more information on formulating custom reports.
	CustomReports *string `json:"custom_reports,omitempty"`
	// You can add your Custom Analytics report by creating one.
	CustomReportsArray []CustomReportConfig `json:"custom_reports_array,omitempty"`
	// If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
	DataState *DataFreshness `default:"final" json:"data_state"`
	// UTC date in the format YYYY-MM-DD. Any data created after this date will not be replicated. Must be greater or equal to the start date field. Leaving this field blank will replicate all data from the start date onward.
	EndDate *types.Date `json:"end_date,omitempty"`
	// The URLs of the website property attached to your GSC account. Learn more about properties <a href="https://support.google.com/webmasters/answer/34592?hl=en">here</a>.
	SiteUrls []string `json:"site_urls"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.
	StartDate *types.Date `default:"2021-01-01" json:"start_date"`
}

func (s SourceGoogleSearchConsoleUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleUpdate) GetAuthorization() AuthenticationType {
	if o == nil {
		return AuthenticationType{}
	}
	return o.Authorization
}

func (o *SourceGoogleSearchConsoleUpdate) GetCustomReports() *string {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourceGoogleSearchConsoleUpdate) GetCustomReportsArray() []CustomReportConfig {
	if o == nil {
		return nil
	}
	return o.CustomReportsArray
}

func (o *SourceGoogleSearchConsoleUpdate) GetDataState() *DataFreshness {
	if o == nil {
		return nil
	}
	return o.DataState
}

func (o *SourceGoogleSearchConsoleUpdate) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceGoogleSearchConsoleUpdate) GetSiteUrls() []string {
	if o == nil {
		return []string{}
	}
	return o.SiteUrls
}

func (o *SourceGoogleSearchConsoleUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}
