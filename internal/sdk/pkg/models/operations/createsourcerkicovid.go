// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceRkiCovidResponse struct {
	ContentType string
	// Successful operation
	SourceRkiCovidGetResponse *shared.SourceRkiCovidGetResponse
	StatusCode                int
	RawResponse               *http.Response
}
