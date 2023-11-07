// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationXataRequest struct {
	DestinationXataPutRequest *shared.DestinationXataPutRequest `request:"mediaType=application/json"`
	DestinationID             string                            `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationXataRequest) GetDestinationXataPutRequest() *shared.DestinationXataPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationXataPutRequest
}

func (o *PutDestinationXataRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationXataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationXataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationXataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationXataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
