// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceConfigcatRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceConfigcatResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceConfigcatGetResponse *shared.SourceConfigcatGetResponse
	StatusCode                 int
	RawResponse                *http.Response
}
