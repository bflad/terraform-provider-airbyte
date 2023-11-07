// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceLinkedinAdsPutRequest struct {
	Configuration SourceLinkedinAdsUpdate `json:"configuration"`
	Name          string                  `json:"name"`
	WorkspaceID   string                  `json:"workspaceId"`
}

func (o *SourceLinkedinAdsPutRequest) GetConfiguration() SourceLinkedinAdsUpdate {
	if o == nil {
		return SourceLinkedinAdsUpdate{}
	}
	return o.Configuration
}

func (o *SourceLinkedinAdsPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceLinkedinAdsPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
