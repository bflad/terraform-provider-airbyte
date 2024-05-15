resource "airbyte_source_retently" "my_source_retently" {
  configuration = {
    credentials = {
      authenticate_via_retently_o_auth = {
        additional_properties = "{ \"see\": \"documentation\" }"
        client_id             = "...my_client_id..."
        client_secret         = "...my_client_secret..."
        refresh_token         = "...my_refresh_token..."
      }
    }
  }
  definition_id = "f5d1d34f-0cce-4548-aa3a-161dc53f6414"
  name          = "Ernesto Schmitt"
  secret_id     = "...my_secret_id..."
  workspace_id  = "b41d5bf9-4a01-4397-93df-d90aff660497"
}