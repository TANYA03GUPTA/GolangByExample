package main

import "fmt"

func isBipartite(graph [][]int) bool {
	nodeMap := make(map[int][]int)

	numNodes := len(graph)

	if numNodes == 1 {
		return true
	}
	for i := 0; i < numNodes; i++ {
		nodes := graph[i]
		for j := 0; j < len(nodes); j++ {
			nodeMap[i] = append(nodeMap[i], nodes[j])
		}
	}

	color := make(map[int]int)

	for i := 0; i < numNodes; i++ {
		if color[i] == 0 {
			color[i] = 1
			isBiPartite := visit(i, nodeMap, &color)
			if !isBiPartite {
				return false
			}
		}
	}

	return true
}

func visit(source int, nodeMap map[int][]int, color *map[int]int) bool {

	for _, neighbour := range nodeMap[source] {
		if (*color)[neighbour] == 0 {
			if (*color)[source] == 1 {
				(*color)[neighbour] = 2
			} else {
				(*color)[neighbour] = 1
			}
			isBipartite := visit(neighbour, nodeMap, color)
			if !isBipartite {
				return false
			}
		} else {
			if (*color)[source] == (*color)[neighbour] {
				return false
			} 
		}
	}

	return true
}

func main() {
	output := isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}})
	fmt.Println(output)

	output = isBipartite([][]int{{1, 4}, {0, 2}, {1, 3}, {2, 4}, {0, 3}})
	fmt.Println(output)

}