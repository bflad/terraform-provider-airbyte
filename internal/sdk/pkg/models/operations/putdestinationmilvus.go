// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationMilvusRequest struct {
	DestinationMilvusPutRequest *shared.DestinationMilvusPutRequest `request:"mediaType=application/json"`
	DestinationID               string                              `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationMilvusRequest) GetDestinationMilvusPutRequest() *shared.DestinationMilvusPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationMilvusPutRequest
}

func (o *PutDestinationMilvusRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationMilvusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationMilvusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationMilvusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationMilvusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
