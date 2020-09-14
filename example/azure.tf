// Create a random name to be used with example
resource "random_pet" "name" {}

resource "azurerm_resource_group" "example" {
  name     = random_pet.name.id
  location = "westus"
}

resource "azurerm_app_service_plan" "example" {
  name                = "${random_pet.name.id}-service-plan"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  kind                = "Linux"
  reserved            = true

  sku {
    tier = "Free"
    size = "F1"
  }
}

resource "azurerm_app_service" "example" {
  name                = "${random_pet.name.id}-web"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id
  https_only          = true
  site_config {
    # 64 bit not supported in F1
    use_32_bit_worker_process = true
    linux_fx_version          = "DOCKER|nginxdemos/hello:latest"
  }

  logs {
    application_logs {
      file_system_level = "Information"
    }
  }
}