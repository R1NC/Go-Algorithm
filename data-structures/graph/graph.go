package graph

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
	inDegreeMap map[string]int
}

func (graph *Graph) BreadFirstTraverse(startVertex *Vertex) {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	fmt.Printf("%s ", startVertex.Label)
	startVertex.isVisited = true
	queue := &queue.LinkedQueue{}
	queue.Add(startVertex)
	for queue.Size() > 0 {// Visit the nearest vertices that haven't been visited.
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
	for stack.Size() > 0 {// Visit the the vertices by edges that hasn't been visited, until the path ends.
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
			} else {// There's a ring, remove the edge hand its opposite edge.
				treeEdges = removeEdgeInEdges(treeEdges, minWeightUnUsedEdge)
				treeEdges = removeEdgeInEdges(treeEdges, oppositeEdge)
			}
		}
	}
	for _, edge := range treeEdges {
		if edge.isUsed {
			fmt.Printf("%s->%s(%d)\n", edge.FromVertex.Label, edge.ToVertex.Label, edge.Weight)
			opEdge := getOppositeEdgeInEdges(treeEdges, edge)
			if opEdge != nil && opEdge.isUsed {// Filter opposite edges.
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
		nearestVertex := graph.getNearestVertex(startVertex, distanceMap)
		if nearestVertex == endVertex {//Reached EndVertex.
			break
		}
		if distanceMap[nearestVertex.Label] == MAX_INT {//There's no path between two vertices.
			break
		}
		for _, edge := range nearestVertex.Edges {// Update distance map.
			toVertex := edge.ToVertex
			distance := distanceMap[nearestVertex.Label] + edge.Weight
			if distance < distanceMap[toVertex.Label] {
				distanceMap[toVertex.Label] = distance
				prevVertexMap[toVertex.Label] = nearestVertex
			}
		}
		nearestVertex.isVisited = true
	}
	graph.clearVerticesVisitHistory()
	for label, vertex := range prevVertexMap {
		if vertex == nil {// Filter StartVertex.
			delete(prevVertexMap, label)
		} else {// Filter the vertices that can't reach StartVertex and EndVertex.
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

func (graph *Graph) getNearestVertex(startVertex *Vertex, distanceMap map[string]int) *Vertex {
	distance := -1
	index := -1
	for i, v := range graph.Vertices {
		if !v.isVisited {
			if distance == -1 || distance > distanceMap[v.Label] {
				distance = distanceMap[v.Label]
				index = i
			}
		}
	}
	if index == -1 {// First scanning, return StartVertex.
		return startVertex
	} else {
		return graph.Vertices[index]
	}
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
	graph.inDegreeMap = make(map[string]int)
	for _, v := range graph.Vertices {
		graph.inDegreeMap[v.Label] = 0
	}
	for _, v := range graph.Vertices {
		for _, e := range v.Edges {
			graph.inDegreeMap[e.ToVertex.Label]++
		}
	}
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		topVertices := graph.getZeroInDegreeVertices()
		for _, v := range topVertices {// Visit the zero-in-degree-vertex, and decrease the next vertices' in-degree.
			fmt.Printf("%s ", v.Label)
			v.isVisited = true
			for _, edge := range v.Edges {
				graph.inDegreeMap[edge.ToVertex.Label]--
			}
		}
		fmt.Println()
	}
	graph.clearVerticesVisitHistory()
}

func (graph *Graph) getZeroInDegreeVertices() []*Vertex {
	vertices := []*Vertex{}
	for _, v := range graph.Vertices {
		if graph.inDegreeMap[v.Label] == 0 && !v.isVisited {
			vertices = append(vertices, v)
		}
	}
	return vertices
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