// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type BothUsernameAndPasswordIsRequiredForAuthenticationRequest string

const (
	BothUsernameAndPasswordIsRequiredForAuthenticationRequestUsernamePassword BothUsernameAndPasswordIsRequiredForAuthenticationRequest = "username_password"
)

func (e BothUsernameAndPasswordIsRequiredForAuthenticationRequest) ToPointer() *BothUsernameAndPasswordIsRequiredForAuthenticationRequest {
	return &e
}

func (e *BothUsernameAndPasswordIsRequiredForAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "username_password":
		*e = BothUsernameAndPasswordIsRequiredForAuthenticationRequest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BothUsernameAndPasswordIsRequiredForAuthenticationRequest: %v", v)
	}
}

// SourceOutbrainAmplifyUpdateUsernamePassword - Credentials for making authenticated requests requires either username/password or access_token.
type SourceOutbrainAmplifyUpdateUsernamePassword struct {
	// Add Password for authentication.
	Password string                                                    `json:"password"`
	type_    BothUsernameAndPasswordIsRequiredForAuthenticationRequest `const:"username_password" json:"type"`
	// Add Username for authentication.
	Username string `json:"username"`
}

func (s SourceOutbrainAmplifyUpdateUsernamePassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOutbrainAmplifyUpdateUsernamePassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOutbrainAmplifyUpdateUsernamePassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceOutbrainAmplifyUpdateUsernamePassword) GetType() BothUsernameAndPasswordIsRequiredForAuthenticationRequest {
	return BothUsernameAndPasswordIsRequiredForAuthenticationRequestUsernamePassword
}

func (o *SourceOutbrainAmplifyUpdateUsernamePassword) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type AccessTokenIsRequiredForAuthenticationRequests string

const (
	AccessTokenIsRequiredForAuthenticationRequestsAccessToken AccessTokenIsRequiredForAuthenticationRequests = "access_token"
)

func (e AccessTokenIsRequiredForAuthenticationRequests) ToPointer() *AccessTokenIsRequiredForAuthenticationRequests {
	return &e
}

func (e *AccessTokenIsRequiredForAuthenticationRequests) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = AccessTokenIsRequiredForAuthenticationRequests(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccessTokenIsRequiredForAuthenticationRequests: %v", v)
	}
}

// SourceOutbrainAmplifyUpdateAccessToken - Credentials for making authenticated requests requires either username/password or access_token.
type SourceOutbrainAmplifyUpdateAccessToken struct {
	// Access Token for making authenticated requests.
	AccessToken string                                         `json:"access_token"`
	type_       AccessTokenIsRequiredForAuthenticationRequests `const:"access_token" json:"type"`
}

func (s SourceOutbrainAmplifyUpdateAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOutbrainAmplifyUpdateAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOutbrainAmplifyUpdateAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceOutbrainAmplifyUpdateAccessToken) GetType() AccessTokenIsRequiredForAuthenticationRequests {
	return AccessTokenIsRequiredForAuthenticationRequestsAccessToken
}

type SourceOutbrainAmplifyUpdateAuthenticationMethodType string

const (
	SourceOutbrainAmplifyUpdateAuthenticationMethodTypeAccessToken      SourceOutbrainAmplifyUpdateAuthenticationMethodType = "AccessToken"
	SourceOutbrainAmplifyUpdateAuthenticationMethodTypeUsernamePassword SourceOutbrainAmplifyUpdateAuthenticationMethodType = "UsernamePassword"
)

type SourceOutbrainAmplifyUpdateAuthenticationMethod struct {
	AccessToken      *SourceOutbrainAmplifyUpdateAccessToken
	UsernamePassword *SourceOutbrainAmplifyUpdateUsernamePassword

	Type SourceOutbrainAmplifyUpdateAuthenticationMethodType
}

func CreateSourceOutbrainAmplifyUpdateAuthenticationMethodAccessToken(accessToken SourceOutbrainAmplifyUpdateAccessToken) SourceOutbrainAmplifyUpdateAuthenticationMethod {
	typ := SourceOutbrainAmplifyUpdateAuthenticationMethodTypeAccessToken

	return SourceOutbrainAmplifyUpdateAuthenticationMethod{
		AccessToken: &accessToken,
		Type:        typ,
	}
}

func CreateSourceOutbrainAmplifyUpdateAuthenticationMethodUsernamePassword(usernamePassword SourceOutbrainAmplifyUpdateUsernamePassword) SourceOutbrainAmplifyUpdateAuthenticationMethod {
	typ := SourceOutbrainAmplifyUpdateAuthenticationMethodTypeUsernamePassword

	return SourceOutbrainAmplifyUpdateAuthenticationMethod{
		UsernamePassword: &usernamePassword,
		Type:             typ,
	}
}

func (u *SourceOutbrainAmplifyUpdateAuthenticationMethod) UnmarshalJSON(data []byte) error {

	accessToken := new(SourceOutbrainAmplifyUpdateAccessToken)
	if err := utils.UnmarshalJSON(data, &accessToken, "", true, true); err == nil {
		u.AccessToken = accessToken
		u.Type = SourceOutbrainAmplifyUpdateAuthenticationMethodTypeAccessToken
		return nil
	}

	usernamePassword := new(SourceOutbrainAmplifyUpdateUsernamePassword)
	if err := utils.UnmarshalJSON(data, &usernamePassword, "", true, true); err == nil {
		u.UsernamePassword = usernamePassword
		u.Type = SourceOutbrainAmplifyUpdateAuthenticationMethodTypeUsernamePassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOutbrainAmplifyUpdateAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.AccessToken != nil {
		return utils.MarshalJSON(u.AccessToken, "", true)
	}

	if u.UsernamePassword != nil {
		return utils.MarshalJSON(u.UsernamePassword, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GranularityForGeoLocationRegion - The granularity used for geo location data in reports.
type GranularityForGeoLocationRegion string

const (
	GranularityForGeoLocationRegionCountry   GranularityForGeoLocationRegion = "country"
	GranularityForGeoLocationRegionRegion    GranularityForGeoLocationRegion = "region"
	GranularityForGeoLocationRegionSubregion GranularityForGeoLocationRegion = "subregion"
)

func (e GranularityForGeoLocationRegion) ToPointer() *GranularityForGeoLocationRegion {
	return &e
}

func (e *GranularityForGeoLocationRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "country":
		fallthrough
	case "region":
		fallthrough
	case "subregion":
		*e = GranularityForGeoLocationRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GranularityForGeoLocationRegion: %v", v)
	}
}

// GranularityForPeriodicReports - The granularity used for periodic data in reports. See <a href="https://amplifyv01.docs.apiary.io/#reference/performance-reporting/periodic/retrieve-performance-statistics-for-all-marketer-campaigns-by-periodic-breakdown">the docs</a>.
type GranularityForPeriodicReports string

const (
	GranularityForPeriodicReportsDaily   GranularityForPeriodicReports = "daily"
	GranularityForPeriodicReportsWeekly  GranularityForPeriodicReports = "weekly"
	GranularityForPeriodicReportsMonthly GranularityForPeriodicReports = "monthly"
)

func (e GranularityForPeriodicReports) ToPointer() *GranularityForPeriodicReports {
	return &e
}

func (e *GranularityForPeriodicReports) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "weekly":
		fallthrough
	case "monthly":
		*e = GranularityForPeriodicReports(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GranularityForPeriodicReports: %v", v)
	}
}

type SourceOutbrainAmplifyUpdate struct {
	// Credentials for making authenticated requests requires either username/password or access_token.
	Credentials SourceOutbrainAmplifyUpdateAuthenticationMethod `json:"credentials"`
	// Date in the format YYYY-MM-DD.
	EndDate *string `json:"end_date,omitempty"`
	// The granularity used for geo location data in reports.
	GeoLocationBreakdown *GranularityForGeoLocationRegion `json:"geo_location_breakdown,omitempty"`
	// The granularity used for periodic data in reports. See <a href="https://amplifyv01.docs.apiary.io/#reference/performance-reporting/periodic/retrieve-performance-statistics-for-all-marketer-campaigns-by-periodic-breakdown">the docs</a>.
	ReportGranularity *GranularityForPeriodicReports `json:"report_granularity,omitempty"`
	// Date in the format YYYY-MM-DD eg. 2017-01-25. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}

func (o *SourceOutbrainAmplifyUpdate) GetCredentials() SourceOutbrainAmplifyUpdateAuthenticationMethod {
	if o == nil {
		return SourceOutbrainAmplifyUpdateAuthenticationMethod{}
	}
	return o.Credentials
}

func (o *SourceOutbrainAmplifyUpdate) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceOutbrainAmplifyUpdate) GetGeoLocationBreakdown() *GranularityForGeoLocationRegion {
	if o == nil {
		return nil
	}
	return o.GeoLocationBreakdown
}

func (o *SourceOutbrainAmplifyUpdate) GetReportGranularity() *GranularityForPeriodicReports {
	if o == nil {
		return nil
	}
	return o.ReportGranularity
}

func (o *SourceOutbrainAmplifyUpdate) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
