---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tembo_instance_secrets Data Source - terraform-provider-tembo"
subcategory: ""
description: |-
  Data Source for Tembo Instance Secrets.
---

# tembo_instance_secrets (Data Source)

Data Source for Tembo Instance Secrets.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_id` (String) Unique ID for the instance generated by Tembo
- `org_id` (String) Id of the organization in which the instance will be created

### Read-Only

- `available_secrets` (Attributes List) (see [below for nested schema](#nestedatt--available_secrets))

<a id="nestedatt--available_secrets"></a>
### Nested Schema for `available_secrets`

Optional:

- `name` (String) Name of the secret
- `possible_keys` (List of String) Possible Keys for the secret
