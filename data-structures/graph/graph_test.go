package graph

import (
	"fmt"
	"testing"
)

func Test_graph(t *testing.T) {
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

	unDirectedGraph := &Graph{}
	unDirectedGraph.Vertices = []*Vertex{Va, Vb, Vc, Vd, Ve, Vf, Vg}
	
	fmt.Println("BreadthFirstSearch:")
	unDirectedGraph.BreadthFirstSearch(Va)

	fmt.Println("\nDepthFirstSearch:")
	unDirectedGraph.DepthFirstSearch(Va)

	fmt.Println("\nPrimMinimumSpanningTree:")
	unDirectedGraph.PrimMinimumSpanningTree(Va)

	fmt.Println("KruskalMinimumSpanningTree:")
	unDirectedGraph.KruskalMinimumSpanningTree()

	fmt.Println("DijkstraShortestPath from A to G:")
	unDirectedGraph.DijkstraShortestPath(Va, Vg)

	Vh := &Vertex{}
	Vh.Label = "H"

	Vi := &Vertex{}
	Vi.Label = "I"

	Vj := &Vertex{}
	Vj.Label = "J"

	Vk := &Vertex{}
	Vk.Label = "K"

	Vl := &Vertex{}
	Vl.Label = "L"

	Vm := &Vertex{}
	Vm.Label = "M"

	Vn := &Vertex{}
	Vn.Label = "N"

	Ehj := &Edge{}
	Ehj.FromVertex = Vh
	Ehj.ToVertex = Vj

	Vh.Edges = []*Edge{Ehj}

	Eij := &Edge{}
	Eij.FromVertex = Vi
	Eij.ToVertex = Vj

	Eik := &Edge{}
	Eik.FromVertex = Vi
	Eik.ToVertex = Vk

	Eil := &Edge{}
	Eil.FromVertex = Vi
	Eil.ToVertex = Vl

	Vi.Edges = []*Edge{Eij, Eik, Eil}

	Ejn := &Edge{}
	Ejn.FromVertex = Vj
	Ejn.ToVertex = Vn

	Vj.Edges = []*Edge{Ejn}

	Ekm := &Edge{}
	Ekm.FromVertex = Vk
	Ekm.ToVertex = Vm

	Ekn := &Edge{}
	Ekn.FromVertex = Vk
	Ekn.ToVertex = Vn

	Vk.Edges = []*Edge{Ekm, Ekn}

	Elm := &Edge{}
	Elm.FromVertex = Vl
	Elm.ToVertex = Vm

	Vl.Edges = []*Edge{Elm}

	directedGraph := &Graph{}
	directedGraph.Vertices = []*Vertex{Vh, Vi, Vj, Vk, Vl, Vm, Vn}

	fmt.Println("TopologicalSort:")
	directedGraph.TopologicalSort()
}