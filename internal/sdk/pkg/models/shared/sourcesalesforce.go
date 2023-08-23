// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceSalesforceAuthType string

const (
	SourceSalesforceAuthTypeClient SourceSalesforceAuthType = "Client"
)

func (e SourceSalesforceAuthType) ToPointer() *SourceSalesforceAuthType {
	return &e
}

func (e *SourceSalesforceAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceSalesforceAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceAuthType: %v", v)
	}
}

type SourceSalesforceSalesforce string

const (
	SourceSalesforceSalesforceSalesforce SourceSalesforceSalesforce = "salesforce"
)

func (e SourceSalesforceSalesforce) ToPointer() *SourceSalesforceSalesforce {
	return &e
}

func (e *SourceSalesforceSalesforce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "salesforce":
		*e = SourceSalesforceSalesforce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceSalesforce: %v", v)
	}
}

type SourceSalesforceStreamsCriteriaSearchCriteria string

const (
	SourceSalesforceStreamsCriteriaSearchCriteriaStartsWith    SourceSalesforceStreamsCriteriaSearchCriteria = "starts with"
	SourceSalesforceStreamsCriteriaSearchCriteriaEndsWith      SourceSalesforceStreamsCriteriaSearchCriteria = "ends with"
	SourceSalesforceStreamsCriteriaSearchCriteriaContains      SourceSalesforceStreamsCriteriaSearchCriteria = "contains"
	SourceSalesforceStreamsCriteriaSearchCriteriaExacts        SourceSalesforceStreamsCriteriaSearchCriteria = "exacts"
	SourceSalesforceStreamsCriteriaSearchCriteriaStartsNotWith SourceSalesforceStreamsCriteriaSearchCriteria = "starts not with"
	SourceSalesforceStreamsCriteriaSearchCriteriaEndsNotWith   SourceSalesforceStreamsCriteriaSearchCriteria = "ends not with"
	SourceSalesforceStreamsCriteriaSearchCriteriaNotContains   SourceSalesforceStreamsCriteriaSearchCriteria = "not contains"
	SourceSalesforceStreamsCriteriaSearchCriteriaNotExacts     SourceSalesforceStreamsCriteriaSearchCriteria = "not exacts"
)

func (e SourceSalesforceStreamsCriteriaSearchCriteria) ToPointer() *SourceSalesforceStreamsCriteriaSearchCriteria {
	return &e
}

func (e *SourceSalesforceStreamsCriteriaSearchCriteria) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "starts with":
		fallthrough
	case "ends with":
		fallthrough
	case "contains":
		fallthrough
	case "exacts":
		fallthrough
	case "starts not with":
		fallthrough
	case "ends not with":
		fallthrough
	case "not contains":
		fallthrough
	case "not exacts":
		*e = SourceSalesforceStreamsCriteriaSearchCriteria(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceStreamsCriteriaSearchCriteria: %v", v)
	}
}

type SourceSalesforceStreamsCriteria struct {
	Criteria SourceSalesforceStreamsCriteriaSearchCriteria `json:"criteria"`
	Value    string                                        `json:"value"`
}

type SourceSalesforce struct {
	AuthType *SourceSalesforceAuthType `json:"auth_type,omitempty"`
	// Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client ID</a>
	ClientID string `json:"client_id"`
	// Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client secret</a>
	ClientSecret string `json:"client_secret"`
	// Toggle to use Bulk API (this might cause empty fields for some streams)
	ForceUseBulkAPI *bool `json:"force_use_bulk_api,omitempty"`
	// Toggle if you're using a <a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&type=5">Salesforce Sandbox</a>
	IsSandbox *bool `json:"is_sandbox,omitempty"`
	// Enter your application's <a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm">Salesforce Refresh Token</a> used for Airbyte to access your Salesforce account.
	RefreshToken string                     `json:"refresh_token"`
	SourceType   SourceSalesforceSalesforce `json:"sourceType"`
	// Enter the date (or date-time) in the YYYY-MM-DD or YYYY-MM-DDTHH:mm:ssZ format. Airbyte will replicate the data updated on and after this date. If this field is blank, Airbyte will replicate the data for last two years.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Add filters to select only required stream based on `SObject` name. Use this field to filter which tables are displayed by this connector. This is useful if your Salesforce account has a large number of tables (>1000), in which case you may find it easier to navigate the UI and speed up the connector's performance if you restrict the tables displayed by this connector.
	StreamsCriteria []SourceSalesforceStreamsCriteria `json:"streams_criteria,omitempty"`
}
