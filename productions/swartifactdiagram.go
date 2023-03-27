package productions

import (
	"fmt"
	"log"
	"os"
	"ploshml/semanticresolver"
	"strings"

	"github.com/blushft/go-diagrams/diagram"
	"gonum.org/v1/gonum/graph"
)

func Name(f string) diagram.Option {
	return func(o *diagram.Options) {
		o.Name = f
	}
}
func SWartifactdiagram(g graph.Directed, processes []graph.Node, procgraph []ProcessGraph) {

	filepath := "ARNE"

	d, err := diagram.New(Name(filepath), diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))

	if err != nil {

		log.Fatal(err)
	}
	translate := make(map[int64]int)
	for l1, n := range processes {
		if n != nil {
			translate[n.ID()] = l1
		}
	}
	nodes := nodetypes(processes)
	for _, l1 := range nodes {

		if l1 != nil {

			d.Add(l1)

			// fmt.Println("DD", d.Nodes())
		}
	}
	for _, n := range procgraph {
		if n.From != n.To {
			indf := translate[n.From]
			indt := translate[n.To]
			d.Connect(nodes[indf], nodes[indt], diagram.Forward())
		}
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) == false {
		err := os.RemoveAll(filepath)
		if err != nil {
			fmt.Println("Error removing directory ", filepath)
			os.Exit(1)
		}
	}

	d.Render()
	if err := d.Render(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

}

func nodetypes(processes []graph.Node) []*diagram.Node {
	r := make([]*diagram.Node, len(processes))

	atf := artifacts()
	for i, l1 := range processes {
		if l1 != nil {
			H := strings.ToLower(l1.(semanticresolver.NodeType).Handler)
			name := "Unknown"
			for l2, _ := range atf {
				contained := strings.Contains(H, l2)
				if contained {
					name = l2
					break
				}
			}
			if name == "Unknown" {
				name = "rack"
			}
			handler := tohumanname(l1.(semanticresolver.NodeType).Handler)
			tmp := atf[name]().Label(handler)
			r[i] = tmp

		}

	}
	return r
}

func tohumanname(in string) string {
	tmp := strings.Split(in, ".")
	return tmp[len(tmp)-1]
}
