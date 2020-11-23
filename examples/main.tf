terraform {
  required_providers {
    pixela = {
      versions = ["0.0.3"]
      source   = "github.com/budougumi0617/pixela"
    }
  }
}

provider pixela {
  username = "budougumi0617"
}

resource "pixela_graph" "sample" {
  graph_id              = "sample"
  name                  = "update from terraform"
  unit                  = "page"
  type                  = "int"
  color                 = "ajisai"
  timezone              = "Asia/Tokyo"
  self_sufficient       = "none"
  is_secret             = true
  publish_optional_data = false
}

output sample {
  value = pixela_graph.sample.graph_id
}
