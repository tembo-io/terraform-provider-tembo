---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tembo_instance Resource - terraform-provider-tembo"
subcategory: ""
description: |-
  Resource for creating a tembo instance.
---

# tembo_instance (Resource)

Resource for creating a tembo instance.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cpu` (String) CPU
- `environment` (String) Environment
- `instance_name` (String) Unique name of the instance
- `memory` (String) Memory
- `org_id` (String) Id of the organization in which the instance will be created
- `stack_type` (String) Stack type for the instance.
- `storage` (String) Storage

### Optional

- `autoscaling` (Attributes) (see [below for nested schema](#nestedatt--autoscaling))
- `connection_pooler` (Attributes) (see [below for nested schema](#nestedatt--connection_pooler))
- `extensions` (Attributes List) Extensions to install in the instance (see [below for nested schema](#nestedatt--extensions))
- `extra_domains_rw` (List of String) Custom domain. Read more [here](https://tembo.io/docs/tembo-cloud/custom-domains/)
- `first_recoverability_time` (String) The time at which the instance first became recoverable
- `ip_allow_list` (List of String) Allowed IP list
- `pg_version` (String) PostgresSQL version to deploy to Tembo Cloud
- `postgres_configs` (Attributes List) Postgres configs (see [below for nested schema](#nestedatt--postgres_configs))
- `provider_id` (String) The Cloud Provider ID where the instance will be created
- `region_id` (String) The cloud provider region where the instance will be created
- `replicas` (Number) Instance replicas
- `restore` (Attributes) (see [below for nested schema](#nestedatt--restore))
- `trunk_installs` (Attributes List) Trunk installs (see [below for nested schema](#nestedatt--trunk_installs))

### Read-Only

- `instance_id` (String) Unique ID for the instance generated by Tembo
- `last_updated` (String) Last updated date time in UTC
- `state` (String) Instance state. Values: Submitted, Up, Configuring, Error, Restarting, Starting, Stopping, Stopped, Deleting, Deleted

<a id="nestedatt--autoscaling"></a>
### Nested Schema for `autoscaling`

Optional:

- `autostop` (Attributes) (see [below for nested schema](#nestedatt--autoscaling--autostop))
- `storage` (Attributes) (see [below for nested schema](#nestedatt--autoscaling--storage))

<a id="nestedatt--autoscaling--autostop"></a>
### Nested Schema for `autoscaling.autostop`

Optional:

- `enabled` (Boolean)


<a id="nestedatt--autoscaling--storage"></a>
### Nested Schema for `autoscaling.storage`

Optional:

- `enabled` (Boolean)
- `maximum` (String)



<a id="nestedatt--connection_pooler"></a>
### Nested Schema for `connection_pooler`

Required:

- `enabled` (Boolean)

Optional:

- `pooler` (Attributes) (see [below for nested schema](#nestedatt--connection_pooler--pooler))

<a id="nestedatt--connection_pooler--pooler"></a>
### Nested Schema for `connection_pooler.pooler`

Required:

- `parameters` (Map of String) Parameters
- `pool_mode` (String)



<a id="nestedatt--extensions"></a>
### Nested Schema for `extensions`

Required:

- `locations` (Attributes List) Locations for extension (see [below for nested schema](#nestedatt--extensions--locations))
- `name` (String) Extension Name

Optional:

- `description` (String) Extension Description

<a id="nestedatt--extensions--locations"></a>
### Nested Schema for `extensions.locations`

Required:

- `database` (String) Database to install extension
- `enabled` (Boolean) Enabled

Optional:

- `schema` (String) The name of the schema
- `version` (String) The version of the extension to install



<a id="nestedatt--postgres_configs"></a>
### Nested Schema for `postgres_configs`

Required:

- `name` (String) Postgres config name

Optional:

- `value` (String) Postgres config value


<a id="nestedatt--restore"></a>
### Nested Schema for `restore`

Required:

- `instance_id` (String)

Optional:

- `recovery_target_time` (String)


<a id="nestedatt--trunk_installs"></a>
### Nested Schema for `trunk_installs`

Required:

- `name` (String) Trunk install name

Optional:

- `version` (String) Trunk install version
