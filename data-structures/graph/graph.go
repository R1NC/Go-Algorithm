package main

import (
	"fmt"
	"github.com/RincLiu/Go-Algorithm/data-structures/queue"
)

type Vertex struct {
	Label string
	Edges []*Edge
	isVisited bool	
}

type Edge struct {
	FromVertex *Vertex
	ToVertex *Vertex
	Weight int
	isUsed bool
}

type Graph struct {
	Vertexes []*Vertex
	FirstVertex *Vertex	
}

func (graph *Graph) HasVertex(vertex *Vertex) bool {
	if graph.Vertexes == nil || len(graph.Vertexes) == 0 {
		panic("Graph has no vertexes.")	
	}
	for _, v := range graph.Vertexes {
		if v == vertex {
			return true
		}
	}
	return false
}

func (graph *Graph) HasPathBetweenVertexes(v1 *Vertex, v2 *Vertex) bool {
	if v1 == v2 {
		panic("Must be two different vertexes.")
	}
	if !graph.HasVertex(v1) {
		panic("v1 is not in graph.")
	}
	if !graph.HasVertex(v2) {
		panic("v2 is not in graph.")
	}
		
	//TODO
	return false	
}

func (graph *Graph) clearVisitHistory() {
	for _, v := range graph.Vertexes {
		for _, e := range v.Edges {
			e.isUsed = false
		}
		v.isVisited = false
        }
}

func main() {
	Va := &Vertex{}
	Va.Label = "A"

	Vb := &Vertex{}
	Vb.Label = "B"

	Vc := &Vertex{}
        Vc.Label = "C"

	Vd := &Vertex{}
        Vd.Label = "D"

	Ve := &Vertex{}
        Ve.Label = "E"

	Vf := &Vertex{}
        Vf.Label = "F"

	Vg := &Vertex{}
        Vg.Label = "G"

	Eab := &Edge{}
	Eab.FromVertex = Va
	Eab.ToVertex = Vb
	Eab.Weight = 2

	Eac := &Edge{}
        Eac.FromVertex = Va
        Eac.ToVertex = Vc
        Eac.Weight = 4

	Ead := &Edge{}
        Ead.FromVertex = Va
        Ead.ToVertex = Vd
        Ead.Weight = 1

	Va.Edges = []*Edge{Eab, Eac, Ead}
	
	Eba := &Edge{}
        Eba.FromVertex = Vb
        Eba.ToVertex = Va
        Eba.Weight = 2

	Ebd := &Edge{}
        Ebd.FromVertex = Vb
        Ebd.ToVertex = Vd
        Ebd.Weight = 3

	Ebe := &Edge{}
        Ebe.FromVertex = Vb
        Ebe.ToVertex = Ve
        Ebe.Weight = 10

	Vb.Edges = []*Edge{Eba, Ebd, Ebe}
	
	Eca := &Edge{}
        Eca.FromVertex = Vc
        Eca.ToVertex = Va
        Eca.Weight = 4

	Ecd := &Edge{}
        Ecd.FromVertex = Vc
        Ecd.ToVertex = Vd
        Ecd.Weight = 2

	Ecf := &Edge{}
        Ecf.FromVertex = Vc
        Ecf.ToVertex = Vf
        Ecf.Weight = 5

	Vc.Edges = []*Edge{Eca, Ecd, Ecf}
	
	Eda := &Edge{}
        Eda.FromVertex = Vd
        Eda.ToVertex = Va
        Eda.Weight = 1

	Edb := &Edge{}
        Edb.FromVertex = Vd
        Edb.ToVertex = Vb
        Edb.Weight = 3

	Edc := &Edge{}
        Edc.FromVertex = Vd
        Edc.ToVertex = Vc
        Edc.Weight = 2

	Ede := &Edge{}
        Ede.FromVertex = Vd
        Ede.ToVertex = Ve
        Ede.Weight = 7

	Edf := &Edge{}
        Edf.FromVertex = Vd
        Edf.ToVertex = Vf
        Edf.Weight = 8

	Edg := &Edge{}
        Edg.FromVertex = Vd
        Edg.ToVertex = Vg
        Edg.Weight = 4

	Vd.Edges = []*Edge{Eda, Edb, Edc, Ede, Edf, Edg}

	Eeb := &Edge{}
        Eeb.FromVertex = Ve
        Eeb.ToVertex = Vb
        Eeb.Weight = 10

	Eed := &Edge{}
        Eed.FromVertex = Ve
        Eed.ToVertex = Vd
        Eed.Weight = 7

	Eeg := &Edge{}
        Eeg.FromVertex = Ve
        Eeg.ToVertex = Vg
        Eeg.Weight = 6

	Ve.Edges = []*Edge{Eeb, Eed, Eeg}

	Efc := &Edge{}
        Efc.FromVertex = Vf
        Efc.ToVertex = Vc
        Efc.Weight = 5

	Efd := &Edge{}
        Efd.FromVertex = Vf
        Efd.ToVertex = Vd
        Efd.Weight = 8

	Efg := &Edge{}
        Efg.FromVertex = Vf
        Efg.ToVertex = Vg
        Efg.Weight = 1
	
	Vf.Edges = []*Edge{Efc, Efd, Efg}

	Egd := &Edge{}
        Egd.FromVertex = Vg
        Egd.ToVertex = Vd
        Egd.Weight = 4

	Ege := &Edge{}
        Ege.FromVertex = Vg
        Ege.ToVertex = Ve
        Ege.Weight = 6

	Egf := &Edge{}
        Egf.FromVertex = Vg
        Egf.ToVertex = Vf
        Egf.Weight = 1
	
	Vg.Edges = []*Edge{Egd, Ege, Egf}

	graph := &Graph{}
	graph.Vertexes = []*Vertex{Va, Vb, Vc, Vd, Ve, Vf, Vg}
	graph.FirstVertex = Va

	fmt.Println("hasVertexD:", graph.HasVertex(Vd))
	fmt.Println("hasPathBetweenAE:", graph.HasPathBetweenVertexes(Va, Ve))
}
