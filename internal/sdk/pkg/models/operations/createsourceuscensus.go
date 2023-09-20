// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceUsCensusResponse struct {
	ContentType string
	// Successful operation
	SourceUsCensusGetResponse *shared.SourceUsCensusGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
