// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTypeform struct {
	Credentials SourceTypeformAuthorizationMethod `tfsdk:"credentials"`
	FormIds     []types.String                    `tfsdk:"form_ids"`
	SourceType  types.String                      `tfsdk:"source_type"`
	StartDate   types.String                      `tfsdk:"start_date"`
}
