---
page_title: "morpheus_cloud_formation_spec_template Resource - terraform-provider-morpheus"
subcategory: ""
description: |-
  Provides a Morpheus cloud formation spec template resource
---

# morpheus_cloud_formation_spec_template

Provides a Morpheus cloud formation spec template resource

## Example Usage

Creating the cloud formation spec template with local content:

```terraform
resource "morpheus_cloud_formation_spec_template" "tfexample_cloud_formation_spec_template_local" {
  name         = "tf-cloud-formation-spec-example-local"
  category     = ""
  source_type  = "local"
  spec_content = <<TFEOF

TFEOF

}
```

Creating the cloud formation spec template with the template fetched from a url:

```terraform
resource "morpheus_cloud_formation_spec_template" "tfexample_cloud_formation_spec_template_url" {
  name        = "tf-cloud-formation-spec-example-url"
  source_type = "url"
  spec_path   = "http://example.com/spec.yaml"
}
```

Creating the cloud formation spec template with the template fetched via git:

```terraform
resource "morpheus_cloud_formation_spec_template" "tfexample_cloud_formation_spec_template_git" {
  name          = "tf-cloud-formation-spec-example-git"
  source_type   = "repository"
  repository_id = 2
  version_ref   = "main"
  spec_path     = "./spec.yaml"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the cloud formation spec template
- `source_type` (String) The source of the cloud formation spec template (local, url or repository)

### Optional

- `capability_auto_expand` (Boolean) Whether cloud init is enabled
- `capability_iam` (Boolean) Whether cloud init is enabled
- `capability_named_iam` (Boolean) Whether cloud init is enabled
- `repository_id` (Number) The ID of the git repository integration
- `spec_content` (String) The content of the cloud formation spec template. Used when the local source type is specified
- `spec_path` (String) The path of the cloud formation spec template, either the url or the path in the repository
- `version_ref` (String) The git reference of the repository to pull (main, master, etc.)

### Read-Only

- `id` (String) The ID of the cloud formation spec template

## Import

Import is supported using the following syntax:

```shell
terraform import morpheus_cloud_formation_spec_template.tf_example_cloud_formation_spec_template 1
```