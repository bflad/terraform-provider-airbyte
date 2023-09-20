// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceZohoCrmRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceZohoCrmResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceZohoCrmGetResponse *shared.SourceZohoCrmGetResponse
	StatusCode               int
	RawResponse              *http.Response
}
