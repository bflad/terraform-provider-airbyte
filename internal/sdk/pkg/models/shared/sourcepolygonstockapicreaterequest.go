// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePolygonStockAPICreateRequest struct {
	Configuration SourcePolygonStockAPI `json:"configuration"`
	Name          string                `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourcePolygonStockAPICreateRequest) GetConfiguration() SourcePolygonStockAPI {
	if o == nil {
		return SourcePolygonStockAPI{}
	}
	return o.Configuration
}

func (o *SourcePolygonStockAPICreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePolygonStockAPICreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourcePolygonStockAPICreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
