resource "tembo_instance" "test_instance" {
  instance_name = "test-instance"
  org_id        = "org_test" # Replace this with your Tembo organization id
  cpu           = "1"
  stack_type    = "Standard"
  environment   = "dev"
  memory        = "4Gi"
  storage       = "10Gi"
  replicas      = 1
  #ip_allow_list = ["71.190.46.60"]
  #extra_domains_rw = ["sample-invalid-domain.test.tembo-development.com"]
  #postgres_configs = [
  #  {
  #    name  = "max_connections"
  #    value = "200"
  #  },
  #  {
  #    name  = "wal_buffers"
  #    value = "10"
  #  }
  #]
  trunk_installs = [
    {
      name    = "pgmq"
      version = "0.24.0"
    }
  ]
  extensions = [{
    name        = "plperl"
    description = "PL/Perl procedural language"
    locations = [{
      database = "app"
      schema   = "public"
      version  = "1.0"
      enabled  = false
      },
      {
        database = "postgres"
        schema   = "public"
        version  = "1.0"
        enabled  = true
    }]
  }]
  connection_pooler = {
    enabled = true,
    pooler = {
      pool_mode = "transaction",
      parameters = {
        max_client_conn   = "20"
        default_pool_size = "100"
      }
    }
  }
}

output "instance" {
  value = tembo_instance.test_instance
}
