// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceOrbUpdate struct {
	// Orb API Key, issued from the Orb admin console.
	APIKey string `json:"api_key"`
	// When set to N, the connector will always refresh resources created within the past N days. By default, updated objects that are not newly created are not incrementally synced.
	LookbackWindowDays *int64 `default:"0" json:"lookback_window_days"`
	// Property key names to extract from all events, in order to enrich ledger entries corresponding to an event deduction.
	NumericEventPropertiesKeys []string `json:"numeric_event_properties_keys,omitempty"`
	// Orb Plan ID to filter subscriptions that should have usage fetched.
	PlanID *string `json:"plan_id,omitempty"`
	// UTC date and time in the format 2022-03-01T00:00:00Z. Any data with created_at before this data will not be synced. For Subscription Usage, this becomes the `timeframe_start` API parameter.
	StartDate string `json:"start_date"`
	// Property key names to extract from all events, in order to enrich ledger entries corresponding to an event deduction.
	StringEventPropertiesKeys []string `json:"string_event_properties_keys,omitempty"`
	// Property key name to group subscription usage by.
	SubscriptionUsageGroupingKey *string `json:"subscription_usage_grouping_key,omitempty"`
}

func (s SourceOrbUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOrbUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOrbUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceOrbUpdate) GetLookbackWindowDays() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindowDays
}

func (o *SourceOrbUpdate) GetNumericEventPropertiesKeys() []string {
	if o == nil {
		return nil
	}
	return o.NumericEventPropertiesKeys
}

func (o *SourceOrbUpdate) GetPlanID() *string {
	if o == nil {
		return nil
	}
	return o.PlanID
}

func (o *SourceOrbUpdate) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

func (o *SourceOrbUpdate) GetStringEventPropertiesKeys() []string {
	if o == nil {
		return nil
	}
	return o.StringEventPropertiesKeys
}

func (o *SourceOrbUpdate) GetSubscriptionUsageGroupingKey() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionUsageGroupingKey
}
