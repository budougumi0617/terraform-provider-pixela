# Resource: pixela_graph 

The pixelation graph definition. 

~> Please set the environment values `PIXELA_USERNAME` and `PIXELA_TOKEN`, before each `terraform` command.

## Example Usage

```hcl
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

* `graph_id` - (Required) It is an ID for identifying the pixelation graph.
                Validation rule: `^[a-z][a-z0-9-]{1,16}`
* `name` - (Required) It is the name of the pixelation graph.
* `unit` - (Required) It is a unit of the quantity recorded in the pixelation graph. Ex. commit, kilogram, calorie.
* `type` - (Required) It is the type of quantity to be handled in the graph. Only int or float are supported.
* `color` - (Required) Defines the display color of the pixel in the pixelation graph.
                    `shibafu` (green), `momiji` (red), `sora` (blue), `ichou` (yellow), `ajisai` (purple) and `kuro` (black) are supported as color kind.
* `timezone` - (Optional) Specify the timezone for handling this graph as `Asia/Tokyo`. If not specified, it is treated as `UTC`.
* `self_sufficient` - (Optional) If SVG graph with this field increment or decrement is referenced, Pixel of this graph itself will be incremented or decremented.
                                 It is suitable when you want to record the PVs on a web page or site simultaneously.
                                 The specification of increment or decrement is the same as Increment a Pixel and Decrement a Pixel with webhook.
                                 If not specified, it is treated as `none`.
* `is_secret` - (Optional) Graphs with this property's value true are not displayed on the graph list page and can be kept secret.
* `publish_optional_data` - (Optional) If this property is `true`, each pixel's `optionalData` will be added to the generated SVG data as a `data-optional` attribute.

## Attribute Reference
There is nothing other than the above arguments values yet.