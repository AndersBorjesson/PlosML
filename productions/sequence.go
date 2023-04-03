package productions

import (
	"path"
	"ploshml/semanticresolver"

	"gonum.org/v1/gonum/graph"
)

func Sequencediag(g graph.Directed, processes []graph.Node, procgraph []ProcessGraph, outpath, projname string) {
	seq := ""
	seq = seq + "@startuml\n"
	pumlid := make(map[int64]string)
	for _, l1 := range processes {

		if l1 != nil {
			id := RandString(15)
			pumlid[l1.ID()] = id
			name := tohumanname(l1.(semanticresolver.NodeType).Handler)
			artifact_type := "participant"
			artifact := artifact_type + " \"" + name + "\" as " + id + "\n"
			seq = seq + artifact
		}

	}
	for _, l1 := range procgraph {
		if l1.From != l1.To {
			node := g.Node(l1.Mid)
			delivery := tohumanname(node.(semanticresolver.NodeType).HumanName)
			left := pumlid[l1.From]
			right := pumlid[l1.To]

			artifact := left + " --> " + right + " :\"" + delivery + "\"\n"
			seq = seq + artifact
		}
	}
	seq = seq + "@enduml\n"

	filename := path.Join(outpath, projname+"_sequence_diagram.puml")
	savestring2file(filename, seq)
}
