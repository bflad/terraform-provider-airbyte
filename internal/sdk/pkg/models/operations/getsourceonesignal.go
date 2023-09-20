// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceOnesignalRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceOnesignalResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceOnesignalGetResponse *shared.SourceOnesignalGetResponse
	StatusCode                 int
	RawResponse                *http.Response
}
