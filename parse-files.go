// This files provides utilities for parsing the text files that contain the YouTube data
// it will use simple regex to split the values and then ultimately join them together into nodes
// then it will take all of the nodes and create a graph

package main

import (
	"os"
	"bufio"
	"strings"
	"errors"
)

func createGraphFromFile(fileName string) graph {
	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)


	g := NewGraph()

	for scanner.Scan() {
		err, n := createNodeFromText(scanner.Text())
		if err != nil {
			continue
		}
		g.InsertNode(n)

	}

	return g
}

func createNodeFromText(line string) (error, node) {
	values := strings.Split(line, "	")


	if len(values) < 7 {
		e := errors.New("incomplete video data")
		return e, node{}
	}

	id := values[0]
	author := values[1]
	age := values[2]
	category := values[3]
	length := values[4]
	rate := values[5]
	ratings := values[6]
	comments := values[7]

	// everything past 8 is related videos
	related := values[8:]

	edges := make([]edge, 0)
	for i := range related {
		e := edge{fromNode:id, toNode:related[i]}
		edges = append(edges, e)
	}
	n := node{
			id:id,
			author: author,
			age: age,
			category: category,
			length: length,
			rate: rate,
			ratings: ratings,
			comments: comments,
			edges:edges,
	}

	return nil, n
}

