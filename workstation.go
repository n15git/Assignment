package main

import (
	"fmt"
)

type Connection struct {
	a, b int
}

func findParent(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return findParent(parent, parent[i])
}

func unionSets(parent []int, x, y int) {
	xset := findParent(parent, x)
	yset := findParent(parent, y)
	parent[xset] = yset
}

func makeConnected(n int, connections []Connection, numConnections int) int {
	if numConnections < n-1 {
		return -1 
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1 
	}

	extraConnections := 0 
	for i := 0; i < numConnections; i++ {
		x := findParent(parent, connections[i].a)
		y := findParent(parent, connections[i].b)
		if x != y { 
			unionSets(parent, x, y) 
		} else {
			extraConnections++ 
		}
	}

	distinctParents := make(map[int]bool)
	for i := 0; i < n; i++ {
		distinctParents[findParent(parent, i)] = true 
	}

	if len(distinctParents)-1 < extraConnections {
		return -1
	}
	return len(distinctParents) - 1
}

func main() {
	var n, numConnections int
	fmt.Print("Enter the number of workstations: ")
	fmt.Scan(&n)

	fmt.Print("Enter the number of connections: ")
	fmt.Scan(&numConnections)

	connections := make([]Connection, numConnections)
	fmt.Println("Enter the connections (format: a b): ")
	for i := 0; i < numConnections; i++ {
		fmt.Scan(&connections[i].a, &connections[i].b)
	}

	fmt.Println("Minimum number of times to reconnect all workstations:", makeConnected(n, connections, numConnections))
}
