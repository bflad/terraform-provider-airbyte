resource "airbyte_destination_mongodb" "my_destination_mongodb" {
  configuration = {
    auth_type = {
      login_password = {
        password = "...my_password..."
        username = "Jarrell.Schuppe"
      }
    }
    database = "...my_database..."
    instance_type = {
      mongo_db_atlas = {
        cluster_url = "...my_cluster_url..."
        instance    = "atlas"
      }
    }
    tunnel_method = {
      destination_mongodb_no_tunnel = {}
    }
  }
  name         = "Eloise Stark"
  workspace_id = "5e758afd-2ad3-48ed-bed0-e5e2212dc7d0"
}