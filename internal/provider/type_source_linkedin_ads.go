// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLinkedinAds struct {
	AccountIds         []types.Int64                                     `tfsdk:"account_ids"`
	AdAnalyticsReports []SourceLinkedinAdsAdAnalyticsReportConfiguration `tfsdk:"ad_analytics_reports"`
	Credentials        *SourceLinkedinAdsAuthentication                  `tfsdk:"credentials"`
	SourceType         types.String                                      `tfsdk:"source_type"`
	StartDate          types.String                                      `tfsdk:"start_date"`
}
