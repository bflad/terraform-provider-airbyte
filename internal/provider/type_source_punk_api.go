// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePunkAPI struct {
	BrewedAfter  types.String `tfsdk:"brewed_after"`
	BrewedBefore types.String `tfsdk:"brewed_before"`
	ID           types.String `tfsdk:"id"`
}
