package main

import (
	"fmt"
	"github.com/RincLiu/Go-Algorithm/data-structures/queue"
	"github.com/RincLiu/Go-Algorithm/data-structures/stack"
)

const MAX_INT = 999999999

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
	Vertices []*Vertex
}

func (graph *Graph) BreadFirstTraverse(startVertex *Vertex) {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	fmt.Printf("%s ", startVertex.Label)
	startVertex.isVisited = true
	queue := &queue.LinkedQueue{}
	queue.Add(startVertex)
	for queue.Size() > 0 {
		vertex := convertToVertex(queue.Peek())
		for _, edge := range vertex.Edges {
			if !edge.ToVertex.isVisited {
				fmt.Printf("%s ", edge.ToVertex.Label)
				edge.ToVertex.isVisited = true
				queue.Add(edge.ToVertex)
			}
		}	
		queue.Remove()
	}
	graph.clearVerticesVisitHistory()
}

func (graph *Graph) DepthFirstTraverse(startVertex *Vertex) {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}       
	fmt.Printf("%s ", startVertex.Label)
	startVertex.isVisited = true
	stack := &stack.LinkedStack{}
	stack.Push(startVertex)
	for stack.Size() > 0 {
		vertex := convertToVertex(stack.Peek())
		hasAddedNewVertex := false
		for _, edge := range vertex.Edges {
			if !edge.ToVertex.isVisited {
				fmt.Printf("%s ", edge.ToVertex.Label)
				edge.ToVertex.isVisited = true
				hasAddedNewVertex = true
				stack.Push(edge.ToVertex)
			}
		}
		if !hasAddedNewVertex {
			stack.Pop()
		}
	}
	graph.clearVerticesVisitHistory()
}

func (graph *Graph) PrimMinimumSpanningTree(startVertex *Vertex) {
	treeEdges := []*Edge{}
	startVertex.isVisited = true
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		minWeightEdge := getMinWeightEdgeInVertices(graph.getVisitedVertices())
		if minWeightEdge != nil {
			treeEdges = append(treeEdges, minWeightEdge)
		}
		minWeightEdge.ToVertex.isVisited = true
	}
	graph.clearVerticesVisitHistory()
	for _, edge := range treeEdges {
		fmt.Printf("%s->%s(%d)\n", edge.FromVertex.Label, edge.ToVertex.Label, edge.Weight)
	}	
}

func getMinWeightEdgeInVertices(vertices []*Vertex) *Edge {
	var minWeightEdge *Edge
	for _, vertex := range vertices {
		for _, edge := range vertex.Edges {
			if !edge.ToVertex.isVisited {
				if minWeightEdge == nil || minWeightEdge.Weight > edge.Weight {
					minWeightEdge = edge
				}
			}
		}
	}
	return minWeightEdge
}

func (graph *Graph) KruskalMinimumSpanningTree() {
	treeEdges := []*Edge{}
	for _, vertex := range graph.Vertices {
		for _, edge := range vertex.Edges {
			treeEdges = append(treeEdges, edge)
		}
	}
	treeCount := len(graph.Vertices)
	for treeCount > 1 {
		minWeightUnUsedEdge := getMinWeightUnUsedEdgeInEdges(treeEdges)
		if minWeightUnUsedEdge != nil {
			oppositeEdge := getOppositeEdgeInEdges(treeEdges, minWeightUnUsedEdge)
			if !graph.hasPathBetweenVertices(minWeightUnUsedEdge.FromVertex, minWeightUnUsedEdge.ToVertex) {
				minWeightUnUsedEdge.isUsed = true
				oppositeEdge.isUsed = true
				treeCount--
			} else {
				treeEdges = removeEdgeInEdges(treeEdges, minWeightUnUsedEdge)
				treeEdges = removeEdgeInEdges(treeEdges, oppositeEdge)
			}
		}
	}
	for _, edge := range treeEdges {
		if edge.isUsed {
			fmt.Printf("%s->%s(%d)\n", edge.FromVertex.Label, edge.ToVertex.Label, edge.Weight)
			opEdge := getOppositeEdgeInEdges(treeEdges, edge)
			if opEdge != nil && opEdge.isUsed {
				opEdge.isUsed = false
			}
		}
	}
	graph.clearEdgesUseHistory()
}

func getMinWeightUnUsedEdgeInEdges(edges []*Edge) *Edge {
	var minWeightUnUsedEdge *Edge
	for _, edge := range edges {
		if !edge.isUsed {
			if minWeightUnUsedEdge == nil || minWeightUnUsedEdge.Weight > edge.Weight {
				minWeightUnUsedEdge = edge
			}	
		}
	}
	return minWeightUnUsedEdge
}

func getOppositeEdgeInEdges(edges []*Edge, edge *Edge) *Edge {
	for _, e := range edges {
		if e.FromVertex == edge.ToVertex && e.ToVertex == edge.FromVertex && e.Weight == edge.Weight {
			return e
		}
	}
	return nil
}

func (graph *Graph) hasPathBetweenVertices(v1 *Vertex, v2 *Vertex) bool {
	queue := &queue.LinkedQueue{}
	v1.isVisited = true
	for _, edge := range v1.Edges {
		if edge.isUsed {
			queue.Add(edge.ToVertex)
		}
	}
	for queue.Size() > 0 {
		vertex := convertToVertex(queue.Peek())
		if vertex == v2 {
			return true
		} else {
			for _, e := range vertex.Edges {
				if e.isUsed && !e.ToVertex.isVisited {
					queue.Add(e.ToVertex)
				}
			}
		}
		vertex.isVisited = true
		queue.Remove()
	}
	graph.clearVerticesVisitHistory()
	return false
}

func removeEdgeInEdges(edges []*Edge, e *Edge) []*Edge {
	var es []*Edge
	for _, x := range edges {
		if x != e {
			es = append(es, x)
		}
	}
	return es
}

func (graph *Graph) DijkstraShortestPath(startVertex *Vertex, endVertex *Vertex) {
	distanceMap := make(map[string]int)
	prevVertexMap := make(map[string]*Vertex)
	for _, v := range graph.Vertices {
		distanceMap[v.Label] = MAX_INT
		prevVertexMap[v.Label] = nil
	}
	distanceMap[startVertex.Label] = 0
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		minDistanceVertex := graph.getMinDistanceVertex(distanceMap)
		if minDistanceVertex == endVertex {
			break
		}
		if distanceMap[minDistanceVertex.Label] == MAX_INT {
			break
		}
		for _, edge := range minDistanceVertex.Edges {
			toVertex := edge.ToVertex
			distance := distanceMap[minDistanceVertex.Label] + edge.Weight
			if distance < distanceMap[toVertex.Label] {
				distanceMap[toVertex.Label] = distance
				prevVertexMap[toVertex.Label] = minDistanceVertex
			}
		}
		minDistanceVertex.isVisited = true
	}
	graph.clearVerticesVisitHistory()
	for label, vertex := range prevVertexMap {
		if vertex == nil {
			delete(prevVertexMap, label)
		} else {
			if !canGoToStart(vertex, startVertex, prevVertexMap) {
				delete(prevVertexMap, label)
			}
			if !canGoToEnd(graph.getVertexByLabel(label), endVertex, prevVertexMap) {
				delete(prevVertexMap, label)
			}
		}
	}
	for label, vertex := range prevVertexMap {
		fmt.Printf("%s->%s(%d)\n", vertex.Label, label, getWeightByLabelAndPrevVertex(label, vertex))
	}
}

func (graph *Graph) getMinDistanceVertex(distanceMap map[string]int) *Vertex {
	distance := -1
	index := 0
	for i, v := range graph.Vertices {
		if !v.isVisited {
			if distance == -1 || distance > distanceMap[v.Label] {
				distance = distanceMap[v.Label]
				index = i
			}
		}
	}
	return graph.Vertices[index]
}

func canGoToStart(v *Vertex, startV *Vertex, prevVertexMap map[string]*Vertex) bool {
	if v == startV {
		return true
	}
	prevV := prevVertexMap[v.Label]
	for prevV != nil {
		if prevV == startV {
			return true
		} else {
			prevV = prevVertexMap[prevV.Label]
		}
	}
	return false
}

func canGoToEnd(v *Vertex, endV *Vertex, prevVertexMap map[string]*Vertex) bool {
	if v == endV {
		return true
	}
	prevV := prevVertexMap[endV.Label]
	for prevV != nil {
		if prevV == v {
			return true
		} else {
			prevV = prevVertexMap[prevV.Label]
		}
	}
	return false
}

func (graph *Graph) getVertexByLabel(label string) *Vertex {
	for _, v := range graph.Vertices {
		if v.Label == label {
			return v
		}
	}
	return nil
}

func getWeightByLabelAndPrevVertex(label string, prevVertex *Vertex) int {
	for _, edge := range prevVertex.Edges {
		if edge.ToVertex.Label == label {
			return edge.Weight
		}
	}
	return -1
}

func (graph *Graph) TopologicalSort() {
	//TODO
}

func (graph *Graph) getZeroInDegreeVertex() *Vertex {
	//TODO
	return nil
}

func (graph *Graph) getVisitedVertices() []*Vertex {
	vertices := []*Vertex{}
	for _, vertex := range graph.Vertices {
		if vertex.isVisited {
			vertices = append(vertices, vertex)
		}
	}
	return vertices
}

func (graph *Graph) clearVerticesVisitHistory() {
	for _, v := range graph.Vertices {
		v.isVisited = false
	}
}

func (graph *Graph) clearEdgesUseHistory() {
	for _, v := range graph.Vertices {
		for _, e := range v.Edges {
			e.isUsed = false
		}
	}
}

func convertToVertex(x interface{}) *Vertex {
	if v, ok := x.(*Vertex); ok {
		return v
	} else {
		panic("Type convertion exception.")
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
	graph.Vertices = []*Vertex{Va, Vb, Vc, Vd, Ve, Vf, Vg}
	
	fmt.Println("BreadFirstTraverse:")
	graph.BreadFirstTraverse(Va)

	fmt.Println("\nDepthFirstTraverse:")
	graph.DepthFirstTraverse(Va)

	fmt.Println("\nPrimMinimumSpanningTree:")
	graph.PrimMinimumSpanningTree(Va)

	fmt.Println("KruskalMinimumSpanningTree:")
	graph.KruskalMinimumSpanningTree()

	fmt.Println("DijkstraShortestPath from A to G:")
	graph.DijkstraShortestPath(Va, Vg)

	fmt.Println("TopologicalSort:")
	graph.TopologicalSort()
}
