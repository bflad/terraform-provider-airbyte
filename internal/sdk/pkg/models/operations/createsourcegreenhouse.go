// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceGreenhouseResponse struct {
	ContentType string
	// Successful operation
	SourceGreenhouseGetResponse *shared.SourceGreenhouseGetResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
