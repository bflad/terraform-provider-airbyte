// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type AuthType string

const (
	AuthTypeClient AuthType = "Client"
)

func (e AuthType) ToPointer() *AuthType {
	return &e
}

func (e *AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthType: %v", v)
	}
}

type SearchCriteria string

const (
	SearchCriteriaStartsWith    SearchCriteria = "starts with"
	SearchCriteriaEndsWith      SearchCriteria = "ends with"
	SearchCriteriaContains      SearchCriteria = "contains"
	SearchCriteriaExacts        SearchCriteria = "exacts"
	SearchCriteriaStartsNotWith SearchCriteria = "starts not with"
	SearchCriteriaEndsNotWith   SearchCriteria = "ends not with"
	SearchCriteriaNotContains   SearchCriteria = "not contains"
	SearchCriteriaNotExacts     SearchCriteria = "not exacts"
)

func (e SearchCriteria) ToPointer() *SearchCriteria {
	return &e
}

func (e *SearchCriteria) UnmarshalJSON(data []byte) error {
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
		*e = SearchCriteria(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SearchCriteria: %v", v)
	}
}

type StreamsCriteria struct {
	Criteria *SearchCriteria `default:"contains" json:"criteria"`
	Value    string          `json:"value"`
}

func (s StreamsCriteria) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StreamsCriteria) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StreamsCriteria) GetCriteria() *SearchCriteria {
	if o == nil {
		return nil
	}
	return o.Criteria
}

func (o *StreamsCriteria) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type SourceSalesforceUpdate struct {
	authType *AuthType `const:"Client" json:"auth_type,omitempty"`
	// Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client ID</a>
	ClientID string `json:"client_id"`
	// Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client secret</a>
	ClientSecret string `json:"client_secret"`
	// Toggle to use Bulk API (this might cause empty fields for some streams)
	ForceUseBulkAPI *bool `default:"false" json:"force_use_bulk_api"`
	// Toggle if you're using a <a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&type=5">Salesforce Sandbox</a>
	IsSandbox *bool `default:"false" json:"is_sandbox"`
	// Enter your application's <a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm">Salesforce Refresh Token</a> used for Airbyte to access your Salesforce account.
	RefreshToken string `json:"refresh_token"`
	// Enter the date (or date-time) in the YYYY-MM-DD or YYYY-MM-DDTHH:mm:ssZ format. Airbyte will replicate the data updated on and after this date. If this field is blank, Airbyte will replicate the data for last two years.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Add filters to select only required stream based on `SObject` name. Use this field to filter which tables are displayed by this connector. This is useful if your Salesforce account has a large number of tables (>1000), in which case you may find it easier to navigate the UI and speed up the connector's performance if you restrict the tables displayed by this connector.
	StreamsCriteria []StreamsCriteria `json:"streams_criteria,omitempty"`
}

func (s SourceSalesforceUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSalesforceUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSalesforceUpdate) GetAuthType() *AuthType {
	return AuthTypeClient.ToPointer()
}

func (o *SourceSalesforceUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSalesforceUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSalesforceUpdate) GetForceUseBulkAPI() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUseBulkAPI
}

func (o *SourceSalesforceUpdate) GetIsSandbox() *bool {
	if o == nil {
		return nil
	}
	return o.IsSandbox
}

func (o *SourceSalesforceUpdate) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceSalesforceUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceSalesforceUpdate) GetStreamsCriteria() []StreamsCriteria {
	if o == nil {
		return nil
	}
	return o.StreamsCriteria
}
