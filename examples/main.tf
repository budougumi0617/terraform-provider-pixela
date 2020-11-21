terraform {
  required_providers {
    pixela = {
      versions = ["0.1"]
      source   = "github.com/budougumi0617/pixela"
    }
  }
}

provider pixela {
  version = "0.1.0"
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

module graph {
  source = "./graph"

  graph_name = "egiu"
}

output egiu {
  value = module.graph.a_graph
}

output all {
  value = module.graph.all_graphs
}