terraform {
  required_providers {
    pixela = {
      versions = ["0.1"]
      source   = "github.com/budougumi0617/pixela"
    }
  }
}

provider pixela {}

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