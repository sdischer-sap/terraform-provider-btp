package provider

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Needed for JSON mapping - fails with data types of resource
type subaccountRoleCollectionRoleRefTestType struct {
	Name              string `json:"name"`
	RoleTemplateAppId string `json:"role_template_app_id"`
	RoleTemplateName  string `json:"role_template_name"`
}

func TestResourceSubAccountRoleCollection(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_role_collection")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						"Description of my new role collection",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Viewer",
							RoleTemplateAppId: "cis-local!b2",
							RoleTemplateName:  "Subaccount_Viewer",
						},
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Destination Viewer",
							RoleTemplateAppId: "destination-xsappname!b9",
							RoleTemplateName:  "Destination_Viewer",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "2"),
					),
				},
				{
					ResourceName:      "btp_subaccount_role_collection.uut",
					ImportStateId:     "ef23ace8-6ade-4d78-9c1f-8df729548bbf,My new role collection",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("happy path - update", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_role_collection.update")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						"Description of my new role collection",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Viewer",
							RoleTemplateAppId: "cis-local!b2",
							RoleTemplateName:  "Subaccount_Viewer",
						},
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Destination Viewer",
							RoleTemplateAppId: "destination-xsappname!b9",
							RoleTemplateName:  "Destination_Viewer",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "2"),
					),
				},
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						"Updated description of my new role collection",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Service Auditor",
							RoleTemplateAppId: "service-manager!b3",
							RoleTemplateName:  "Subaccount_Service_Auditor",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Updated description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "1"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.0.name", "Subaccount Service Auditor"),
					),
				},
				{
					ResourceName:      "btp_subaccount_role_collection.uut",
					ImportStateId:     "ef23ace8-6ade-4d78-9c1f-8df729548bbf,My new role collection",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("happy path - update removing description", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_role_collection.update_rm_desc")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						"Description of my new role collection",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Viewer",
							RoleTemplateAppId: "cis-local!b2",
							RoleTemplateName:  "Subaccount_Viewer",
						},
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Destination Viewer",
							RoleTemplateAppId: "destination-xsappname!b9",
							RoleTemplateName:  "Destination_Viewer",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "2"),
					),
				},
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						"",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Service Auditor",
							RoleTemplateAppId: "service-manager!b3",
							RoleTemplateName:  "Subaccount_Service_Auditor",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", ""),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "1"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.0.name", "Subaccount Service Auditor"),
					),
				},
				{
					ResourceName:      "btp_subaccount_role_collection.uut",
					ImportStateId:     "ef23ace8-6ade-4d78-9c1f-8df729548bbf,My new role collection",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("happy path - update omitting description", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_role_collection.update_wo_desc")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						"Description of my new role collection",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Viewer",
							RoleTemplateAppId: "cis-local!b2",
							RoleTemplateName:  "Subaccount_Viewer",
						},
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Destination Viewer",
							RoleTemplateAppId: "destination-xsappname!b9",
							RoleTemplateName:  "Destination_Viewer",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "2"),
					),
				},
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollectionWoDescription(
						"uut",
						"ef23ace8-6ade-4d78-9c1f-8df729548bbf",
						"My new role collection",
						subaccountRoleCollectionRoleRefTestType{
							Name:              "Subaccount Service Auditor",
							RoleTemplateAppId: "service-manager!b3",
							RoleTemplateName:  "Subaccount_Service_Auditor",
						}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "1"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.0.name", "Subaccount Service Auditor"),
					),
				},
				{
					ResourceName:      "btp_subaccount_role_collection.uut",
					ImportStateId:     "ef23ace8-6ade-4d78-9c1f-8df729548bbf,My new role collection",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("error path - import with wrong key", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_role_collection.import_error")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubAccountRoleCollection("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", "My new role collection", "Description of my new role collection"),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "name", "My new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "description", "Description of my new role collection"),
						resource.TestCheckResourceAttr("btp_subaccount_role_collection.uut", "roles.#", "0"),
					),
				},
				{
					ResourceName:      "btp_subaccount_role_collection.uut",
					ImportStateId:     "ef23ace8-6ade-4d78-9c1f-8df729548bbf",
					ImportState:       true,
					ImportStateVerify: true,
					ExpectError:       regexp.MustCompile(`Expected import identifier with format: subaccount_id, name. Got:`),
				},
			},
		})
	})

	t.Run("error path - subaccount_id not a valid UUID", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(nil),
			Steps: []resource.TestStep{
				{
					Config:      hclResourceSubAccountRoleCollection("uut", "this-is-not-a-uuid", "My new role collection", "Description of my new role collection"),
					ExpectError: regexp.MustCompile(`Attribute subaccount_id value must be a valid UUID, got: this-is-not-a-uuid`),
				},
			},
		})
	})

	t.Run("error path - subacount_id mandatory", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(nil),
			Steps: []resource.TestStep{
				{
					Config:      hclResourceSubAccountRoleCollectionNoSubaccountId("uut", "My new role collection", "Description of my new role collection"),
					ExpectError: regexp.MustCompile(`The argument "subaccount_id" is required, but no definition was found`),
				},
			},
		})
	})

}

func hclResourceSubAccountRoleCollection(resourceName string, subaccountId string, displayName string, description string, roles ...subaccountRoleCollectionRoleRefTestType) string {
	if roles == nil {
		roles = []subaccountRoleCollectionRoleRefTestType{}
	}
	rolesJson, _ := json.Marshal(roles)

	return fmt.Sprintf(`resource "btp_subaccount_role_collection" "%s" {
        subaccount_id       = "%s"
		name      			= "%s"
        description      	= "%s"
		roles               = %v
    }`, resourceName, subaccountId, displayName, description, string(rolesJson))
}

func hclResourceSubAccountRoleCollectionNoSubaccountId(resourceName string, displayName string, description string) string {

	roles := []subaccountRoleCollectionRoleRefTestType{}

	roles = append(roles, subaccountRoleCollectionRoleRefTestType{
		Name:              "Subaccount Viewer",
		RoleTemplateAppId: "cis-local!b2",
		RoleTemplateName:  "Subaccount_Viewer",
	},
		subaccountRoleCollectionRoleRefTestType{
			Name:              "Destination Viewer",
			RoleTemplateAppId: "destination-xsappname!b9",
			RoleTemplateName:  "Destination_Viewer",
		},
	)

	rolesJson, _ := json.Marshal(roles)

	return fmt.Sprintf(`resource "btp_subaccount_role_collection" "%s" {
        name      			= "%s"
        description      	= "%s"
		roles               = %v
    }`, resourceName, displayName, description, string(rolesJson))
}

func hclResourceSubAccountRoleCollectionWoDescription(resourceName string, subaccountId string, displayName string, roles ...subaccountRoleCollectionRoleRefTestType) string {
	if roles == nil {
		roles = []subaccountRoleCollectionRoleRefTestType{}
	}
	rolesJson, _ := json.Marshal(roles)

	return fmt.Sprintf(`resource "btp_subaccount_role_collection" "%s" {
        subaccount_id       = "%s"
		name      			= "%s"
		roles               = %v
    }`, resourceName, subaccountId, displayName, string(rolesJson))
}
