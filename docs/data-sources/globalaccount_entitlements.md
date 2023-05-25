---
page_title: "btp_globalaccount_entitlements Data Source - terraform-provider-btp"
subcategory: ""
description: |-
  Get all the entitlements and quota assignments for a global account.
  To view all the resources a global account:
  * Target only the global account in the command line.
  * You must be assigned to either the global account admin or global account viewers role.
  Tips
  You must be assigned to one of these roles: global account admin, global account viewer.
---

# btp_globalaccount_entitlements (Data Source)

Get all the entitlements and quota assignments for a global account.

To view all the resources a global account:
* Target only the global account in the command line.
* You must be assigned to either the global account admin or global account viewers role.

__Tips__
You must be assigned to one of these roles: global account admin, global account viewer.

## Example Usage

```terraform
data "btp_globalaccount_entitlements" "all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `values` (Attributes Map) (see [below for nested schema](#nestedatt--values))

<a id="nestedatt--values"></a>
### Nested Schema for `values`

Read-Only:

- `plan_description` (String) The description of the entitled service plan.
- `plan_display_name` (String) The display name of the entitled service plan.
- `plan_name` (String) The name of the entitled service plan.
- `quota_assigned` (Number) The overall quota assigned.
- `quota_remaining` (Number) The quota which is not used.
- `service_display_name` (String) The display name of the entitled service.
- `service_name` (String) The name of the entitled service.