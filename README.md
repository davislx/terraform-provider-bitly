# terraform-provider-bitly
A terraform provider for Bitly, this is a example terraform provider to show case the creation of Terraform provider.

## How to Run the example
Prep work

- Install Azure CLI
    - Ubuntu: `sudo apt install azure-cli`
    - OS X: `brew install azure-cli`

- Azure login
    - `az login`

- Get the bitly token
    - Login to bitly, top right cornor "${Username}" -> "Profile Settings" -> "Generate Access Token"
    - Set the token in shell: `export BITLY_TOKEN=xxxxxxxx`


Now, run the usual terraform command to provision
```bash
cd example
terraform init
# [Optional] Plan to see what will be created
terraform plan

terraform apply
```

## Implementation Notes
This provider uses the high level Terraform provider SDK, [a.k.a the"Schema" API](https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/schema).

The bitly golang client is [github.com/retgits/bitly](https://github.com/retgits/bitly), which is a thin wrapper around [Bitly v4 API](https://dev.bitly.com/docs/getting-started/introduction).

For other useful info, read more at the [offcial HashiCorp Tutorial](https://learn.hashicorp.com/collections/terraform/providers).

## Directory Structure

```
.
├── bitly
│   ├── provider.go
│   └── resource_link.go
├── example
│   ├── azure.tf
│   ├── crash.log
│   ├── main.tf
│   ├── terraform.tfstate
│   └── terraform.tfstate.backup
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
└── terraform-provider-bitly
```