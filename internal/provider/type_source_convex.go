// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceConvex struct {
	AccessKey     types.String `tfsdk:"access_key"`
	DeploymentURL types.String `tfsdk:"deployment_url"`
	SourceType    types.String `tfsdk:"source_type"`
}
