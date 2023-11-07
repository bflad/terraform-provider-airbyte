// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSyncModeEnum string

const (
	ConnectionSyncModeEnumFullRefreshOverwrite      ConnectionSyncModeEnum = "full_refresh_overwrite"
	ConnectionSyncModeEnumFullRefreshAppend         ConnectionSyncModeEnum = "full_refresh_append"
	ConnectionSyncModeEnumIncrementalAppend         ConnectionSyncModeEnum = "incremental_append"
	ConnectionSyncModeEnumIncrementalDedupedHistory ConnectionSyncModeEnum = "incremental_deduped_history"
)

func (e ConnectionSyncModeEnum) ToPointer() *ConnectionSyncModeEnum {
	return &e
}

func (e *ConnectionSyncModeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "full_refresh_overwrite":
		fallthrough
	case "full_refresh_append":
		fallthrough
	case "incremental_append":
		fallthrough
	case "incremental_deduped_history":
		*e = ConnectionSyncModeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSyncModeEnum: %v", v)
	}
}
