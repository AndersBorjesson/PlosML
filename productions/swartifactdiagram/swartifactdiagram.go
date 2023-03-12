package swartifactdiagram

import (
	"fmt"
	"log"
	"os"
	"ploshml/productions"
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
func Apa(g graph.Directed) {
	processes := productions.ResolveStructure(g)
	filepath := "ARNE"

	d, err := diagram.New(Name(filepath), diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))

	if err != nil {

		log.Fatal(err)
	}
	nodes := nodetypes(processes)
	for _, l1 := range nodes {

		if l1 != nil {

			d.Add(l1)

			fmt.Println("DD", d.Nodes())
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
			tmp := atf[name]().Label(l1.(semanticresolver.NodeType).Handler)
			r[i] = tmp

		}

	}
	return r
}
