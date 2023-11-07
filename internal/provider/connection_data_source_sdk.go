// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ConnectionDataSourceModel) RefreshFromGetResponse(resp *shared.ConnectionResponse) {
	r.Configurations.Streams = nil
	for _, streamsItem := range resp.Configurations.Streams {
		var streams1 StreamConfiguration
		streams1.CursorField = nil
		for _, v := range streamsItem.CursorField {
			streams1.CursorField = append(streams1.CursorField, types.StringValue(v))
		}
		streams1.Name = types.StringValue(streamsItem.Name)
		streams1.PrimaryKey = nil
		for _, primaryKeyItem := range streamsItem.PrimaryKey {
			var primaryKey1 []types.String
			primaryKey1 = nil
			for _, v := range primaryKeyItem {
				primaryKey1 = append(primaryKey1, types.StringValue(v))
			}
			streams1.PrimaryKey = append(streams1.PrimaryKey, primaryKey1)
		}
		if streamsItem.SyncMode != nil {
			streams1.SyncMode = types.StringValue(string(*streamsItem.SyncMode))
		} else {
			streams1.SyncMode = types.StringNull()
		}
		r.Configurations.Streams = append(r.Configurations.Streams, streams1)
	}
	r.ConnectionID = types.StringValue(resp.ConnectionID)
	if resp.DataResidency != nil {
		r.DataResidency = types.StringValue(string(*resp.DataResidency))
	} else {
		r.DataResidency = types.StringNull()
	}
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.Name = types.StringValue(resp.Name)
	if resp.NamespaceDefinition != nil {
		r.NamespaceDefinition = types.StringValue(string(*resp.NamespaceDefinition))
	} else {
		r.NamespaceDefinition = types.StringNull()
	}
	if resp.NamespaceFormat != nil {
		r.NamespaceFormat = types.StringValue(*resp.NamespaceFormat)
	} else {
		r.NamespaceFormat = types.StringNull()
	}
	if resp.NonBreakingSchemaUpdatesBehavior != nil {
		r.NonBreakingSchemaUpdatesBehavior = types.StringValue(string(*resp.NonBreakingSchemaUpdatesBehavior))
	} else {
		r.NonBreakingSchemaUpdatesBehavior = types.StringNull()
	}
	if resp.Prefix != nil {
		r.Prefix = types.StringValue(*resp.Prefix)
	} else {
		r.Prefix = types.StringNull()
	}
	if resp.Schedule.BasicTiming != nil {
		r.Schedule.BasicTiming = types.StringValue(*resp.Schedule.BasicTiming)
	} else {
		r.Schedule.BasicTiming = types.StringNull()
	}
	if resp.Schedule.CronExpression != nil {
		r.Schedule.CronExpression = types.StringValue(*resp.Schedule.CronExpression)
	} else {
		r.Schedule.CronExpression = types.StringNull()
	}
	r.Schedule.ScheduleType = types.StringValue(string(resp.Schedule.ScheduleType))
	r.SourceID = types.StringValue(resp.SourceID)
	r.Status = types.StringValue(string(resp.Status))
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
