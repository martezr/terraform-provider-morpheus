#!/bin/bash

cd ..
go build -o ~/.terraform.d/plugins/github.com/morpheusdata/morpheus/0.0.1/darwin_amd64/terraform-provider-morpheus
cd examples
#cp ~/.terraform.d/plugins/github.com/morpheusdata/morpheus/0.0.1/darwin_amd64/terraform-provider-morpheus .terraform/plugins/github.com/morpheusdata/morpheus/0.0.1/darwin_amd64/
terraform init
terraform apply --auto-approve
