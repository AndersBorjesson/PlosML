package productions

import (
	"path"
	"ploshml/semanticresolver"

	// "github.com/dominikbraun/graph"
	// "github.com/dominikbraun/graph/draw"
	// "github.com/goccy/go-graphviz"
	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"gonum.org/v1/gonum/graph"
)

// var a *[]byte

// type wr struct {
// 	B []byte
// 	a *int
// }

// func (w wr) Init(a *int) {
// 	w.B = make([]byte, 100)
// 	w.a = a
// }
// func (w wr) Write(p []byte) (n int, err error) {
// 	fmt.Println("WB", w.B)
// 	for _, l1 := range p {
// 		fmt.Println("L!", l1)
// 		// *w.a = append(*w.a, l1)
// 	}
// 	fmt.Println("WB", w.a)
// 	fmt.Println("P", p)
// 	return len(p), nil
// }

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

// func tmpmain() {
// 	A := simple.DirectedGraph{}
// 	fmt.Println(A)
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
// 	// // }
// 	G := NewGraphNode(0)
// 	e := NewGraphNode(10)
// 	f := NewGraphNode(120)
// 	e.AddRoot(f)
// 	G.AddRoot(e)
// 	bitar, _ := enc.Marshal(G, "Arne", "", "")
// 	fmt.Println(string(bitar))
// 	graph, _ := graphviz.ParseBytes(bitar)
// 	g2 := graphviz.New()
// 	graph2, _ := g2.Graph()

// 	node, _ := graph2.CreateNode("Arne")
// 	n, _ := graph2.CreateNode("n")
// 	n.SetShape("box")
// 	n.SetColor("blue")
// 	m, _ := graph2.CreateNode("m")

// 	e2, _ := graph2.CreateEdge("e", n, m)
// 	graph2.CreateEdge("22e", node, m)
// 	e2.SetLabel("e")
// 	fmt.Println(graph)
// 	fmt.Println(node)
// 	g2.RenderFilename(graph2, graphviz.PNG, "graph.png")
// }

func getarrows() (map[string]cgraph.ArrowType, map[string]cgraph.ArrowType) {

	heads := map[string]cgraph.ArrowType{
		"triggers":    "dot",
		"handles":     "dot",
		"relates":     "none",
		"requires":    "none",
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
		"requires":    "odot",
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

func makegv_2(relations []semanticresolver.Relation, items []semanticresolver.Item) {
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
		e.SetDir("both")
		e.SetArrowHead("none")
		// fmt.Println(l1.Typ)
		// fmt.Println(heads[l1.Typ], tails[l1.Typ])
		e.SetArrowHead(heads[l1.Typ])
		e.SetArrowTail(tails[l1.Typ])
		if l1.Typ == "triggers" {
			e.SetHeadLabel("e")
		}
	}
	gv2 := graph.SubGraph("cluster_AAA", 2)
	gv2.CreateNode("Klasse")

	gv.RenderFilename(graph, graphviz.PNG, "graph.png")

}

func Makegv(g graph.Directed, outpath, projname string) {
	gv := graphviz.New()
	graph, _ := gv.Graph()

	nodes := make(map[string]*cgraph.Node)
	heads, tails := getarrows()
	tmp := g.Nodes()
	node_ids := []int64{}
	len_tmp := tmp.Len()
	for l1 := 0; l1 < len_tmp; l1++ {
		tmp.Next()
		n := tmp.Node().(semanticresolver.NodeType)
		node_ids = append(node_ids, n.Id)
		nodes[n.Name], _ = graph.CreateNode(n.Name)
		nodes[n.Name].SetLabel(n.HumanName)
		if n.Type == "P" {
			nodes[n.Name].SetColor("blue")
			nodes[n.Name].SetShape("oval")
		} else if n.Type == "O" {
			nodes[n.Name].SetColor("green")
			nodes[n.Name].SetShape("box")
		} else {
			nodes[n.Name].SetColor("black")
			nodes[n.Name].SetShape("none")
		}

	}
	order, _ := DeriveOrder(g)

	for _, l1 := range order {
		id := l1.(semanticresolver.NodeType).Id
		for _, l2 := range node_ids {
			if l2 != id {
				has_edge := g.HasEdgeFromTo(id, l2)
				if has_edge {
					edge := g.Edge(id, l2)
					F := edge.(semanticresolver.EdgeType).F
					T := edge.(semanticresolver.EdgeType).T
					e, _ := graph.CreateEdge("", nodes[F.(semanticresolver.NodeType).Name], nodes[T.(semanticresolver.NodeType).Name])
					e.SetDir("both")
					e.SetArrowHead("none")
					// fmt.Println(l1.Typ)
					// fmt.Println(edge.(EdgeType).Type, heads[edge.(EdgeType).Type], tails[edge.(EdgeType).Type])
					e.SetArrowHead(heads[edge.(semanticresolver.EdgeType).Type])
					e.SetArrowTail(tails[edge.(semanticresolver.EdgeType).Type])
					if edge.(semanticresolver.EdgeType).Type == "triggers" {
						e.SetHeadLabel("e")
					}
				}
			}
		}
	}
	basename := path.Join(outpath, projname+"_base_yield")

	gv.RenderFilename(graph, graphviz.PNG, basename+".png")
	// gv.RenderFilename(graph, graphviz.SVG, basename+".svg")
	// gv.RenderFilename(graph, graphviz.JPG, basename+".jpg")
	// gv.RenderFilename(graph, graphviz.XDOT, basename+".dot")

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
