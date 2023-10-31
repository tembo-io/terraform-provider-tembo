resource "tembo_instance" "test_db" {
  instance_name = "tfprovider-4"
  org_id        = "org_2UJ2WPYFsE42Cos6mlmIuwIIJ4V"
  cpu           = "1"
  stack_type    = "Standard"
  environment   = "dev"
  memory        = "4Gi"
  storage       = "10Gi"
  replicas      = 1
  #ip_allow_list = ["71.190.46.69"]
  # extra_domains_rw = ["sample-invalid-domain.test.tembo-development.com"]
  postgres_configs = [
    {
      name  = "max_connections"
      value = "200"
    },
    {
      name  = "wal_buffers"
      value = "16MB"
    }
  ]
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
    },
    {
      "name" : "pltclu",
      "description" : "PL/TclU untrusted procedural language",
      "locations" : [
        {
          "database" : "app",
          "schema" : "public",
          "version" : "1.0",
          "enabled" : false,
          "error" : false,
          "error_message" : null
        },
        {
          "database" : "postgres",
          "schema" : "public",
          "version" : "1.0",
          "enabled" : false,
          "error" : false,
          "error_message" : null
        }
      ]
  }]
  connection_pooler = {
    enabled = true,
    pooler = {
      pool_mode = "transaction",
      parameters = [
        {
          name  = "max_client_conn"
          value = "50"
        },
        {
          name  = "default_pool_size"
          value = "5000"
        }
      ]
    }
  }
}

data "tembo_instance_secrets" "test" {
  org_id      = "org_2UJ2WPYFsE42Cos6mlmIuwIIJ4V"
  instance_id = tembo_instance.test_db.instance_id
}

data "tembo_instance_secret" "test_sec" {
  org_id      = "org_2UJ2WPYFsE42Cos6mlmIuwIIJ4V"
  instance_id = tembo_instance.test_db.instance_id
  secret_name = "readonly-role"
}


output "instance" {
  value = tembo_instance.test_db
}

output "data" {
  value = data.tembo_instance_secrets.test
}

output "data_secret" {
  value = data.tembo_instance_secret.test_sec
}
