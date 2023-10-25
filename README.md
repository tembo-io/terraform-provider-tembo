# Tembo Terraform Provider

The Terraform provider for [Tembo](https://tembo.io/)

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Using the provider

```
terraform {
  required_providers {
    tembo = {
      source  = "tembo-io/tembo"
      version = ">= 0.1.0"
    }
  }
}

provider "tembo" {
}

resource "tembo_instance" "test_db" {
  instance_name = "tfprovider-de2345"
  org_id        = "org_test" # Replace this with your Tembo organization id
  cpu           = "1"
  stack_type    = "Standard"
  environment   = "dev"
  memory        = "4Gi"
  storage       = "10Gi"
  replicas      = 1
  # ip_allow_list = ["71.190.46.69"]
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
}

data "tembo_instance_secrets" "test" {
  org_id      = "org_test" # Replace this with your Tembo organization id
  instance_id = tembo_instance.test_db.instance_id
}

data "tembo_instance_secret" "test_sec" {
  org_id      = "org_test" # Replace this with your Tembo organization id
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
```

## Managing `trunk_installs` & `extensions`

Provider currently only checks whether there is an error or not with `trunk_installs` & `extensions` when creating or updating `tembo_instance`. If there is no error then it assumes the terraform apply is successful. It doesn’t compare the desired and actual state for them since some extensions get installed as part of the stack or base image.

Also, when importing `tembo_instance` it doesn’t import `trunk_installs` and `extensions` because of the same reason. After you import you can run terraform apply and it updates the state appropriately. It tries re-installing extensions but control plane ignores the request.

If in future there was a way to identify why extensions were installed (base image , stack , trunk_installs or extensions) then the provider can be updated to compare desired and actual state

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```

## Generating Go Client from API

[OpenAPI Generator](https://openapi-generator.tech/) tool is used to generate GO Client.

Install OpenAPI Generator if not already by following steps [here](https://openapi-generator.tech/docs/installation)

### Control plane API client

Go to `internal/provider/temboclient` directory in your terminal.

Delete the contents of the directory first and then run following command to re-generate the go client code for the API.

```bash
openapi-generator generate -i https://api.coredb.io/api-docs/openapi.json  -g go -o . --additional-properties=packageName=temboclient
```

### Data plane API client

Go to `internal/provider/tembodataclient` directory in your terminal.

Delete the contents of the directory first and then run following command to re-generate the go client code for the API.

```bash
openapi-generator generate -i https://api.data-1.use1.tembo.io/api-docs/openapi.json  -g go -o . --additional-properties=packageName=tembodataclient
```

## Releasing the Provider to Terraform Registry

The GitHub Action will trigger and create a release for your provider whenever a new valid version tag is pushed to the repository. Terraform provider versions must follow the [Semantic Versioning](https://semver.org/) standard (vMAJOR.MINOR.PATCH).

First, make the changes you want to the provider & push/merge changes to the main branch.

Next, create a new tag & push the changes to github.

```shell
git tag v0.2.1 && git push origin v0.2.1
```
