// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceHubspotResponse struct {
	ContentType string
	// Successful operation
	SourceHubspotGetResponse *shared.SourceHubspotGetResponse
	StatusCode               int
	RawResponse              *http.Response
}
