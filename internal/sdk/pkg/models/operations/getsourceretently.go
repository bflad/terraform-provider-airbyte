// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceRetentlyRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceRetentlyResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceRetentlyGetResponse *shared.SourceRetentlyGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
