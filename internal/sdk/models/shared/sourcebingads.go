// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
)

// SourceBingAdsOperator - An Operator that will be used to filter accounts. The Contains predicate has features for matching words, matching inflectional forms of words, searching using wildcard characters, and searching using proximity. The Equals is used to return all rows where account name is equal(=) to the string that you provided
type SourceBingAdsOperator string

const (
	SourceBingAdsOperatorContains SourceBingAdsOperator = "Contains"
	SourceBingAdsOperatorEquals   SourceBingAdsOperator = "Equals"
)

func (e SourceBingAdsOperator) ToPointer() *SourceBingAdsOperator {
	return &e
}
func (e *SourceBingAdsOperator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Contains":
		fallthrough
	case "Equals":
		*e = SourceBingAdsOperator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBingAdsOperator: %v", v)
	}
}

// SourceBingAdsAccountNames - Account Names Predicates Config.
type SourceBingAdsAccountNames struct {
	// Account Name is a string value for comparing with the specified predicate.
	Name string `json:"name"`
	// An Operator that will be used to filter accounts. The Contains predicate has features for matching words, matching inflectional forms of words, searching using wildcard characters, and searching using proximity. The Equals is used to return all rows where account name is equal(=) to the string that you provided
	Operator SourceBingAdsOperator `json:"operator"`
}

func (o *SourceBingAdsAccountNames) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceBingAdsAccountNames) GetOperator() SourceBingAdsOperator {
	if o == nil {
		return SourceBingAdsOperator("")
	}
	return o.Operator
}

type SourceBingAdsAuthMethod string

const (
	SourceBingAdsAuthMethodOauth20 SourceBingAdsAuthMethod = "oauth2.0"
)

func (e SourceBingAdsAuthMethod) ToPointer() *SourceBingAdsAuthMethod {
	return &e
}
func (e *SourceBingAdsAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceBingAdsAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBingAdsAuthMethod: %v", v)
	}
}

// SourceBingAdsReportingDataObject - The name of the the object derives from the ReportRequest object. You can find it in Bing Ads Api docs - Reporting API - Reporting Data Objects.
type SourceBingAdsReportingDataObject string

const (
	SourceBingAdsReportingDataObjectAccountPerformanceReportRequest               SourceBingAdsReportingDataObject = "AccountPerformanceReportRequest"
	SourceBingAdsReportingDataObjectAdDynamicTextPerformanceReportRequest         SourceBingAdsReportingDataObject = "AdDynamicTextPerformanceReportRequest"
	SourceBingAdsReportingDataObjectAdExtensionByAdReportRequest                  SourceBingAdsReportingDataObject = "AdExtensionByAdReportRequest"
	SourceBingAdsReportingDataObjectAdExtensionByKeywordReportRequest             SourceBingAdsReportingDataObject = "AdExtensionByKeywordReportRequest"
	SourceBingAdsReportingDataObjectAdExtensionDetailReportRequest                SourceBingAdsReportingDataObject = "AdExtensionDetailReportRequest"
	SourceBingAdsReportingDataObjectAdGroupPerformanceReportRequest               SourceBingAdsReportingDataObject = "AdGroupPerformanceReportRequest"
	SourceBingAdsReportingDataObjectAdPerformanceReportRequest                    SourceBingAdsReportingDataObject = "AdPerformanceReportRequest"
	SourceBingAdsReportingDataObjectAgeGenderAudienceReportRequest                SourceBingAdsReportingDataObject = "AgeGenderAudienceReportRequest"
	SourceBingAdsReportingDataObjectAudiencePerformanceReportRequest              SourceBingAdsReportingDataObject = "AudiencePerformanceReportRequest"
	SourceBingAdsReportingDataObjectCallDetailReportRequest                       SourceBingAdsReportingDataObject = "CallDetailReportRequest"
	SourceBingAdsReportingDataObjectCampaignPerformanceReportRequest              SourceBingAdsReportingDataObject = "CampaignPerformanceReportRequest"
	SourceBingAdsReportingDataObjectConversionPerformanceReportRequest            SourceBingAdsReportingDataObject = "ConversionPerformanceReportRequest"
	SourceBingAdsReportingDataObjectDestinationURLPerformanceReportRequest        SourceBingAdsReportingDataObject = "DestinationUrlPerformanceReportRequest"
	SourceBingAdsReportingDataObjectDsaAutoTargetPerformanceReportRequest         SourceBingAdsReportingDataObject = "DSAAutoTargetPerformanceReportRequest"
	SourceBingAdsReportingDataObjectDsaCategoryPerformanceReportRequest           SourceBingAdsReportingDataObject = "DSACategoryPerformanceReportRequest"
	SourceBingAdsReportingDataObjectDsaSearchQueryPerformanceReportRequest        SourceBingAdsReportingDataObject = "DSASearchQueryPerformanceReportRequest"
	SourceBingAdsReportingDataObjectGeographicPerformanceReportRequest            SourceBingAdsReportingDataObject = "GeographicPerformanceReportRequest"
	SourceBingAdsReportingDataObjectGoalsAndFunnelsReportRequest                  SourceBingAdsReportingDataObject = "GoalsAndFunnelsReportRequest"
	SourceBingAdsReportingDataObjectHotelDimensionPerformanceReportRequest        SourceBingAdsReportingDataObject = "HotelDimensionPerformanceReportRequest"
	SourceBingAdsReportingDataObjectHotelGroupPerformanceReportRequest            SourceBingAdsReportingDataObject = "HotelGroupPerformanceReportRequest"
	SourceBingAdsReportingDataObjectKeywordPerformanceReportRequest               SourceBingAdsReportingDataObject = "KeywordPerformanceReportRequest"
	SourceBingAdsReportingDataObjectNegativeKeywordConflictReportRequest          SourceBingAdsReportingDataObject = "NegativeKeywordConflictReportRequest"
	SourceBingAdsReportingDataObjectProductDimensionPerformanceReportRequest      SourceBingAdsReportingDataObject = "ProductDimensionPerformanceReportRequest"
	SourceBingAdsReportingDataObjectProductMatchCountReportRequest                SourceBingAdsReportingDataObject = "ProductMatchCountReportRequest"
	SourceBingAdsReportingDataObjectProductNegativeKeywordConflictReportRequest   SourceBingAdsReportingDataObject = "ProductNegativeKeywordConflictReportRequest"
	SourceBingAdsReportingDataObjectProductPartitionPerformanceReportRequest      SourceBingAdsReportingDataObject = "ProductPartitionPerformanceReportRequest"
	SourceBingAdsReportingDataObjectProductPartitionUnitPerformanceReportRequest  SourceBingAdsReportingDataObject = "ProductPartitionUnitPerformanceReportRequest"
	SourceBingAdsReportingDataObjectProductSearchQueryPerformanceReportRequest    SourceBingAdsReportingDataObject = "ProductSearchQueryPerformanceReportRequest"
	SourceBingAdsReportingDataObjectProfessionalDemographicsAudienceReportRequest SourceBingAdsReportingDataObject = "ProfessionalDemographicsAudienceReportRequest"
	SourceBingAdsReportingDataObjectPublisherUsagePerformanceReportRequest        SourceBingAdsReportingDataObject = "PublisherUsagePerformanceReportRequest"
	SourceBingAdsReportingDataObjectSearchCampaignChangeHistoryReportRequest      SourceBingAdsReportingDataObject = "SearchCampaignChangeHistoryReportRequest"
	SourceBingAdsReportingDataObjectSearchQueryPerformanceReportRequest           SourceBingAdsReportingDataObject = "SearchQueryPerformanceReportRequest"
	SourceBingAdsReportingDataObjectShareOfVoiceReportRequest                     SourceBingAdsReportingDataObject = "ShareOfVoiceReportRequest"
	SourceBingAdsReportingDataObjectUserLocationPerformanceReportRequest          SourceBingAdsReportingDataObject = "UserLocationPerformanceReportRequest"
)

func (e SourceBingAdsReportingDataObject) ToPointer() *SourceBingAdsReportingDataObject {
	return &e
}
func (e *SourceBingAdsReportingDataObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AccountPerformanceReportRequest":
		fallthrough
	case "AdDynamicTextPerformanceReportRequest":
		fallthrough
	case "AdExtensionByAdReportRequest":
		fallthrough
	case "AdExtensionByKeywordReportRequest":
		fallthrough
	case "AdExtensionDetailReportRequest":
		fallthrough
	case "AdGroupPerformanceReportRequest":
		fallthrough
	case "AdPerformanceReportRequest":
		fallthrough
	case "AgeGenderAudienceReportRequest":
		fallthrough
	case "AudiencePerformanceReportRequest":
		fallthrough
	case "CallDetailReportRequest":
		fallthrough
	case "CampaignPerformanceReportRequest":
		fallthrough
	case "ConversionPerformanceReportRequest":
		fallthrough
	case "DestinationUrlPerformanceReportRequest":
		fallthrough
	case "DSAAutoTargetPerformanceReportRequest":
		fallthrough
	case "DSACategoryPerformanceReportRequest":
		fallthrough
	case "DSASearchQueryPerformanceReportRequest":
		fallthrough
	case "GeographicPerformanceReportRequest":
		fallthrough
	case "GoalsAndFunnelsReportRequest":
		fallthrough
	case "HotelDimensionPerformanceReportRequest":
		fallthrough
	case "HotelGroupPerformanceReportRequest":
		fallthrough
	case "KeywordPerformanceReportRequest":
		fallthrough
	case "NegativeKeywordConflictReportRequest":
		fallthrough
	case "ProductDimensionPerformanceReportRequest":
		fallthrough
	case "ProductMatchCountReportRequest":
		fallthrough
	case "ProductNegativeKeywordConflictReportRequest":
		fallthrough
	case "ProductPartitionPerformanceReportRequest":
		fallthrough
	case "ProductPartitionUnitPerformanceReportRequest":
		fallthrough
	case "ProductSearchQueryPerformanceReportRequest":
		fallthrough
	case "ProfessionalDemographicsAudienceReportRequest":
		fallthrough
	case "PublisherUsagePerformanceReportRequest":
		fallthrough
	case "SearchCampaignChangeHistoryReportRequest":
		fallthrough
	case "SearchQueryPerformanceReportRequest":
		fallthrough
	case "ShareOfVoiceReportRequest":
		fallthrough
	case "UserLocationPerformanceReportRequest":
		*e = SourceBingAdsReportingDataObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBingAdsReportingDataObject: %v", v)
	}
}

type SourceBingAdsCustomReportConfig struct {
	// The name of the custom report, this name would be used as stream name
	Name string `json:"name"`
	// A list of available aggregations.
	ReportAggregation *string `default:"[Hourly]" json:"report_aggregation"`
	// A list of available report object columns. You can find it in description of reporting object that you want to add to custom report.
	ReportColumns []string `json:"report_columns"`
	// The name of the the object derives from the ReportRequest object. You can find it in Bing Ads Api docs - Reporting API - Reporting Data Objects.
	ReportingObject SourceBingAdsReportingDataObject `json:"reporting_object"`
}

func (s SourceBingAdsCustomReportConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBingAdsCustomReportConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBingAdsCustomReportConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceBingAdsCustomReportConfig) GetReportAggregation() *string {
	if o == nil {
		return nil
	}
	return o.ReportAggregation
}

func (o *SourceBingAdsCustomReportConfig) GetReportColumns() []string {
	if o == nil {
		return []string{}
	}
	return o.ReportColumns
}

func (o *SourceBingAdsCustomReportConfig) GetReportingObject() SourceBingAdsReportingDataObject {
	if o == nil {
		return SourceBingAdsReportingDataObject("")
	}
	return o.ReportingObject
}

type BingAds string

const (
	BingAdsBingAds BingAds = "bing-ads"
)

func (e BingAds) ToPointer() *BingAds {
	return &e
}
func (e *BingAds) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bing-ads":
		*e = BingAds(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BingAds: %v", v)
	}
}

type SourceBingAds struct {
	// Predicates that will be used to sync data by specific accounts.
	AccountNames []SourceBingAdsAccountNames `json:"account_names,omitempty"`
	authMethod   *SourceBingAdsAuthMethod    `const:"oauth2.0" json:"auth_method,omitempty"`
	// The Client ID of your Microsoft Advertising developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Advertising developer application.
	ClientSecret *string `default:"" json:"client_secret"`
	// You can add your Custom Bing Ads report by creating one.
	CustomReports []SourceBingAdsCustomReportConfig `json:"custom_reports,omitempty"`
	// Developer token associated with user. See more info <a href="https://docs.microsoft.com/en-us/advertising/guides/get-started?view=bingads-13#get-developer-token"> in the docs</a>.
	DeveloperToken string `json:"developer_token"`
	// Also known as attribution or conversion window. How far into the past to look for records (in days). If your conversion window has an hours/minutes granularity, round it up to the number of days exceeding. Used only for performance report streams in incremental mode without specified Reports Start Date.
	LookbackWindow *int64 `default:"0" json:"lookback_window"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// The start date from which to begin replicating report data. Any data generated before this date will not be replicated in reports. This is a UTC date in YYYY-MM-DD format. If not set, data from previous and current calendar year will be replicated.
	ReportsStartDate *types.Date `json:"reports_start_date,omitempty"`
	sourceType       BingAds     `const:"bing-ads" json:"sourceType"`
	// The Tenant ID of your Microsoft Advertising developer application. Set this to "common" unless you know you need a different value.
	TenantID *string `default:"common" json:"tenant_id"`
}

func (s SourceBingAds) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBingAds) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBingAds) GetAccountNames() []SourceBingAdsAccountNames {
	if o == nil {
		return nil
	}
	return o.AccountNames
}

func (o *SourceBingAds) GetAuthMethod() *SourceBingAdsAuthMethod {
	return SourceBingAdsAuthMethodOauth20.ToPointer()
}

func (o *SourceBingAds) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceBingAds) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *SourceBingAds) GetCustomReports() []SourceBingAdsCustomReportConfig {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourceBingAds) GetDeveloperToken() string {
	if o == nil {
		return ""
	}
	return o.DeveloperToken
}

func (o *SourceBingAds) GetLookbackWindow() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindow
}

func (o *SourceBingAds) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceBingAds) GetReportsStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReportsStartDate
}

func (o *SourceBingAds) GetSourceType() BingAds {
	return BingAdsBingAds
}

func (o *SourceBingAds) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}
