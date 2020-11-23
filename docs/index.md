# Pixela Provider

Lifecycle management of Pixela resources, including graph.
**This provider is not official.** 

- https://pixe.la

## Preparing
The provider needs to be configured with the proper credentials before it can be used.
Please set below environment values before each `terraform` command.

|Key|Value|
|---|---|
|`PIXELA_USERNAME`|Created user name by [POST - /v1/users][post_user]
|`PIXELA_TOKEN`|A token string used to authenticate as `PIXELA_USERNAME`|


[post_user]: https://docs.pixe.la/entry/post-user


Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
terraform {
  required_providers {
    pixela = {
      versions = ">= 0.0.1"
      source   = "github.com/budougumi0617/pixela"
    }
  }
}

provider pixela {}

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