// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationDynamodbGetResponse struct {
	Configuration DestinationDynamodb `json:"configuration"`
	DestinationID *string             `json:"destinationId,omitempty"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}
