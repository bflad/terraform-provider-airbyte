// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceSurveymonkeyResponse struct {
	ContentType string
	// Successful operation
	SourceSurveymonkeyGetResponse *shared.SourceSurveymonkeyGetResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
