package graph

import (
	"testing"

	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
)

var graphFromA Graph
var graphFromB Graph

func TestMain(m *testing.M) {
	utl.SetupTests(m, createGraph)
}

func createGraph() {
	/*
		graph
		  e __  d __ a __ b __ f __ g __ i
				\	/				|
				  c					h -- k

	*/
	vertexK := Vertex{value: "K"}
	vertexI := Vertex{value: "I"}
	vertexH := Vertex{value: "H", edges: []Edge{
		{end: &vertexK},
	}}
	vertexG := Vertex{value: "G", edges: []Edge{
		{end: &vertexH},
		{end: &vertexI},
	}}
	vertexF := Vertex{value: "F", edges: []Edge{
		{end: &vertexG},
	}}
	vertexE := Vertex{value: "E"}
	vertexD := Vertex{value: "D", edges: []Edge{
		{end: &vertexE},
	}}
	vertexC := Vertex{value: "C"}
	vertexB := Vertex{value: "B"}
	vertexA := Vertex{value: "A", edges: []Edge{
		{end: &vertexB},
		{end: &vertexC},
		{end: &vertexD},
	}}
	vertexC.edges = []Edge{
		{end: &vertexA},
		{end: &vertexD},
	}
	vertexB.edges = []Edge{
		{end: &vertexA},
		{end: &vertexF},
	}
	graphFromA.start = &vertexA
	graphFromB.start = &vertexB
}

func TestBFS(t *testing.T) {
	bfsTraversalFromA := graphFromA.TraverseBreadthFirst()
	if bfsTraversalFromA != "ABCDFEGHIK" {
		t.Errorf("Breadth First Traversal should be in the order -- ABCDFEGHIK, got  -- %v \n", bfsTraversalFromA)
	}

	bfsTraversalFromB := graphFromB.TraverseBreadthFirst()
	if bfsTraversalFromB != "BAFCDGEHIK" {
		t.Errorf("Breadth First Traversal should be in the order -- BAFCDGEHIK, got  -- %v \n", bfsTraversalFromB)
	}
}

func TestDFS(t *testing.T) {
	dfsTraversalFromA := graphFromA.TraverseDepthFirst()
	if dfsTraversalFromA != "ABFGHKICDE" {
		t.Errorf("Depth First Traversal should be in the order -- ABFGHKICDE, got -- %v \n", dfsTraversalFromA)
	}

	dfsTraversalFromB := graphFromB.TraverseDepthFirst()
	if dfsTraversalFromB != "BACDEFGHKI" {
		t.Errorf("Depth First Traversal should be in the order -- BACDEFGHKI, got -- %v \n", dfsTraversalFromB)
	}
}
