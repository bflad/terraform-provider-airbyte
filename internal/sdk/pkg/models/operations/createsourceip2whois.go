// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceIp2whoisResponse struct {
	ContentType string
	// Successful operation
	SourceIp2whoisGetResponse *shared.SourceIp2whoisGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
