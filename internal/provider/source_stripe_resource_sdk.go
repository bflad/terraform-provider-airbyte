// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceStripeResourceModel) ToCreateSDKType() *shared.SourceStripeCreateRequest {
	accountID := r.Configuration.AccountID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	sliceRange := new(int64)
	if !r.Configuration.SliceRange.IsUnknown() && !r.Configuration.SliceRange.IsNull() {
		*sliceRange = r.Configuration.SliceRange.ValueInt64()
	} else {
		sliceRange = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceStripe{
		AccountID:          accountID,
		ClientSecret:       clientSecret,
		LookbackWindowDays: lookbackWindowDays,
		SliceRange:         sliceRange,
		StartDate:          startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceStripeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceStripeResourceModel) ToGetSDKType() *shared.SourceStripeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceStripeResourceModel) ToUpdateSDKType() *shared.SourceStripePutRequest {
	accountID := r.Configuration.AccountID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	sliceRange := new(int64)
	if !r.Configuration.SliceRange.IsUnknown() && !r.Configuration.SliceRange.IsNull() {
		*sliceRange = r.Configuration.SliceRange.ValueInt64()
	} else {
		sliceRange = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceStripeUpdate{
		AccountID:          accountID,
		ClientSecret:       clientSecret,
		LookbackWindowDays: lookbackWindowDays,
		SliceRange:         sliceRange,
		StartDate:          startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceStripePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceStripeResourceModel) ToDeleteSDKType() *shared.SourceStripeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceStripeResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceStripeResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
