// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceBigqueryRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceBigqueryResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceBigqueryGetResponse *shared.SourceBigqueryGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
