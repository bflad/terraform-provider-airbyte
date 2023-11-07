// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationAzureBlobStorageCreateRequest struct {
	Configuration DestinationAzureBlobStorage `json:"configuration"`
	Name          string                      `json:"name"`
	WorkspaceID   string                      `json:"workspaceId"`
}

func (o *DestinationAzureBlobStorageCreateRequest) GetConfiguration() DestinationAzureBlobStorage {
	if o == nil {
		return DestinationAzureBlobStorage{}
	}
	return o.Configuration
}

func (o *DestinationAzureBlobStorageCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationAzureBlobStorageCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
