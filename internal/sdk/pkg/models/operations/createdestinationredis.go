// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationRedisResponse struct {
	ContentType string
	// Successful operation
	DestinationRedisGetResponse *shared.DestinationRedisGetResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
