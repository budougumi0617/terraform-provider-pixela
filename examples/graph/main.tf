terraform {
  required_providers {
    pixela = {
      versions = ["0.1"]
      source = "github.com/budougumi0617/pixela"
    }
  }
}

variable "graph_name" {
  type    = string
  default = "egiu"
}

data "pixela_graphs" "all" {
  id = "budougumi0617/graphs"
}

# Returns all graphs.
output "all_graphs" {
  value = data.pixela_graphs.all
}

# Only Returns graph by id.
output "a_graph" {
  value = {
    for graph in data.pixela_graphs.all.graphs :
    graph.id => graph
    if graph.id == var.graph_name
  }
}
