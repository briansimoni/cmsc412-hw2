package main

import (
	"fmt"
)

// the main function will create a graph object
// read the graph.txt file
// and then output the unnormalized weighted degree centrality
// to wdegree.txt
func main() {

	g := createGraphFromFile("youtube/testg.txt")
	fmt.Println(g)


	//fileString := ""
	//for i := range g.nodes {
	//	fileString += strconv.FormatFloat(g.nodes[i].closenessCentrality, 'f', -1, 64) + "\r\n"
	//}
	//
	//f, err := os.Create("closeness.txt")
	//check(err)
	//f.Write([]byte(fileString))


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