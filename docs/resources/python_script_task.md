---
page_title: "morpheus_python_script_task Resource - terraform-provider-morpheus"
subcategory: ""
description: |-
  Provides a Morpheus python script task resource
---

# morpheus_python_script_task

Provides a Morpheus python script task resource

## Example Usage

Creating the python script task with local script content:

```terraform
resource "morpheus_python_script_task" "tfexample_python_local" {
  name                = "tfexample_python_local"
  code                = "tfexample_python_local"
  source_type         = "local"
  script_content      = <<EOF
print('morpheus')
print('python')
EOF
  command_arguments   = "example"
  additional_packages = "pyyaml"
  python_binary       = "/usr/bin/python3"
  retryable           = true
  retry_count         = 1
  retry_delay_seconds = 10
  allow_custom_config = true
}
```

Creating the python script task with the script fetched from a url:

```terraform
resource "morpheus_python_script_task" "tfexample_python_url" {
  name                = "tfexample_python_url"
  code                = "tfexample_python_url"
  source_type         = "url"
  result_type         = "json"
  script_path         = "https://example.com/example.py"
  command_arguments   = "example"
  additional_packages = "pyyaml"
  python_binary       = "/usr/bin/python3"
  retryable           = true
  retry_count         = 1
  retry_delay_seconds = 10
  allow_custom_config = true
}
```

Creating the python script task with the script fetch via git:

```terraform
resource "morpheus_python_script_task" "tfexample_python_git" {
  name                = "tfexample_python_git"
  code                = "tfexample_python_git"
  source_type         = "repository"
  result_type         = "json"
  script_path         = "example.py"
  version_ref         = "master"
  repository_id       = 1
  command_arguments   = "example"
  additional_packages = "pyyaml"
  python_binary       = "/usr/bin/python3"
  retryable           = true
  retry_count         = 1
  retry_delay_seconds = 10
  allow_custom_config = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the python script task
- **source_type** (String) The source of the python script (local, url or repository)

### Optional

- **additional_packages** (String) Additional python packages to install prior to the execution of the python script
- **allow_custom_config** (Boolean) Custom configuration data to pass during the execution of the python script
- **code** (String) The code of the python script task
- **command_arguments** (String) Arguments to pass to the python script
- **python_binary** (String) The system path of the python binary to execute
- **repository_id** (Number) The ID of the git repository integration
- **result_type** (String) The expected result type (single value, key pairs, json)
- **retry_count** (Number) The number of times to retry the task if there is a failure
- **retry_delay_seconds** (Number) The number of seconds to wait between retry attempts
- **retryable** (Boolean) Whether to retry the task if there is a failure
- **script_content** (String) The content of the python script. Used when the local source type is specified
- **script_path** (String) The path of the python script, either the url or the path in the repository
- **version_ref** (String) The git reference of the repository to pull (main, master, etc.)

### Read-Only

- **id** (String) The ID of the python script task

## Import

Import is supported using the following syntax:

```shell
terraform import morpheus_checkbox_option_type.tf_example_checkbox_option_type 1
```