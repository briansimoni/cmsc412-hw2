package main

import "reflect"

type node struct {
	id       string
	author	string
	age      string
	category string
	length   string
	views	string
	rate     string
	ratings  string
	comments string
	edges    []edge

	// using distance and parent for breadth first search algorithm
	distance int
	parent   *node

	closenessCentrality float64
}

func (n node) WeightedDegreeCentrality() float64 {

	sum := 0.0

	for i := range n.edges {
		sum += n.edges[i].Weight
	}

	return sum
}

func (n *node) AddEdge(e edge) {
	for i := range n.edges {
		if reflect.DeepEqual(e, n.edges[i]) {
			return
		}
	}

	n.edges = append(n.edges, e)
}
