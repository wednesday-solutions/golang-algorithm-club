package graph

import (
	hash "github.com/wednesday-solutions/golang-algorithm-club/pkg/datastructures/hashtable"
	queue "github.com/wednesday-solutions/golang-algorithm-club/pkg/datastructures/queue"
)

/*

Using Adjacency list (Object oriented model) - https://en.wikipedia.org/wiki/Adjacency_list

*/

type (

	// Graph A Graph, need a start vertex pointer for the methods
	Graph struct {
		start *Vertex
	}

	// Interface BFS and DFS traversal methods on the graph
	Interface interface {
		TraverseBreadthFirst() string
		TraverseDepthFirst() string
	}

	// Vertex represents each vertex in the graph
	Vertex struct {
		value string
		edges []Edge
	}

	// Edge Only one end of the edge is stored
	Edge struct {
		end *Vertex
	}
)

/*
* 1. Get the hashed value of the current vertex value.
* 2. If not visited, append to result, add the edge vertices to the visit queue
* 3. Insert each edge to the hashmap with value = false and add to the visit queue
* 4. dequeue and recurse.
 */

// TraverseBreadthFirst Traverse the graph breadth first from the starting vertex
func (graph *Graph) TraverseBreadthFirst() string {
	hash.ResetTable()
	return graph.start.traverseBreadthFirst("", queue.NewQueue())
}

func (ver *Vertex) traverseBreadthFirst(result string, VisitQueue queue.Queue) string {
	hashValue := hash.GetValueFromBucket(ver.value)
	if hashValue == false || hashValue != nil {
		hash.Insert(ver.value, true)
		result = result + ver.value
		for _, vertexEdge := range ver.edges {
			vertexEdgeValue := vertexEdge.end.value
			hashValueOfVertexEdge := hash.GetValueFromBucket(vertexEdgeValue)
			if hashValueOfVertexEdge == "" {
				hash.Insert(vertexEdgeValue, false)
				VisitQueue.Enqueue(vertexEdge.end)
			}
		}
	}
	dequedValue := VisitQueue.Dequeue()
	if nextVertex, ok := dequedValue.(*Vertex); ok {
		return nextVertex.traverseBreadthFirst(result, VisitQueue)
	}
	return result
}

/*
* 1. Get the hashed value of the current vertex value, append the result if it has not been visited.
* 2. Go through edges and recurse
 */

// TraverseDepthFirst Traverse the graph Depth first from the starting vertex
func (graph *Graph) TraverseDepthFirst() string {
	hash.ResetTable()
	return graph.start.traverseDepthFirst("")
}

func (ver *Vertex) traverseDepthFirst(result string) string {
	hashValue := hash.GetValueFromBucket(ver.value)
	if hashValue == false || hashValue != nil {
		hash.Insert(ver.value, true)
		result = result + ver.value
		for _, vertexEdge := range ver.edges {
			vertexEdgeValue := vertexEdge.end.value
			hashValueOfVertexEdge := hash.GetValueFromBucket(vertexEdgeValue)
			if hashValueOfVertexEdge == "" {
				hash.Insert(vertexEdgeValue, false)
				result = vertexEdge.end.traverseDepthFirst(result)
			}
		}
	}
	return result
}
