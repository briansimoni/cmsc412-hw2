// This files provides utilities for parsing the text files that contain the YouTube data
// it will use simple regex to split the values and then ultimately join them together into nodes
// then it will take all of the nodes and create a graph

package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func createGraphFromFile(fileName string) graph {
	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)


	g := NewGraph()

	for scanner.Scan() {
		n := createNodeFromText(scanner.Text())
		g.InsertNode(n)

	}

	return g
}

func createNodeFromText(line string) node {
	values := strings.Split(line, "	")
	if len(values) < 10 {
		fmt.Println(len(values))
	}
	id := values[0]
	related := values[len(values)-1]

	// from and to
	e := edge{fromNode:id, toNode:related}
	edges := []edge{e}
	n := node{
			id:id,
			edges:edges,
	}

	return n
}

