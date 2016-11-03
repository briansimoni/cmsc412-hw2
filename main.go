package main

import (
	"fmt"
)

// the main function will create a graph object
// read the graph.txt file
// and then output the unnormalized weighted degree centrality
// to wdegree.txt
func main() {

	g := createGraphFromFile("youtube/1.txt")
	//g.closenessCentrality()
	//
	//
	//fileString := ""
	//for k, _ := range g.nodes {
	//	fileString += strconv.FormatFloat(g.nodes[k].closenessCentrality, 'f', -1, 64) + "\r\n"
	//}
	//
	//f, err := os.Create("closeness.txt")
	//check(err)
	//f.Write([]byte(fileString))
	//
	//for _, node := range g.nodes {
	//	fmt.Println(node.category)
	//}

	//result := g.topCategoryByVideoCount()
	//fmt.Println(result)

	result2 := g.averageCategoryRate()
	for k, v := range result2 {
		fmt.Println(k,v)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// parse the line of text of the file, and extract an edge from a line
//func getEdge(line string) edge {
//
//	e := strings.Split(line, " ")
//	fromNode, _ := strconv.Atoi(e[0])
//	toNode, _ := strconv.Atoi(e[1])
//	var weight float64 = 0.0
//
//	edge := edge{fromNode, toNode, weight}
//	return edge
//
//}

// parse the first line and return the number of nodes in the graph
//func firstLine(line string) int {
//	l := strings.Split(line, " ")
//	n, _ := strconv.Atoi(l[0])
//	return n
//}
//
//// flips the to and from node
//func undirectedEdge(e edge) edge {
//	fromNode := e.toNode
//	toNode := e.fromNode
//	return edge{fromNode, toNode, e.Weight}
//}
