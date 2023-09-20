// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePaypalTransactionGetResponse struct {
	Configuration SourcePaypalTransaction `json:"configuration"`
	Name          string                  `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	SourceID    *string `json:"sourceId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}
