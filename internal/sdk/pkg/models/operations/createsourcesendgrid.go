// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceSendgridResponse struct {
	ContentType string
	// Successful operation
	SourceSendgridGetResponse *shared.SourceSendgridGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
