terraform {
  required_providers {
    morpheus = {
      version = "~> 0.0.1"
      source  = "github.com/morpheusdata/morpheus"
    }
  }
}

provider "morpheus" {
  url          = ""
  username     = ""
  password     = ""
}
