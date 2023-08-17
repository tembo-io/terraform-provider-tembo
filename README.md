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
      source = "registry.terraform.io/tembo-io/tembo"
    }
  }
}

provider "tembo" {}

data "tembo_example" "example" {}
```

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

Go to `internal/provider/temboclient` directory in your terminal.

Run following command to re-generate the go client code for the API. You might want to delete the contents of the directory first.

```bash
openapi-generator generate -i https://api.coredb.io/api-docs/openapi.json  -g go -o . --additional-properties=packageName=temboclient
```

## Releasing the Provider to Terraform Registry

The GitHub Action will trigger and create a release for your provider whenever a new valid version tag is pushed to the repository. Terraform provider versions must follow the [Semantic Versioning](https://semver.org/) standard (vMAJOR.MINOR.PATCH).

First, make the changes you want to the provider & push/merge changes to the main branch.

Next, create a new tag & push the changes to github.

```shell
git tag v0.2.1 && git push origin v0.2.1
```