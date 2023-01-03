package main

import (
	"fmt"

	// "github.com/dominikbraun/graph"
	// "github.com/dominikbraun/graph/draw"
	// "github.com/goccy/go-graphviz"
	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"gonum.org/v1/gonum/graph"
	enc "gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/simple"
)

var a *[]byte

type wr struct {
	B []byte
	a *int
}

func (w wr) Init(a *int) {
	w.B = make([]byte, 100)
	w.a = a
}
func (w wr) Write(p []byte) (n int, err error) {
	fmt.Println("WB", w.B)
	for _, l1 := range p {
		fmt.Println("L!", l1)
		// *w.a = append(*w.a, l1)
	}
	fmt.Println("WB", w.a)
	fmt.Println("P", p)
	return len(p), nil
}

// func makegraphviz() {
// 	g := graph.New(graph.IntHash, graph.Directed())

// 	_ = g.AddVertex(1, graph.VertexAttribute("style", "filled"), graph.VertexAttribute("fillcolor", "red"), graph.VertexAttribute("label", "red"))
// 	_ = g.AddVertex(2)
// 	_ = g.AddVertex(3)
// 	_ = g.AddVertex(4)
// 	_ = g.AddVertex(5)

// 	_ = g.AddEdge(1, 2)
// 	_ = g.AddEdge(1, 4)
// 	_ = g.AddEdge(2, 3)
// 	_ = g.AddEdge(2, 4)
// 	_ = g.AddEdge(2, 5)
// 	_ = g.AddEdge(3, 5)
// 	// var file wr
// 	buf := new(bytes.Buffer)
// 	// a := 1
// 	// file.Init(&a)
// 	// file, _ := os.Create("./mygraph.gv")
// 	_ = draw.DOT(g, buf)
// 	fmt.Println(buf)
// 	graph, _ := graphviz.ParseBytes(buf.Bytes())
// 	g2 := graphviz.New()
// 	g2.RenderFilename(graph, graphviz.PNG, "graph.png")
// }

// GraphNode is a node in an implicit graph.
type GraphNode struct {
	id        int64
	neighbors []graph.Node
	roots     []*GraphNode
}

// NewGraphNode returns a new GraphNode.
func NewGraphNode(id int64) *GraphNode {
	return &GraphNode{id: id}
}

// Node allows GraphNode to satisfy the graph.Graph interface.
func (g *GraphNode) Node(id int64) graph.Node {
	if id == g.id {
		return g
	}

	seen := map[int64]struct{}{g.id: {}}
	for _, root := range g.roots {
		if root.ID() == id {
			return root
		}

		if root.has(seen, id) {
			return root
		}
	}

	for _, n := range g.neighbors {
		if n.ID() == id {
			return n
		}

		if gn, ok := n.(*GraphNode); ok {
			if gn.has(seen, id) {
				return gn
			}
		}
	}

	return nil
}

func (g *GraphNode) has(seen map[int64]struct{}, id int64) bool {
	for _, root := range g.roots {
		if _, ok := seen[root.ID()]; ok {
			continue
		}

		seen[root.ID()] = struct{}{}
		if root.ID() == id {
			return true
		}

		if root.has(seen, id) {
			return true
		}

	}

	for _, n := range g.neighbors {
		if _, ok := seen[n.ID()]; ok {
			continue
		}

		seen[n.ID()] = struct{}{}
		if n.ID() == id {
			return true
		}

		if gn, ok := n.(*GraphNode); ok {
			if gn.has(seen, id) {
				return true
			}
		}
	}

	return false
}

// Nodes allows GraphNode to satisfy the graph.Graph interface.
func (g *GraphNode) Nodes() graph.Nodes {
	nodes := []graph.Node{g}
	seen := map[int64]struct{}{g.id: {}}

	for _, root := range g.roots {
		nodes = append(nodes, root)
		seen[root.ID()] = struct{}{}

		nodes = root.nodes(nodes, seen)
	}

	for _, n := range g.neighbors {
		nodes = append(nodes, n)
		seen[n.ID()] = struct{}{}

		if gn, ok := n.(*GraphNode); ok {
			nodes = gn.nodes(nodes, seen)
		}
	}

	return iterator.NewOrderedNodes(nodes)
}

func (g *GraphNode) nodes(dst []graph.Node, seen map[int64]struct{}) []graph.Node {
	for _, root := range g.roots {
		if _, ok := seen[root.ID()]; ok {
			continue
		}
		seen[root.ID()] = struct{}{}
		dst = append(dst, graph.Node(root))

		dst = root.nodes(dst, seen)
	}

	for _, n := range g.neighbors {
		if _, ok := seen[n.ID()]; ok {
			continue
		}

		dst = append(dst, n)
		if gn, ok := n.(*GraphNode); ok {
			dst = gn.nodes(dst, seen)
		}
	}

	return dst
}

// From allows GraphNode to satisfy the graph.Graph interface.
func (g *GraphNode) From(id int64) graph.Nodes {
	if id == g.ID() {
		return iterator.NewOrderedNodes(g.neighbors)
	}

	seen := map[int64]struct{}{g.id: {}}
	for _, root := range g.roots {
		seen[root.ID()] = struct{}{}

		if result := root.findNeighbors(id, seen); result != nil {
			return iterator.NewOrderedNodes(result)
		}
	}

	for _, n := range g.neighbors {
		seen[n.ID()] = struct{}{}

		if gn, ok := n.(*GraphNode); ok {
			if result := gn.findNeighbors(id, seen); result != nil {
				return iterator.NewOrderedNodes(result)
			}
		}
	}

	return nil
}

func (g *GraphNode) findNeighbors(id int64, seen map[int64]struct{}) []graph.Node {
	if id == g.ID() {
		return g.neighbors
	}

	for _, root := range g.roots {
		if _, ok := seen[root.ID()]; ok {
			continue
		}
		seen[root.ID()] = struct{}{}

		if result := root.findNeighbors(id, seen); result != nil {
			return result
		}
	}

	for _, n := range g.neighbors {
		if _, ok := seen[n.ID()]; ok {
			continue
		}
		seen[n.ID()] = struct{}{}

		if gn, ok := n.(*GraphNode); ok {
			if result := gn.findNeighbors(id, seen); result != nil {
				return result
			}
		}
	}

	return nil
}

// HasEdgeBetween allows GraphNode to satisfy the graph.Graph interface.
func (g *GraphNode) HasEdgeBetween(uid, vid int64) bool {
	return g.EdgeBetween(uid, vid) != nil
}

// Edge allows GraphNode to satisfy the graph.Graph interface.
func (g *GraphNode) Edge(uid, vid int64) graph.Edge {
	return g.EdgeBetween(uid, vid)
}

// EdgeBetween allows GraphNode to satisfy the graph.Graph interface.
func (g *GraphNode) EdgeBetween(uid, vid int64) graph.Edge {
	if uid == g.id || vid == g.id {
		for _, n := range g.neighbors {
			if n.ID() == uid || n.ID() == vid {
				return simple.Edge{F: g, T: n}
			}
		}
		return nil
	}

	seen := map[int64]struct{}{g.id: {}}
	for _, root := range g.roots {
		seen[root.ID()] = struct{}{}
		if result := root.edgeBetween(uid, vid, seen); result != nil {
			return result
		}
	}

	for _, n := range g.neighbors {
		seen[n.ID()] = struct{}{}
		if gn, ok := n.(*GraphNode); ok {
			if result := gn.edgeBetween(uid, vid, seen); result != nil {
				return result
			}
		}
	}

	return nil
}

func (g *GraphNode) edgeBetween(uid, vid int64, seen map[int64]struct{}) graph.Edge {
	if uid == g.id || vid == g.id {
		for _, n := range g.neighbors {
			if n.ID() == uid || n.ID() == vid {
				return simple.Edge{F: g, T: n}
			}
		}
		return nil
	}

	for _, root := range g.roots {
		if _, ok := seen[root.ID()]; ok {
			continue
		}
		seen[root.ID()] = struct{}{}
		if result := root.edgeBetween(uid, vid, seen); result != nil {
			return result
		}
	}

	for _, n := range g.neighbors {
		if _, ok := seen[n.ID()]; ok {
			continue
		}

		seen[n.ID()] = struct{}{}
		if gn, ok := n.(*GraphNode); ok {
			if result := gn.edgeBetween(uid, vid, seen); result != nil {
				return result
			}
		}
	}

	return nil
}

// ID allows GraphNode to satisfy the graph.Node interface.
func (g *GraphNode) ID() int64 {
	return g.id
}

// AddMeighbor adds an edge between g and n.
func (g *GraphNode) AddNeighbor(n *GraphNode) {
	g.neighbors = append(g.neighbors, graph.Node(n))
}

// AddRoot adds provides an entrance into the graph g from n.
func (g *GraphNode) AddRoot(n *GraphNode) {
	g.roots = append(g.roots, n)
}

func tmpmain() {
	A := simple.DirectedGraph{}
	fmt.Println(A)
	// // This example shows the construction of the following graph
	// // using the implicit graph type above.
	// //
	// // The visual representation of the graph can be seen at
	// // https://graphviz.gitlab.io/_pages/Gallery/undirected/fdpclust.html
	// //
	// // graph G {
	// // 	e
	// // 	subgraph clusterA {
	// // 		a -- b
	// // 		subgraph clusterC {
	// // 			C -- D
	// // 		}
	// // 	}
	// // 	subgraph clusterB {
	// // 		d -- f
	// // 	}
	// // 	d -- D
	// // 	e -- clusterB
	// // 	clusterC -- clusterB
	// // }

	// // graph G {
	// G := NewGraphNode(0)
	// // 	e
	// // e := NewGraphNode(1)

	// // // 	subgraph clusterA {
	// // clusterA := NewGraphNode(2)
	// // // 		a -- b
	// // a := NewGraphNode(3)
	// // b := NewGraphNode(4)
	// // a.AddNeighbor(b)
	// // b.AddNeighbor(a)
	// // clusterA.AddRoot(a)
	// // clusterA.AddRoot(b)

	// // // 		subgraph clusterC {
	// // clusterC := NewGraphNode(5)
	// // // 			C -- D
	// // C := NewGraphNode(6)
	// // D := NewGraphNode(7)
	// // C.AddNeighbor(D)
	// // D.AddNeighbor(C)

	// // clusterC.AddRoot(C)
	// // clusterC.AddRoot(D)
	// // // 		}
	// // clusterA.AddRoot(clusterC)
	// // // 	}

	// // // 	subgraph clusterB {
	// // clusterB := NewGraphNode(8)
	// // // 		d -- f
	// // d := NewGraphNode(9)
	// // f := NewGraphNode(10)
	// // d.AddNeighbor(f)
	// // f.AddNeighbor(d)
	// // clusterB.AddRoot(d)
	// // clusterB.AddRoot(f)
	// // // 	}

	// // // 	d -- D
	// // d.AddNeighbor(D)
	// // D.AddNeighbor(d)

	// // // 	e -- clusterB
	// // e.AddNeighbor(clusterB)
	// // clusterB.AddNeighbor(e)

	// // // 	clusterC -- clusterB
	// // clusterC.AddNeighbor(clusterB)
	// // clusterB.AddNeighbor(clusterC)

	// // G.AddRoot(e)
	// // G.AddRoot(clusterA)
	// // G.AddRoot(clusterB)
	// // }

	// // if topo.IsPathIn(G, []graph.Node{C, D, d, f}) {
	// // 	fmt.Println("C--D--d--f is a path in G.")
	// // }
	G := NewGraphNode(0)
	e := NewGraphNode(10)
	f := NewGraphNode(120)
	e.AddRoot(f)
	G.AddRoot(e)
	bitar, _ := enc.Marshal(G, "Arne", "", "")
	fmt.Println(string(bitar))
	graph, _ := graphviz.ParseBytes(bitar)
	g2 := graphviz.New()
	graph2, _ := g2.Graph()

	node, _ := graph2.CreateNode("Arne")
	n, _ := graph2.CreateNode("n")
	n.SetShape("box")
	n.SetColor("blue")
	m, _ := graph2.CreateNode("m")

	e2, _ := graph2.CreateEdge("e", n, m)
	graph2.CreateEdge("22e", node, m)
	e2.SetLabel("e")
	fmt.Println(graph)
	fmt.Println(node)
	g2.RenderFilename(graph2, graphviz.PNG, "graph.png")
}

func getarrows() (map[string]cgraph.ArrowType, map[string]cgraph.ArrowType) {

	heads := map[string]cgraph.ArrowType{
		"triggers":    "none",
		"handles":     "dot",
		"relates":     "none",
		"requires":    "odot",
		"aggregates":  "none",
		"implements":  "none",
		"exhibits":    "none",
		"generalizes": "none",
		"->":          "normal",
		">-":          "normal",
		"-<":          "normal",
		"<-":          "normal"}
	tails := map[string]cgraph.ArrowType{
		"triggers":    "none",
		"handles":     "none",
		"relates":     "none",
		"requires":    "none",
		"aggregates":  "none",
		"implements":  "none",
		"exhibits":    "none",
		"generalizes": "none",
		"->":          "none",
		">-":          "none",
		"-<":          "none",
		"<-":          "none"}
	return heads, tails
}

func makegv(relations []Relation, items []Item) {
	gv := graphviz.New()
	graph, _ := gv.Graph()
	nodes := make(map[string]*cgraph.Node)
	heads, tails := getarrows()
	for _, l1 := range items {
		nodes[l1.Name], _ = graph.CreateNode(l1.Name)
		if l1.Typ == "P" {
			nodes[l1.Name].SetColor("blue")
			nodes[l1.Name].SetShape("oval")
		} else if l1.Typ == "O" {
			nodes[l1.Name].SetColor("green")
			nodes[l1.Name].SetShape("box")
		} else {
			nodes[l1.Name].SetColor("black")
			nodes[l1.Name].SetShape("none")
		}

	}
	for _, l1 := range relations {
		e, _ := graph.CreateEdge("", nodes[l1.From], nodes[l1.To])
		e.SetArrowHead("none")
		// fmt.Println(l1.Typ)
		fmt.Println(heads[l1.Typ], tails[l1.Typ])
		e.SetArrowHead(heads[l1.Typ])
		e.SetArrowTail(tails[l1.Typ])
	}
	gv2 := graph.SubGraph("cluster_AAA", 2)
	gv2.CreateNode("Klasse")

	gv.RenderFilename(graph, graphviz.PNG, "graph.png")

}

// func main() {
// 	// for _, l1 := range return_keywords() {
// 	// 	fmt.Println(l1)
// 	// 	lex := lexer.New([]rune(l1))

// 	// 	_, err := parser.New(lex).Parse()

// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// }
// 	// makegraphviz()
// }
