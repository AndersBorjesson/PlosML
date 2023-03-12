package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"ploshml/semanticresolver"

	"github.com/blushft/go-diagrams/diagram"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"

	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/blushft/go-diagrams/nodes/gcp"
)

func types() {
	tmp := make(map[string]*diagram.Node)
	tmp["dns"] = gcp.Network.Dns()
	tmp["database"] = gcp.Database.Sql()
	tmp["sql"] = gcp.Database.Sql()
	tmp["dataproc"] = gcp.Analytics.Dataproc()
	tmp["computeengine"] = gcp.Compute.ComputeEngine()
	tmp["lambda"] = aws.Compute.Lambda()
	tmp["awsdb"] = aws.Database.Database()
	tmp["sql"] = aws.Compute.Compute()
	tmp["sql"] = gcp.Database.Sql()

}
func Name(f string) diagram.Option {
	return func(o *diagram.Options) {
		o.Name = f
	}
}

func ResolveStructure(g graph.Directed) {

	order, err := topo.Sort(g)
	if err != nil {
		fmt.Println("Handle error here")
		os.Exit(1)
	}
	processes := make([]graph.Node, len(order))
	no := 0
	for _, l1 := range order {
		if l1.(semanticresolver.NodeType).Type == "P" {
			processes[no] = l1
			no += 1
		}
	}

	tt := topo.PathExistsIn(g, processes[2], processes[3])
	fmt.Println(tt)
	tt = topo.PathExistsIn(g, processes[1], processes[0])
	fmt.Println(tt)
}

func astparse() {
	t := token.NewFileSet()
	src := "./"
	f, _ := parser.ParseDir(t, src, nil, parser.AllErrors)
	for i, l1 := range f {
		fmt.Println(i, l1.Files)
	}
}
func ASW() {
	// astparse()
	// os.Exit(2)
	filepath := "ARNE"

	d, err := diagram.New(Name(filepath), diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))

	if err != nil {

		log.Fatal(err)
	}

	A := make(map[string]*diagram.Node)
	A["Sune"] = gcp.Network.Dns()
	dns := gcp.Network.Dns().Label("DNSen")
	lb := gcp.Network.LoadBalancing(diagram.NodeLabel("NLB"))
	cache := gcp.Database.Memorystore(diagram.NodeLabel("Cache"))
	db := gcp.Database.Sql(diagram.NodeLabel("Database"))

	d.Add(dns, lb, cache, db)
	d.Connect(lb, cache)

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

func ASW2() {

	d, err := diagram.New(diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))
	lb := gcp.Network.LoadBalancing(diagram.NodeLabel("NLB"))
	cache := gcp.Database.Memorystore(diagram.NodeLabel("Cache"))
	db := gcp.Database.Sql(diagram.NodeLabel("Database"))

	dc := diagram.NewGroup("GCP")
	dc.NewGroup("services").
		Label("Service Layer").
		Add(
			gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 1")),
			gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 2")),
			gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 3")),
			aws.Analytics.DataPipeline(diagram.NodeLabel("PipeLone 3")),
		).
		ConnectAllFrom(lb.ID(), diagram.Forward()).
		ConnectAllTo(cache.ID(), diagram.Forward())

	dc.NewGroup("data").Label("Data Layer").Add(cache, db).Connect(cache, db)

	d.Connect(dns, lb, diagram.Forward()).Group(dc)
	d.Render()
	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
