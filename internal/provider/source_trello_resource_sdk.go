// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTrelloResourceModel) ToCreateSDKType() *shared.SourceTrelloCreateRequest {
	var boardIds []string = nil
	for _, boardIdsItem := range r.Configuration.BoardIds {
		boardIds = append(boardIds, boardIdsItem.ValueString())
	}
	key := r.Configuration.Key.ValueString()
	sourceType := shared.SourceTrelloTrello(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceTrello{
		BoardIds:   boardIds,
		Key:        key,
		SourceType: sourceType,
		StartDate:  startDate,
		Token:      token,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTrelloCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrelloResourceModel) ToGetSDKType() *shared.SourceTrelloCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrelloResourceModel) ToUpdateSDKType() *shared.SourceTrelloPutRequest {
	var boardIds []string = nil
	for _, boardIdsItem := range r.Configuration.BoardIds {
		boardIds = append(boardIds, boardIdsItem.ValueString())
	}
	key := r.Configuration.Key.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceTrelloUpdate{
		BoardIds:  boardIds,
		Key:       key,
		StartDate: startDate,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTrelloPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrelloResourceModel) ToDeleteSDKType() *shared.SourceTrelloCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrelloResourceModel) RefreshFromGetResponse(resp *shared.SourceTrelloGetResponse) {
	r.Configuration.BoardIds = nil
	for _, v := range resp.Configuration.BoardIds {
		r.Configuration.BoardIds = append(r.Configuration.BoardIds, types.StringValue(v))
	}
	r.Configuration.Key = types.StringValue(resp.Configuration.Key)
	r.Configuration.SourceType = types.StringValue(string(resp.Configuration.SourceType))
	r.Configuration.StartDate = types.StringValue(resp.Configuration.StartDate.Format(time.RFC3339))
	r.Configuration.Token = types.StringValue(resp.Configuration.Token)
	r.Name = types.StringValue(resp.Name)
	if resp.SecretID != nil {
		r.SecretID = types.StringValue(*resp.SecretID)
	} else {
		r.SecretID = types.StringNull()
	}
	if resp.SourceID != nil {
		r.SourceID = types.StringValue(*resp.SourceID)
	} else {
		r.SourceID = types.StringNull()
	}
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTrelloResourceModel) RefreshFromCreateResponse(resp *shared.SourceTrelloGetResponse) {
	r.RefreshFromGetResponse(resp)
}
