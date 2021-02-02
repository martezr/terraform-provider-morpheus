resource "morpheus_app" "tftest2" {
  name         = "sandbox"
  description  = "example app"
  environment  = "sandbox"
  group = {
    id = 1
  }
  blueprintid  = "existing"

  depends_on = [
    morpheus_environment.tftest2,
  ]
}
