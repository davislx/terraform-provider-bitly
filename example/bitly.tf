resource "bitly_bitlink" "az_service_plan" {
  long_url = "http://${azurerm_app_service.example.default_site_hostname}"
}

output "az_service_plan_link" {
    value = bitly_bitlink.az_service_plan.link
}
