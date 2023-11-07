// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourcePostmarkappRequest struct {
	SourcePostmarkappPutRequest *shared.SourcePostmarkappPutRequest `request:"mediaType=application/json"`
	SourceID                    string                              `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourcePostmarkappRequest) GetSourcePostmarkappPutRequest() *shared.SourcePostmarkappPutRequest {
	if o == nil {
		return nil
	}
	return o.SourcePostmarkappPutRequest
}

func (o *PutSourcePostmarkappRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourcePostmarkappResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourcePostmarkappResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourcePostmarkappResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourcePostmarkappResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
