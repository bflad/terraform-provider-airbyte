// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceLaunchdarklyResponse struct {
	ContentType string
	// Successful operation
	SourceLaunchdarklyGetResponse *shared.SourceLaunchdarklyGetResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
