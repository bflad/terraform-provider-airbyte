// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceFreshsalesResponse struct {
	ContentType string
	// Successful operation
	SourceFreshsalesGetResponse *shared.SourceFreshsalesGetResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
