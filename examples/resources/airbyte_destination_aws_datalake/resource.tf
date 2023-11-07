resource "airbyte_destination_aws_datalake" "my_destination_awsdatalake" {
  configuration = {
    aws_account_id = "111111111111"
    bucket_name    = "...my_bucket_name..."
    bucket_prefix  = "...my_bucket_prefix..."
    credentials = {
      iam_role = {
        role_arn = "...my_role_arn..."
      }
    }
    format = {
      json_lines_newline_delimited_json = {
        compression_codec = "GZIP"
        format_type       = "JSONL"
      }
    }
    glue_catalog_float_as_decimal             = false
    lakeformation_database_default_tag_key    = "pii_level"
    lakeformation_database_default_tag_values = "private,public"
    lakeformation_database_name               = "...my_lakeformation_database_name..."
    lakeformation_governed_tables             = false
    partitioning                              = "YEAR/MONTH/DAY"
    region                                    = "eu-west-1"
  }
  name         = "Rosa Harris"
  workspace_id = "0f2a9b0d-e149-47ac-8629-f5a79ed4a852"
}