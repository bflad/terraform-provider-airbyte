// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceDynamodbCredentials struct {
	AuthenticateViaAccessKeys *AuthenticateViaAccessKeys `tfsdk:"authenticate_via_access_keys" tfPlanOnly:"true"`
	RoleBasedAuthentication   *RoleBasedAuthentication   `tfsdk:"role_based_authentication" tfPlanOnly:"true"`
}
