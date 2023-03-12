package productions

import (
	"fmt"
	"os"
	"ploshml/semanticresolver"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

func ResolveStructure(g graph.Directed) []graph.Node {
	processes, objects := ResolveProcessesAndObjects(g)
	ResolveProcessOrder(processes, objects, g)
	return processes
}

func ResolveProcessOrder(processes, object []graph.Node, g graph.Directed) {
	// order, err := topo.Sort(g)
	nodes := g.Nodes()
	no_nodes := nodes.Len()
	node_ids := make([]int64, no_nodes)
	for l1 := 0; l1 < no_nodes; l1++ {
		nodes.Next()
		node_ids[l1] = nodes.Node().ID()
	}
	for _, l1 := range node_ids {
		// n_now := g.Node(l1)
		to := g.To(l1)
		from := g.From(l1)
		fmt.Println(to, from)
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
