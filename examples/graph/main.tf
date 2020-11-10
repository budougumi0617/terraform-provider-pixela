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

data "pixela_graphs" "all" {}

# Returns all coffees
output "all_coffees" {
  value = data.pixela_graphs.all.graphs
}

# Only returns packer spiced latte
output "egiu" {
  value = {
    for graph in data.pixela_graphs.all.graphs :
    graph.id => graph
    if graphh.name == var.graph_name
  }
}
