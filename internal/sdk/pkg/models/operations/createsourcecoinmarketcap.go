// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceCoinmarketcapResponse struct {
	ContentType string
	// Successful operation
	SourceCoinmarketcapGetResponse *shared.SourceCoinmarketcapGetResponse
	StatusCode                     int
	RawResponse                    *http.Response
}
