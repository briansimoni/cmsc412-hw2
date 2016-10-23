package main

import (
	"math"
)

const infinity = math.MaxInt32

// a graph is simply a collection of nodes and edges (the node struct contains edge data)
type graph struct {
	nodes []node
}

// function to see if the node is already in the graph
func (g graph) IsInGraph(n int) bool {
	for i := range g.nodes {
		if g.nodes[i].id == n {
			return true
		}
	}

	return false
}


func (g graph) InsertNode(n node) {
	if !g.IsInGraph(n.id) {
		g.nodes = append(g.nodes, n)
	}
}

func (g *graph) initialize () {
	for i := range g.nodes {
		g.nodes[i].distance = infinity
		g.nodes[i].parent = nil
	}
}


// the breadth first search function sets the shortest path distance to all other nodes in g
func (g *graph) breadthFirstSearch(root *node) {

	g.initialize()

	// create empty queue
	queue := make([]*node, 0)

	root.distance = 0
	// enqueue root (push)
	queue = append(queue, root)

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for i := range current.edges {
			// remember to subtract 1 to account for my 0 based index of nodes
			neighbor := current.edges[i].toNode - 1
			if g.nodes[neighbor].distance == infinity {
				g.nodes[neighbor].distance = current.distance + 1
				g.nodes[neighbor].parent = current
				queue = append(queue, &g.nodes[neighbor])
			}
		}
	}
}


// calculate the normalized closeness centrality for all nodes in g
func (g *graph) closenessCentrality() {
	n := float64(len(g.nodes))
	for i := range g.nodes {
		g.breadthFirstSearch(&g.nodes[i])

		var sum float64 = 0.0
		for j := range g.nodes {
			sum += float64(g.nodes[j].distance)
		}

		sum = sum / (n - 1)
		g.nodes[i].closenessCentrality = math.Pow(sum, -1)

	}
}

// create a new graph with a specified number of nodes
func NewGraph(size int) graph {
	nodes := make([]node, size)
	return graph{nodes: nodes}
}


