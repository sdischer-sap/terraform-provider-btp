package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestResourceSubaccountTrustConfiguration(t *testing.T) {
	t.Parallel()

	t.Run("happy path - complete configuration with update", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_trust_configuration.complete")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubaccountTrustConfigurationComplete("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", "terraformint.accounts400.ondemand.com", "terraformint.accounts400.ondemand.com", "Custom IAS tenant for apps", "Description for terraformint.accounts400.ondemand.com", "custom link text", false, false, "inactive"),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestMatchResourceAttr("btp_subaccount_trust_configuration.uut", "subaccount_id", regexpValidUUID),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "identity_provider", "terraformint.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "domain", "terraformint.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "name", "Custom IAS tenant for apps"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "description", "Description for terraformint.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "link_text", "custom link text"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "available_for_user_logon", "false"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "auto_create_shadow_users", "false"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "origin", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "id", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "type", "Application"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "protocol", "OpenID Connect"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "status", "inactive"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "read_only", "false"),
					),
				},
				{
					Config: hclProviderFor(user) + hclResourceSubaccountTrustConfigurationMinimal("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", "terraformint.accounts400.ondemand.com"),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestMatchResourceAttr("btp_subaccount_trust_configuration.uut", "subaccount_id", regexpValidUUID),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "identity_provider", "terraformint.accounts400.ondemand.com"),
						resource.TestCheckNoResourceAttr("btp_subaccount_trust_configuration.uut", "domain"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "name", "Custom IAS tenant for apps"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "description", "Description for terraformint.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "link_text", "custom link text"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "available_for_user_logon", "true"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "auto_create_shadow_users", "true"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "origin", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "id", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "type", "Application"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "protocol", "OpenID Connect"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "status", "active"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "read_only", "false"),
					),
				},
				{
					ResourceName:      "btp_subaccount_trust_configuration.uut",
					ImportStateIdFunc: getTrustConfigIdForImport("btp_subaccount_trust_configuration.uut"),
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("happy path - minimal configuration with update", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_trust_configuration.minimal")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubaccountTrustConfigurationMinimal("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", "terraformint.accounts400.ondemand.com"),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestMatchResourceAttr("btp_subaccount_trust_configuration.uut", "subaccount_id", regexpValidUUID),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "identity_provider", "terraformint.accounts400.ondemand.com"),
						resource.TestCheckNoResourceAttr("btp_subaccount_trust_configuration.uut", "domain"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "name", "Custom IAS tenant"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "description", "IAS tenant terraformint.accounts400.ondemand.com (OpenID Connect)"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "link_text", "terraformint.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "available_for_user_logon", "true"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "auto_create_shadow_users", "true"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "origin", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "id", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "type", "Application"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "protocol", "OpenID Connect"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "status", "active"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "read_only", "false"),
					),
				},
				{
					Config: hclProviderFor(user) + hclResourceSubaccountTrustConfigurationComplete("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", "terraformtest.accounts400.ondemand.com", "terraformtest.accounts400.ondemand.com", "Custom IAS tenant for apps", "Description for terraformint.accounts400.ondemand.com", "custom link text", false, false, "inactive"),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestMatchResourceAttr("btp_subaccount_trust_configuration.uut", "subaccount_id", regexpValidUUID),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "identity_provider", "terraformtest.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "domain", "terraformtest.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "name", "Custom IAS tenant for apps"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "description", "Description for terraformint.accounts400.ondemand.com"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "link_text", "custom link text"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "available_for_user_logon", "false"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "auto_create_shadow_users", "false"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "origin", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "id", "sap.custom"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "type", "Application"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "protocol", "OpenID Connect"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "status", "inactive"),
						resource.TestCheckResourceAttr("btp_subaccount_trust_configuration.uut", "read_only", "false"),
					),
				},
				{
					ResourceName:      "btp_subaccount_trust_configuration.uut",
					ImportStateIdFunc: getTrustConfigIdForImport("btp_subaccount_trust_configuration.uut"),
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("error path - import failure", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_trust_configuration.import_error")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProviderFor(user) + hclResourceSubaccountTrustConfigurationMinimal("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", "terraformint.accounts400.ondemand.com"),
				},
				{
					ResourceName:      "btp_subaccount_trust_configuration.uut",
					ImportStateId:     "ef23ace8-6ade-4d78-9c1f-8df729548bbf",
					ImportState:       true,
					ImportStateVerify: true,
					ExpectError:       regexp.MustCompile(`Unexpected Import Identifier`),
				},
			},
		})
	})

	t.Run("error path - missing identity provider", func(t *testing.T) {
		rec, user := setupVCR(t, "fixtures/resource_subaccount_trust_configuration.error_missing_identityprovider")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config:      hclProviderFor(user) + hclResourceSubaccountTrustConfigurationMinimal("uut", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", ""),
					ExpectError: regexp.MustCompile(`Empty Identity Provider`),
				},
			},
		})
	})
}

func hclResourceSubaccountTrustConfigurationComplete(resourceName string, subaccountId string, identityProvider string, domain string, name string, description string, linkText string, availableForUserLogin bool, autoCreateShadowUsers bool, status string) string {
	template := `
resource "btp_subaccount_trust_configuration" "%s" {
    subaccount_id            = "%s"
    identity_provider        = "%s"
    domain                   = "%s"
    name                     = "%s"
    description              = "%s"
    link_text                = "%s"
    available_for_user_logon = %t
    auto_create_shadow_users = %t
    status                   = "%s"
}`

	return fmt.Sprintf(template, resourceName, subaccountId, identityProvider, domain, name, description, linkText, availableForUserLogin, autoCreateShadowUsers, status)
}

func hclResourceSubaccountTrustConfigurationMinimal(resourceName string, subaccountId string, identityProvider string) string {
	template := `
resource "btp_subaccount_trust_configuration" "%s" {
    subaccount_id 		= "%s"
    identity_provider  	= "%s"
}`

	return fmt.Sprintf(template, resourceName, subaccountId, identityProvider)
}

func getTrustConfigIdForImport(resourceName string) resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("%s,%s", "ef23ace8-6ade-4d78-9c1f-8df729548bbf", rs.Primary.ID), nil
	}
}
