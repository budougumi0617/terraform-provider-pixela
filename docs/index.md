# Pixela Provider

Lifecycle management of Pixela resources, including graph.
**This provider is not official.** 

- https://pixe.la

## Preparing
The provider needs to be configured with the proper credentials before it can be used.
Please set below environment values before each `terraform` command.

|Key|Value|
|---|---|
|`PIXELA_TOKEN`|A token string used to authenticate as `username`|

`PIXELA_TOKEN` is a token for the created pixela user.
Related pixela endpoint is [POST - /v1/users][post_user]

[post_user]: https://docs.pixe.la/entry/post-user


Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
terraform {
  required_providers {
    pixela = {
      versions = ">= 0.0.3"
      source   = "budougumi0617/pixela"
    }
  }
}

provider pixela {
  username = "PIXELA_USERNAME"
}

resource "pixela_graph" "sample" {
  graph_id              = "sample"
  name                  = "sample from terraform"
  unit                  = "page"
  type                  = "int"
  color                 = "ajisai"
  timezone              = "Asia/Tokyo"
  self_sufficient       = "none"
  is_secret             = true
  publish_optional_data = false
}
```

## Argument Reference

* Not be defined yet...