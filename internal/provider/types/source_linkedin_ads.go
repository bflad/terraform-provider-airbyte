// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLinkedinAds struct {
	AccountIds         []types.Int64                    `tfsdk:"account_ids"`
	AdAnalyticsReports []AdAnalyticsReportConfiguration `tfsdk:"ad_analytics_reports"`
	Credentials        *SourceLinkedinAdsAuthentication `tfsdk:"credentials"`
	LookbackWindow     types.Int64                      `tfsdk:"lookback_window"`
	StartDate          types.String                     `tfsdk:"start_date"`
}
