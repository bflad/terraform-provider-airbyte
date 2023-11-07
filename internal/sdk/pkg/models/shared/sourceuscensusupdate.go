// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceUsCensusUpdate struct {
	// Your API Key. Get your key <a href="https://api.census.gov/data/key_signup.html">here</a>.
	APIKey string `json:"api_key"`
	// The query parameters portion of the GET request, without the api key
	QueryParams *string `json:"query_params,omitempty"`
	// The path portion of the GET request
	QueryPath string `json:"query_path"`
}

func (o *SourceUsCensusUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceUsCensusUpdate) GetQueryParams() *string {
	if o == nil {
		return nil
	}
	return o.QueryParams
}

func (o *SourceUsCensusUpdate) GetQueryPath() string {
	if o == nil {
		return ""
	}
	return o.QueryPath
}
