package generate_docs

import "fmt"

func LinkBuilder(s string) string {
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
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/aws/latest/docs")
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
