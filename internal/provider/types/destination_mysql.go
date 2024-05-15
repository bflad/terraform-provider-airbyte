// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMysql struct {
	Database          types.String                          `tfsdk:"database"`
	DisableTypeDedupe types.Bool                            `tfsdk:"disable_type_dedupe"`
	Host              types.String                          `tfsdk:"host"`
	JdbcURLParams     types.String                          `tfsdk:"jdbc_url_params"`
	Password          types.String                          `tfsdk:"password"`
	Port              types.Int64                           `tfsdk:"port"`
	RawDataSchema     types.String                          `tfsdk:"raw_data_schema"`
	TunnelMethod      *DestinationClickhouseSSHTunnelMethod `tfsdk:"tunnel_method"`
	Username          types.String                          `tfsdk:"username"`
}
