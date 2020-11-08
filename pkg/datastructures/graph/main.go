package main

import (
	"fmt"

	linkedlist "github.com/wednesday-solutions/golang-algorithm-club/pkg/datastructures/linkedlist"
)

/*
 	   2
	/  |  \
  /	   |    \
1 ---- 5 ---- 3
  \	   |  	/
	\  |  /
	   4
*/

var adjacencyList []linkedlist.LinkedList

func createAdjacencyList(adjacentItems [][]string) {
	for _, v := range adjacentItems {
		adjacencyList = append(adjacencyList, *linkedlist.CreateLinkedListFromArray(v))
	}
}

func printAdjacency(edgeLink linkedlist.Node) {
	fmt.Printf(edgeLink.Value)
	if edgeLink.Next != nil {
		fmt.Print(" --> ")
		printAdjacency(*edgeLink.Next)
	} else {
		fmt.Println()
	}
}

func edgeValues(list linkedlist.Node, result []string) []string {
	result = append(result, list.Value)
	if list.Next != nil {
		return edgeValues(*list.Next, result)
	}
	return result
}

func getAdjacents(val string, input []string) []string {
	l := len(input)
	if l < 2 {
		return []string{""}
	}
	result := []string{}
	for i, v := range input {
		if v == val && l != i {
			switch {
			case i == 0:
				result = append(result, input[i+1])
			case i == l-1:
				result = append(result, input[i-1])
			default:
				result = append(result, input[i-1], input[i+1])
			}
		}
	}
	return result
}

func traverseBreadthFirst(startAt string, result string) {
	// visited := []string{}
	allEdgeValues := make([][]string, len(adjacencyList))
	for i, v := range adjacencyList {
		allEdgeValues[i] = edgeValues(*v.Head, allEdgeValues[i])
	}
	fmt.Println(allEdgeValues)
	adjacents := getAdjacents("3", allEdgeValues[0])
	fmt.Println(adjacents)
}

func main() {
	edge1 := []string{"1", "5", "3"}
	edge2 := []string{"1", "2", "3"}
	edge3 := []string{"1", "4", "3"}
	edge4 := []string{"1", "2", "5", "4"}
	edge5 := []string{"3", "2", "5", "4"}
	edge6 := []string{"2", "1", "5", "3"}
	edge7 := []string{"4", "1", "5", "3"}

	input := [][]string{edge1, edge2, edge3, edge4, edge5, edge6, edge7}
	createAdjacencyList(input)
	printAdjacency(*adjacencyList[6].Head)
	result := ""
	traverseBreadthFirst("1", result)
}
