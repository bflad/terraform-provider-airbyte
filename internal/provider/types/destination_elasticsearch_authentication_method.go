// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type DestinationElasticsearchAuthenticationMethod struct {
	APIKeySecret     *APIKeySecret     `tfsdk:"api_key_secret" tfPlanOnly:"true"`
	None             *Fake             `tfsdk:"none" tfPlanOnly:"true"`
	UsernamePassword *UsernamePassword `tfsdk:"username_password" tfPlanOnly:"true"`
}
