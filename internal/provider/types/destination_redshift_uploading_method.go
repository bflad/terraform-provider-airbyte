// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type DestinationRedshiftUploadingMethod struct {
	AWSS3Staging *AWSS3Staging `tfsdk:"awss3_staging" tfPlanOnly:"true"`
	Standard     *Fake         `tfsdk:"standard" tfPlanOnly:"true"`
}
