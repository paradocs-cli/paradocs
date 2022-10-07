package generate_docs

import (
	"fmt"
	"strings"
)

func ProviderLinkBuilder(s string) string {
	switch s {
	case "aws":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs")
	case "google":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/google/latest/docs")
	case "azurerm":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs")
	case "kubernetes":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs")
	case "azuredevops":
		return fmt.Sprintf("https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs")
	case "tfe":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/tfe/latest/docs")
	case "helm":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/helm/latest/docs")
	case "azuread":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/azuread/latest/docs")
	case "consul":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/consul/latest/docs")
	case "nomad":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/nomad/latest/docs")
	case "vault":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/vault/latest/docs")
	default:
		return fmt.Sprintf("https://registry.terraform.io/browse/providers")
	}
}

func ResourceLinkBuilder(s string, r string) string {
	switch s {
	case "aws":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/%s", strings.Replace(r, "aws_", "", 2))
	case "google":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/%s", strings.Replace(r, "google_", "", 2))
	case "azurerm":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/%s", strings.Replace(r, "azurerm_", "", 2))
	case "kubernetes":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/%s", strings.Replace(r, "google_", "", 2))
	case "azuredevops":
		return fmt.Sprintf("https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs/resources/%s", strings.Replace(r, "azuredevops_", "", 2))
	case "tfe":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/tfe/latest/docs/resources/%s", strings.Replace(r, "tfe_", "", 2))
	case "helm":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/helm/latest/docs/resources/%s", strings.Replace(r, "helm_", "", 2))
	case "azuread":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/%s", strings.Replace(r, "azuread_", "", 2))
	case "consul":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/consul/latest/docs/resources/%s", strings.Replace(r, "consul_", "", 2))
	case "nomad":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/nomad/latest/docs/resources/%s", strings.Replace(r, "nomad_", "", 2))
	case "vault":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/vault/latest/docs/resources/%s", strings.Replace(r, "vault_", "", 2))
	default:
		return fmt.Sprintf("https://registry.terraform.io/browse/providers")
	}
}

func DataResourceLinkBuilder(s string, r string) string {
	switch s {
	case "aws":
		strings.Replace(r, "aws_", "", 2)
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/%s", strings.Replace(r, "aws_", "", 2))
	case "google":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/%s", strings.Replace(r, "google_", "", 2))
	case "azurerm":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/%s", strings.Replace(r, "azurerm_", "", 2))
	case "kubernetes":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/data-sources/%s", strings.Replace(r, "kubernetes_", "", 2))
	case "azuredevops":
		return fmt.Sprintf("https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs/data-sources/%s", strings.Replace(r, "azuredevops_", "", 2))
	case "tfe":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/tfe/latest/docs/data-sources/%s", strings.Replace(r, "tfe_", "", 2))
	case "helm":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/helm/latest/docs/data-sources/%s", strings.Replace(r, "helm_", "", 2))
	case "azuread":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/data-sources/%s", strings.Replace(r, "azuread_", "", 2))
	case "consul":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/consul/latest/docs/data-sources/%s", strings.Replace(r, "consul_", "", 2))
	case "nomad":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/nomad/latest/docs/data-sources/%s", strings.Replace(r, "nomad_", "", 2))
	case "vault":
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/vault/latest/docs/data-sources/%s", strings.Replace(r, "vault_", "", 2))
	default:
		return fmt.Sprintf("https://registry.terraform.io/browse/providers")
	}
}
