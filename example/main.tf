terraform {
  required_version = ">=0.13"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 2.27.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 2.3.0"
    }
    bitly = {
      source  = "davislx.github.io/terraform/bitly"
      version = ">=0.0.1"
    }
  }
}

provider "azurerm" {
  features {}
}
