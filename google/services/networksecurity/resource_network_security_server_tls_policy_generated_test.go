// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package networksecurity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityServerTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_server_tls_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_server_tls_policy" "default" {
  name                   = "tf-test-my-server-tls-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  allow_open             = "false"
  server_certificate {
    certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
  }
  mtls_policy {
    client_validation_ca {
      grpc_endpoint {
        target_uri = "unix:mypath"
      }
    }
  }
}
`, context)
}

func TestAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityServerTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_security_server_tls_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_server_tls_policy" "default" {
  name                   = "tf-test-my-server-tls-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  location               = "global"
  allow_open             = "false"
  mtls_policy {
    client_validation_mode = "ALLOW_INVALID_OR_MISSING_CLIENT_CERT"
  }
}
`, context)
}

func TestAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyServerCertExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityServerTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyServerCertExample(context),
			},
			{
				ResourceName:            "google_network_security_server_tls_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyServerCertExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_server_tls_policy" "default" {
  name                   = "tf-test-my-server-tls-policy%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  location               = "global"
  allow_open             = "false"
  server_certificate {
    grpc_endpoint {
        target_uri = "unix:mypath"
      }
  }
}
`, context)
}

func TestAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyMtlsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityServerTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyMtlsExample(context),
			},
			{
				ResourceName:            "google_network_security_server_tls_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityServerTlsPolicy_networkSecurityServerTlsPolicyMtlsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_network_security_server_tls_policy" "default" {
  name     = "tf-test-my-server-tls-policy%{random_suffix}"

  description = "my description"
  location    = "global"
  allow_open  = "false"

  mtls_policy {
    client_validation_mode         = "REJECT_INVALID"
    client_validation_trust_config = "projects/${data.google_project.project.number}/locations/global/trustConfigs/${google_certificate_manager_trust_config.default.name}"
  }

  labels = {
    foo = "bar"
  }
}

resource "google_certificate_manager_trust_config" "default" {
  name        = "tf-test-my-trust-config%{random_suffix}"
  description = "sample trust config description"
  location    = "global"

  trust_stores {
    trust_anchors {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
    intermediate_cas {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
  }

  labels = {
    foo = "bar"
  }
}
`, context)
}

func testAccCheckNetworkSecurityServerTlsPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_server_tls_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkSecurityServerTlsPolicy still exists at %s", url)
			}
		}

		return nil
	}
}
