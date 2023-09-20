// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceZendeskSupportRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceZendeskSupportResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceZendeskSupportGetResponse *shared.SourceZendeskSupportGetResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
