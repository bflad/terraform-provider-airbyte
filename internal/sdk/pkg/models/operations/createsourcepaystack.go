// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourcePaystackResponse struct {
	ContentType string
	// Successful operation
	SourcePaystackGetResponse *shared.SourcePaystackGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
