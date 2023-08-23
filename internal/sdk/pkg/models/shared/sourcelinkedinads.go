// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy - Select value from list to pivot by
type SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy string

const (
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByCompany                     SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "COMPANY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByAccount                     SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "ACCOUNT"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByShare                       SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "SHARE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByCampaign                    SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CAMPAIGN"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByCreative                    SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CREATIVE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByCampaignGroup               SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CAMPAIGN_GROUP"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByConversion                  SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CONVERSION"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByConversationNode            SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CONVERSATION_NODE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByConversationNodeOptionIndex SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CONVERSATION_NODE_OPTION_INDEX"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByServingLocation             SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "SERVING_LOCATION"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByCardIndex                   SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "CARD_INDEX"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberCompanySize           SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_COMPANY_SIZE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberIndustry              SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_INDUSTRY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberSeniority             SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_SENIORITY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberJobTitle              SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_JOB_TITLE "
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberJobFunction           SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_JOB_FUNCTION "
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberCountryV2             SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_COUNTRY_V2 "
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberRegionV2              SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_REGION_V2"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByMemberCompany               SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "MEMBER_COMPANY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByPlacementName               SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "PLACEMENT_NAME"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotByImpressionDeviceType        SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy = "IMPRESSION_DEVICE_TYPE"
)

func (e SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy) ToPointer() *SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy {
	return &e
}

func (e *SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMPANY":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "SHARE":
		fallthrough
	case "CAMPAIGN":
		fallthrough
	case "CREATIVE":
		fallthrough
	case "CAMPAIGN_GROUP":
		fallthrough
	case "CONVERSION":
		fallthrough
	case "CONVERSATION_NODE":
		fallthrough
	case "CONVERSATION_NODE_OPTION_INDEX":
		fallthrough
	case "SERVING_LOCATION":
		fallthrough
	case "CARD_INDEX":
		fallthrough
	case "MEMBER_COMPANY_SIZE":
		fallthrough
	case "MEMBER_INDUSTRY":
		fallthrough
	case "MEMBER_SENIORITY":
		fallthrough
	case "MEMBER_JOB_TITLE ":
		fallthrough
	case "MEMBER_JOB_FUNCTION ":
		fallthrough
	case "MEMBER_COUNTRY_V2 ":
		fallthrough
	case "MEMBER_REGION_V2":
		fallthrough
	case "MEMBER_COMPANY":
		fallthrough
	case "PLACEMENT_NAME":
		fallthrough
	case "IMPRESSION_DEVICE_TYPE":
		*e = SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy: %v", v)
	}
}

// SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity - Set time granularity for report:
// ALL - Results grouped into a single result across the entire time range of the report.
// DAILY - Results grouped by day.
// MONTHLY - Results grouped by month.
// YEARLY - Results grouped by year.
type SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity string

const (
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityAll     SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "ALL"
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityDaily   SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "DAILY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityMonthly SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "MONTHLY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityYearly  SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "YEARLY"
)

func (e SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity) ToPointer() *SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity {
	return &e
}

func (e *SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "DAILY":
		fallthrough
	case "MONTHLY":
		fallthrough
	case "YEARLY":
		*e = SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity: %v", v)
	}
}

// SourceLinkedinAdsAdAnalyticsReportConfiguration - Config for custom ad Analytics Report
type SourceLinkedinAdsAdAnalyticsReportConfiguration struct {
	// The name for the report
	Name string `json:"name"`
	// Select value from list to pivot by
	PivotBy SourceLinkedinAdsAdAnalyticsReportConfigurationPivotBy `json:"pivot_by"`
	// Set time granularity for report:
	// ALL - Results grouped into a single result across the entire time range of the report.
	// DAILY - Results grouped by day.
	// MONTHLY - Results grouped by month.
	// YEARLY - Results grouped by year.
	TimeGranularity SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity `json:"time_granularity"`
}

type SourceLinkedinAdsAuthenticationAccessTokenAuthMethod string

const (
	SourceLinkedinAdsAuthenticationAccessTokenAuthMethodAccessToken SourceLinkedinAdsAuthenticationAccessTokenAuthMethod = "access_token"
)

func (e SourceLinkedinAdsAuthenticationAccessTokenAuthMethod) ToPointer() *SourceLinkedinAdsAuthenticationAccessTokenAuthMethod {
	return &e
}

func (e *SourceLinkedinAdsAuthenticationAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinAdsAuthenticationAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAuthenticationAccessTokenAuthMethod: %v", v)
	}
}

type SourceLinkedinAdsAuthenticationAccessToken struct {
	// The token value generated using the authentication code. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-ads#authentication">docs</a> to obtain yours.
	AccessToken string                                                `json:"access_token"`
	AuthMethod  *SourceLinkedinAdsAuthenticationAccessTokenAuthMethod `json:"auth_method,omitempty"`
}

type SourceLinkedinAdsAuthenticationOAuth20AuthMethod string

const (
	SourceLinkedinAdsAuthenticationOAuth20AuthMethodOAuth20 SourceLinkedinAdsAuthenticationOAuth20AuthMethod = "oAuth2.0"
)

func (e SourceLinkedinAdsAuthenticationOAuth20AuthMethod) ToPointer() *SourceLinkedinAdsAuthenticationOAuth20AuthMethod {
	return &e
}

func (e *SourceLinkedinAdsAuthenticationOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinAdsAuthenticationOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAuthenticationOAuth20AuthMethod: %v", v)
	}
}

type SourceLinkedinAdsAuthenticationOAuth20 struct {
	AuthMethod *SourceLinkedinAdsAuthenticationOAuth20AuthMethod `json:"auth_method,omitempty"`
	// The client ID of the LinkedIn Ads developer application.
	ClientID string `json:"client_id"`
	// The client secret the LinkedIn Ads developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token.
	RefreshToken string `json:"refresh_token"`
}

type SourceLinkedinAdsAuthenticationType string

const (
	SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationOAuth20     SourceLinkedinAdsAuthenticationType = "source-linkedin-ads_Authentication_OAuth2.0"
	SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationAccessToken SourceLinkedinAdsAuthenticationType = "source-linkedin-ads_Authentication_Access token"
)

type SourceLinkedinAdsAuthentication struct {
	SourceLinkedinAdsAuthenticationOAuth20     *SourceLinkedinAdsAuthenticationOAuth20
	SourceLinkedinAdsAuthenticationAccessToken *SourceLinkedinAdsAuthenticationAccessToken

	Type SourceLinkedinAdsAuthenticationType
}

func CreateSourceLinkedinAdsAuthenticationSourceLinkedinAdsAuthenticationOAuth20(sourceLinkedinAdsAuthenticationOAuth20 SourceLinkedinAdsAuthenticationOAuth20) SourceLinkedinAdsAuthentication {
	typ := SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationOAuth20

	return SourceLinkedinAdsAuthentication{
		SourceLinkedinAdsAuthenticationOAuth20: &sourceLinkedinAdsAuthenticationOAuth20,
		Type:                                   typ,
	}
}

func CreateSourceLinkedinAdsAuthenticationSourceLinkedinAdsAuthenticationAccessToken(sourceLinkedinAdsAuthenticationAccessToken SourceLinkedinAdsAuthenticationAccessToken) SourceLinkedinAdsAuthentication {
	typ := SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationAccessToken

	return SourceLinkedinAdsAuthentication{
		SourceLinkedinAdsAuthenticationAccessToken: &sourceLinkedinAdsAuthenticationAccessToken,
		Type: typ,
	}
}

func (u *SourceLinkedinAdsAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceLinkedinAdsAuthenticationOAuth20 := new(SourceLinkedinAdsAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinAdsAuthenticationOAuth20); err == nil {
		u.SourceLinkedinAdsAuthenticationOAuth20 = sourceLinkedinAdsAuthenticationOAuth20
		u.Type = SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationOAuth20
		return nil
	}

	sourceLinkedinAdsAuthenticationAccessToken := new(SourceLinkedinAdsAuthenticationAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinAdsAuthenticationAccessToken); err == nil {
		u.SourceLinkedinAdsAuthenticationAccessToken = sourceLinkedinAdsAuthenticationAccessToken
		u.Type = SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLinkedinAdsAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinAdsAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceLinkedinAdsAuthenticationOAuth20)
	}

	if u.SourceLinkedinAdsAuthenticationAccessToken != nil {
		return json.Marshal(u.SourceLinkedinAdsAuthenticationAccessToken)
	}

	return nil, nil
}

type SourceLinkedinAdsLinkedinAds string

const (
	SourceLinkedinAdsLinkedinAdsLinkedinAds SourceLinkedinAdsLinkedinAds = "linkedin-ads"
)

func (e SourceLinkedinAdsLinkedinAds) ToPointer() *SourceLinkedinAdsLinkedinAds {
	return &e
}

func (e *SourceLinkedinAdsLinkedinAds) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linkedin-ads":
		*e = SourceLinkedinAdsLinkedinAds(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsLinkedinAds: %v", v)
	}
}

type SourceLinkedinAds struct {
	// Specify the account IDs separated by a space, to pull the data from. Leave empty, if you want to pull the data from all associated accounts. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn Ads docs</a> for more info.
	AccountIds         []int64                                           `json:"account_ids,omitempty"`
	AdAnalyticsReports []SourceLinkedinAdsAdAnalyticsReportConfiguration `json:"ad_analytics_reports,omitempty"`
	Credentials        *SourceLinkedinAdsAuthentication                  `json:"credentials,omitempty"`
	SourceType         SourceLinkedinAdsLinkedinAds                      `json:"sourceType"`
	// UTC date in the format 2020-09-17. Any data before this date will not be replicated.
	StartDate types.Date `json:"start_date"`
}
