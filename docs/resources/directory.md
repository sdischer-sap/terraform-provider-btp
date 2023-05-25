---
page_title: "btp_directory Resource - terraform-provider-btp"
subcategory: ""
description: |-
  Directories allow you to organize and manage your subaccounts according to your technical and business needs. The use of directories is optional.
  You can create up to 5 levels of directories in your account hierarchy. If you have directories, you can still create subaccounts directly under your global account.
  Directory features: Set the '--features' parameter to specify which features to enable for the directory. Use either the feature name or its character.
  * DEFAULT (D): (Required) All directories provide the following basic features: (i) Group and filter subaccounts for reports and filters, (ii) monitor usage and costs on a directory level (costs only available for contracts that use the consumption-based commercial model), and (iii) assign labels to the directory for identification and reporting purposes.
  * ENTITLEMENTS (E): (Optional) Enables the assignment of a quota for services and applications to the directory from the global account quota for distribution to the directory's subaccounts and subdirectories.
  * AUTHORIZATIONS (A): (Optional) Allows you to assign users as administrators or viewers of this directory. For example, allow certain users to manage the directory's entitlements. You can enable this feature only in combination with the ENTITLEMENTS (E) feature.
  Tips
  * You must be assigned to the global account admin role, or the directory admin if the directory is configured to manage its authorizations.
  * A directory path in the account hierarchy can have only one directory that is enabled with the ENTITLEMENTS (E) or AUTHORIZATIONS (A) features. If such a directory exists, other directories in that path can only be enabled with the DEFAULT (D) features.
  Further documentation
  https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/8ed4a705efa0431b910056c0acdbf377.html
---

# btp_directory (Resource)

Directories allow you to organize and manage your subaccounts according to your technical and business needs. The use of directories is optional.

You can create up to 5 levels of directories in your account hierarchy. If you have directories, you can still create subaccounts directly under your global account.

Directory features: Set the '--features' parameter to specify which features to enable for the directory. Use either the feature name or its character.
* DEFAULT (D): (Required) All directories provide the following basic features: (i) Group and filter subaccounts for reports and filters, (ii) monitor usage and costs on a directory level (costs only available for contracts that use the consumption-based commercial model), and (iii) assign labels to the directory for identification and reporting purposes.
* ENTITLEMENTS (E): (Optional) Enables the assignment of a quota for services and applications to the directory from the global account quota for distribution to the directory's subaccounts and subdirectories.
* AUTHORIZATIONS (A): (Optional) Allows you to assign users as administrators or viewers of this directory. For example, allow certain users to manage the directory's entitlements. You can enable this feature only in combination with the ENTITLEMENTS (E) feature.

__Tips__
* You must be assigned to the global account admin role, or the directory admin if the directory is configured to manage its authorizations.
* A directory path in the account hierarchy can have only one directory that is enabled with the ENTITLEMENTS (E) or AUTHORIZATIONS (A) features. If such a directory exists, other directories in that path can only be enabled with the DEFAULT (D) features.

__Further documentation__
https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/8ed4a705efa0431b910056c0acdbf377.html

## Example Usage

```terraform
resource "btp_directory" "parent" {
  name        = "my-parent-directory"
  description = "This is a parent directory."
}

resource "btp_directory" "child" {
  parent_id = btp_directory.parent.id
  name        = "my-child-directory"
  description = "This is a child directory."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The display name of the directory.

### Optional

- `description` (String) A description of the directory.
- `parent_id` (String) The GUID of the directory's parent entity. Typically this is the global account.
- `subdomain` (String) Applies only to directories that have the user authorization management feature enabled. The subdomain becomes part of the path used to access the authorization tenant of the directory. Unique within the defined region.

### Read-Only

- `created_by` (String) Details of the user that created the directory.
- `created_date` (String) The date and time the resource was created in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format.
- `features` (Set of String) The features that are enabled for the directory. Valid values: - DEFAULT: (Mandatory) All directories have the following basic feature enabled. (1) Group and filter subaccounts for reports and filters, (2) monitor usage and costs on a directory level (costs only available for contracts that use the consumption-based commercial model), and (3) set custom properties and tags to the directory for identification and reporting purposes. - ENTITLEMENTS: (Optional) Allows the assignment of a quota for services and applications to the directory from the global account quota for distribution to the subaccounts under this directory.  - AUTHORIZATIONS: (Optional) Allows the assignment of users as administrators or viewers of this directory. You must apply this feature in combination with the ENTITLEMENTS feature. <br/><b>Valid values:</b>  [DEFAULT] [DEFAULT,ENTITLEMENTS] [DEFAULT,ENTITLEMENTS,AUTHORIZATIONS]<br/>
- `id` (String) The ID of the directory.
- `labels` (Map of Set of String) Contains information about the labels assigned to a specified global account. Labels are represented in a JSON array of key-value pairs; each key has up to 10 corresponding values. This field replaces the deprecated "customProperties" field, which supports only single values per key.
- `last_modified` (String) The date and time the resource was last modified in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format.
- `state` (String) The current state of the directory. * <b>STARTED:</b> CRUD operation on an entity has started. * <b>CREATING:</b> Creating entity operation is in progress. * <b>UPDATING:</b> Updating entity operation is in progress. * <b>MOVING:</b> Moving entity operation is in progress. * <b>PROCESSING:</b> A series of operations related to the entity is in progress. * <b>DELETING:</b> Deleting entity operation is in progress. * <b>OK:</b> The CRUD operation or series of operations completed successfully. * <b>PENDING REVIEW:</b> The processing operation has been stopped for reviewing and can be restarted by the operator. * <b>CANCELLED:</b> The operation or processing was canceled by the operator. * <b>CREATION_FAILED:</b> The creation operation failed, and the entity was not created or was created but cannot be used. * <b>UPDATE_FAILED:</b> The update operation failed, and the entity was not updated. * <b>PROCESSING_FAILED:</b> The processing operations failed. * <b>DELETION_FAILED:</b> The delete operation failed, and the entity was not deleted. * <b>MOVE_FAILED:</b> Entity could not be moved to a different location. * <b>MIGRATING:</b> Migrating entity from NEO to CF.

## Import

Import is supported using the following syntax:

```terraform
# terraform import btp_directory.<resource_name> <directory_id>

terraform import btp_directory.parent dd005d8b-1fee-4e6b-b6ff-cb9a197b7fe0
```