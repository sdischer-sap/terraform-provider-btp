package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDataSourceSubaccountUser(t *testing.T) {
	t.Parallel()
	t.Run("happy path", func(t *testing.T) {
		rec := setupVCR(t, "fixtures/datasource_subaccount_user")
		defer stopQuietly(rec)

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(rec.GetDefaultClient()),
			Steps: []resource.TestStep{
				{
					Config: hclProvider() + hclDatasourceSubaccountUserWithCustomIdp("uut", "5381d6a4-d67f-45b1-93a0-624876f74d03", "jenny.doe@test.com", "sap.default"),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "subaccount_id", "5381d6a4-d67f-45b1-93a0-624876f74d03"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "user_name", "jenny.doe@test.com"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "origin", "sap.default"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "active", "true"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "family_name", "unknown"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "given_name", "unknown"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "id", "de350a51-fa8f-4bdf-bd75-79179b846911"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "role_collections.#", "0"),
						resource.TestCheckResourceAttr("data.btp_subaccount_user.uut", "verified", "false"),
					),
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
					Config:      hclProvider() + hclDatasourceSubaccountUserWithCustomIdp("uut", "this-is-not-a-uuid", "jenny.doe@test.com", "sap.default"),
					ExpectError: regexp.MustCompile(`Attribute subaccount_id value must be a valid UUID, got: this-is-not-a-uuid`),
				},
			},
		})
	})
	t.Run("error path - subaccount_id, user_name and origin are mandatory", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(nil),
			Steps: []resource.TestStep{
				{
					Config:      hclProvider() + `data "btp_directory_user" "uut" {}`,
					ExpectError: regexp.MustCompile(`The argument "(subaccount_id|user_name)" is required, but no definition was found.`),
				},
			},
		})
	})
	t.Run("error path - user_name must not be empty", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(nil),
			Steps: []resource.TestStep{
				{
					Config:      hclProvider() + hclDatasourceDirectoryUserCustomIdp("uut", "5381d6a4-d67f-45b1-93a0-624876f74d03", "", "terraformint"),
					ExpectError: regexp.MustCompile(`Attribute user_name string length must be between 1 and 256, got: 0`),
				},
			},
		})
	})
	t.Run("error path - origin must not be empty if given", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(nil),
			Steps: []resource.TestStep{
				{
					Config:      hclProvider() + hclDatasourceDirectoryUserCustomIdp("uut", "5381d6a4-d67f-45b1-93a0-624876f74d03", "jenny.doe@test.com", ""),
					ExpectError: regexp.MustCompile(`Attribute origin string length must be at least 1, got: 0`),
				},
			},
		})
	})
	t.Run("error path - cli server returns error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/login/") {
				fmt.Fprintf(w, "{}")
				return
			}
			w.WriteHeader(http.StatusNotFound)
		}))
		defer srv.Close()

		resource.Test(t, resource.TestCase{
			IsUnitTest:               true,
			ProtoV6ProviderFactories: getProviders(srv.Client()),
			Steps: []resource.TestStep{
				{
					Config:      hclProviderWithCLIServerURL(srv.URL) + hclDatasourceSubaccountUserWithCustomIdp("uut", "5381d6a4-d67f-45b1-93a0-624876f74d03", "jenny.doe@test.com", "sap.default"),
					ExpectError: regexp.MustCompile(`Received response with unexpected status \[Status: 404; Correlation ID:\s+[a-f0-9\-]+\]`),
				},
			},
		})
	})
}

func hclDatasourceSubaccountUserWithCustomIdp(resourceName string, subaccountId string, userName string, origin string) string {
	template := `
data "btp_subaccount_user" "%s" {
	subaccount_id = "%s"
	user_name 	 = "%s"
  	origin    	 = "%s"
}`

	return fmt.Sprintf(template, resourceName, subaccountId, userName, origin)
}
