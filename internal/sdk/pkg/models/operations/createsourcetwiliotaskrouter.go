// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceTwilioTaskrouterResponse struct {
	ContentType string
	// Successful operation
	SourceTwilioTaskrouterGetResponse *shared.SourceTwilioTaskrouterGetResponse
	StatusCode                        int
	RawResponse                       *http.Response
}
