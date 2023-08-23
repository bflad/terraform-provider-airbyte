// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceFakerResourceModel) ToCreateSDKType() *shared.SourceFakerCreateRequest {
	alwaysUpdated := new(bool)
	if !r.Configuration.AlwaysUpdated.IsUnknown() && !r.Configuration.AlwaysUpdated.IsNull() {
		*alwaysUpdated = r.Configuration.AlwaysUpdated.ValueBool()
	} else {
		alwaysUpdated = nil
	}
	count := r.Configuration.Count.ValueInt64()
	parallelism := new(int64)
	if !r.Configuration.Parallelism.IsUnknown() && !r.Configuration.Parallelism.IsNull() {
		*parallelism = r.Configuration.Parallelism.ValueInt64()
	} else {
		parallelism = nil
	}
	recordsPerSlice := new(int64)
	if !r.Configuration.RecordsPerSlice.IsUnknown() && !r.Configuration.RecordsPerSlice.IsNull() {
		*recordsPerSlice = r.Configuration.RecordsPerSlice.ValueInt64()
	} else {
		recordsPerSlice = nil
	}
	seed := new(int64)
	if !r.Configuration.Seed.IsUnknown() && !r.Configuration.Seed.IsNull() {
		*seed = r.Configuration.Seed.ValueInt64()
	} else {
		seed = nil
	}
	sourceType := shared.SourceFakerFaker(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceFaker{
		AlwaysUpdated:   alwaysUpdated,
		Count:           count,
		Parallelism:     parallelism,
		RecordsPerSlice: recordsPerSlice,
		Seed:            seed,
		SourceType:      sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFakerCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFakerResourceModel) ToGetSDKType() *shared.SourceFakerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFakerResourceModel) ToUpdateSDKType() *shared.SourceFakerPutRequest {
	alwaysUpdated := new(bool)
	if !r.Configuration.AlwaysUpdated.IsUnknown() && !r.Configuration.AlwaysUpdated.IsNull() {
		*alwaysUpdated = r.Configuration.AlwaysUpdated.ValueBool()
	} else {
		alwaysUpdated = nil
	}
	count := r.Configuration.Count.ValueInt64()
	parallelism := new(int64)
	if !r.Configuration.Parallelism.IsUnknown() && !r.Configuration.Parallelism.IsNull() {
		*parallelism = r.Configuration.Parallelism.ValueInt64()
	} else {
		parallelism = nil
	}
	recordsPerSlice := new(int64)
	if !r.Configuration.RecordsPerSlice.IsUnknown() && !r.Configuration.RecordsPerSlice.IsNull() {
		*recordsPerSlice = r.Configuration.RecordsPerSlice.ValueInt64()
	} else {
		recordsPerSlice = nil
	}
	seed := new(int64)
	if !r.Configuration.Seed.IsUnknown() && !r.Configuration.Seed.IsNull() {
		*seed = r.Configuration.Seed.ValueInt64()
	} else {
		seed = nil
	}
	configuration := shared.SourceFakerUpdate{
		AlwaysUpdated:   alwaysUpdated,
		Count:           count,
		Parallelism:     parallelism,
		RecordsPerSlice: recordsPerSlice,
		Seed:            seed,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFakerPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFakerResourceModel) ToDeleteSDKType() *shared.SourceFakerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFakerResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFakerResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
