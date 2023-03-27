package productions

import (
	"fmt"
	"os"

	"ploshml/semanticresolver"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

func Produce(g graph.Directed) {
	processes, _, procgraph := ResolveStructure(g)
	SWartifactdiagram(g, processes, procgraph)
}
func ResolveStructure(g graph.Directed) ([]graph.Node, bool, []ProcessGraph) {
	processes, objects := ResolveProcessesAndObjects(g)
	order, cyclic := DeriveOrder(g)
	procgraph := ResolveProcessOrder(order, processes, objects, g)
	return processes, cyclic, procgraph
}

type ProcessGraph struct {
	To   int64
	From int64
	Mid  int64
	Typ  string
}

func DeriveOrder(g graph.Directed) ([]graph.Node, bool) {
	order, err := topo.Sort(g)
	if err == nil {
		return order, false
	}
	nodes := g.Nodes()
	no_nodes := nodes.Len()
	order2 := make([]graph.Node, no_nodes)
	idx := 0
	for nodes.Len() > 0 {
		nodes.Next()
		order2[idx] = nodes.Node()
	}

	return order2, true
}
func ResolveProcessOrder(order []graph.Node, processes, object []graph.Node, g graph.Directed) []ProcessGraph {
	// order, err := topo.Sort(g)

	no_nodes := len(order)
	procgraph := make([]ProcessGraph, no_nodes)
	node_ids := make([]int64, no_nodes)
	for l1 := 0; l1 < no_nodes; l1++ {
		order[l1].ID()
		node_ids[l1] = order[l1].ID()
	}
	idx := 0
	for _, l1 := range node_ids {
		to := g.To(l1)
		if g.Node(l1).(semanticresolver.NodeType).Type == "P" {
			for to.Len() > 0 {
				to.Next()
				if to.Node() != nil {
					tn := to.Node().(semanticresolver.NodeType).ID()
					E := g.Edge(tn, l1).(semanticresolver.EdgeType).Type
					if E == "triggers" {
						procgraph[idx] = ProcessGraph{From: l1, To: tn, Typ: "triggers"}
						idx++
					}
				}
			}
		} else if g.Node(l1).(semanticresolver.NodeType).Type == "O" {
			from := g.From(l1)
			for to.Len() > 0 {
				to.Next()
				for from.Len() > 0 {
					from.Next()
					if (to.Node() != nil) && (from.Node() != nil) {
						tn := to.Node().(semanticresolver.NodeType).ID()
						fn := from.Node().(semanticresolver.NodeType).ID()
						E1 := g.Edge(tn, l1).(semanticresolver.EdgeType)
						E2 := g.Edge(l1, fn).(semanticresolver.EdgeType)
						v1 := (E1.Type == "->") && (E2.Type == ">-")
						v2 := (to.Node().(semanticresolver.NodeType).Type == "P") && (from.Node().(semanticresolver.NodeType).Type == "P")
						if v1 && v2 {
							procgraph[idx] = ProcessGraph{From: tn, To: fn, Typ: "->", Mid: l1}
							idx++
						}
					}
				}
			}

		} else {
			fmt.Println("Unknown error")
			os.Exit(1)
		}

	}
	// if err != nil {
	// 	fmt.Println("Handle error here")
	// 	os.Exit(1)
	// }

	// // nodes := make(map[string]([]graph.Node))
	// ind := 0
	// for l1 := ind; l1 < len(order); l1++ {
	// 	for l2 := ind; l2 < len(order); l2++ {
	// 		exists := topo.PathExistsIn(g, order[l1], order[l2])
	// 		fmt.Println(exists)
	// 	}
	// }
	return procgraph
}
func ResolveProcessesAndObjects(g graph.Directed) ([]graph.Node, []graph.Node) {
	order, err := topo.Sort(g)
	if err != nil {
		fmt.Println("Handle error here")
		os.Exit(1)
	}
	processes := make([]graph.Node, len(order))
	objects := make([]graph.Node, len(order))
	no_proc := 0
	no_obj := 0
	for _, l1 := range order {
		if l1.(semanticresolver.NodeType).Type == "P" {
			processes[no_proc] = l1
			no_proc += 1
		} else if l1.(semanticresolver.NodeType).Type == "O" {
			objects[no_obj] = l1
			no_obj += 1
		}
	}
	return processes, objects

}
