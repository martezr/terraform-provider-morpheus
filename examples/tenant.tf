resource "morpheus_tenant" "tftest" {
  name         = "tftest"
  description  = "A test tenant created by Terraform"
  subdomain    = "tftest"
}

resource "morpheus_tenant" "tftest2" {
  name         = "tftest2"
  description  = "A test tenant created by Terraform"
  subdomain    = "tftest2"
}
