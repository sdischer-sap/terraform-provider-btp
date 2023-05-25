---
page_title: "btp_directory_role_collection Data Source - terraform-provider-btp"
subcategory: ""
description: |-
  Get details about a specific role collection.
---

# btp_directory_role_collection (Data Source)

Get details about a specific role collection.

## Example Usage

```terraform
data "btp_directory_role_collection" "directory_admin" {
  directory_id = "dd005d8b-1fee-4e6b-b6ff-cb9a197b7fe0"
  name         = "Directory Administrator"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `directory_id` (String) The ID of the directory.
- `name` (String) The name of the role collection.

### Read-Only

- `description` (String) The description of the role collection.
- `id` (String, Deprecated) The ID of the directory.
- `read_only` (Boolean) Whether the role collection is readonly.
- `role_references` (Attributes List) (see [below for nested schema](#nestedatt--role_references))

<a id="nestedatt--role_references"></a>
### Nested Schema for `role_references`

Read-Only:

- `description` (String) The description of the referenced role
- `name` (String) The name of the referenced role.
- `role_template_app_id` (String) The name of the referenced template app id
- `role_template_name` (String) The name of the referenced role template.