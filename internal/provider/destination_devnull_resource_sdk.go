// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDevNullResourceModel) ToCreateSDKType() *shared.DestinationDevNullCreateRequest {
	var testDestination shared.DestinationDevNullTestDestination
	var destinationDevNullSilent *shared.DestinationDevNullSilent
	if r.Configuration.TestDestination.Silent != nil {
		destinationDevNullSilent = &shared.DestinationDevNullSilent{}
	}
	if destinationDevNullSilent != nil {
		testDestination = shared.DestinationDevNullTestDestination{
			Silent: destinationDevNullSilent,
		}
	}
	configuration := shared.DestinationDevNull{
		TestDestination: testDestination,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDevNullCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDevNullResourceModel) ToGetSDKType() *shared.DestinationDevNullCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDevNullResourceModel) ToUpdateSDKType() *shared.DestinationDevNullPutRequest {
	var testDestination shared.TestDestination
	var silent *shared.Silent
	if r.Configuration.TestDestination.Silent != nil {
		silent = &shared.Silent{}
	}
	if silent != nil {
		testDestination = shared.TestDestination{
			Silent: silent,
		}
	}
	configuration := shared.DestinationDevNullUpdate{
		TestDestination: testDestination,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDevNullPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDevNullResourceModel) ToDeleteSDKType() *shared.DestinationDevNullCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDevNullResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDevNullResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
