// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAppfigures struct {
	APIKey      types.String `tfsdk:"api_key"`
	GroupBy     types.String `tfsdk:"group_by"`
	SearchStore types.String `tfsdk:"search_store"`
	StartDate   types.String `tfsdk:"start_date"`
}
