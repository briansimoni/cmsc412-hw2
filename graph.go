package main

import (
	"math"
	"strconv"
	"fmt"
)

const infinity = math.MaxInt32

// a graph is simply a collection of nodes and edges (the node struct contains edge data)
type graph struct {
	nodes map[string]node
}

func NewGraph() graph {
	nodes := make(map[string]node)
	return graph{nodes: nodes}
}

// function to see if the node is already in the graph
func (g *graph) IsInGraph(id string) bool {
	if _, ok := g.nodes[id]; ok {
		return true
	}
	return false
}

func (g *graph) InsertNode(n node) {
	if !g.IsInGraph(n.id) {
		g.nodes[n.id] = n
	}
}

func (g *graph) PrintGraphInformation() {
	for id, _ := range g.nodes {
		fmt.Println("Node", id)
		for i := 0; i < len(g.nodes[id].edges); i++ {
			fmt.Println(g.nodes[id].edges[i].toNode)
		}

	}
}

func (g *graph) initialize() {
	for id, node := range g.nodes {
		node.distance = infinity
		node.parent = nil
		g.nodes[id] = node
	}

}

// todo refactor breadthFirstSearch for maps
// the breadth first search function sets the shortest path distance to all other nodes in g
func (g *graph) breadthFirstSearch(root string) {

	g.initialize()

	// create empty queue
	queue := make([]string, 0)

	r := g.nodes[root]
	r.distance = 0
	g.nodes[root] = r
	// enqueue root (push)
	queue = append(queue, root)

	for len(queue) != 0 {
		current := g.nodes[queue[0]]
		queue = queue[1:]
		//for i := range current.edges {
		for i := 0; i < len(current.edges); i ++ {

			// neighbor is a node
			neighbor := g.nodes[current.edges[i].toNode]
			//fmt.Println("what", neighbor.edges[i].toNode)
			if g.nodes[neighbor.id].distance == infinity {

				neighbor.distance = current.distance + 1
				neighbor.parent = &current
				g.nodes[neighbor.id] = neighbor

				queue = append(queue, neighbor.id)
			}
		}
		//}
	}
}

// calculate the normalized closeness centrality for all nodes in g
func (g *graph) closenessCentrality() {
	totalNodes := float64(len(g.nodes))

	for nodeID, node := range g.nodes {
		g.breadthFirstSearch(nodeID)
		var sum float64 = 0.0
		for _, n := range g.nodes {
			sum += float64(n.distance)
		}

		sum = sum / (totalNodes - 1)
		node.closenessCentrality = math.Pow(sum, -1)

	}
}

func (g *graph) topCategoryByVideoCount() map[string]int {

	sums := make(map[string]int, 0)

	for _, node := range g.nodes {
		sums[node.category]++
		sums["total"]++
	}

	return sums
}


func (g *graph) averageCategoryRate() map[string]float64 {

	sums := make(map[string]int, 0)
	floatSums := make(map[string]float64, 0)

	for _, node := range g.nodes {
		sums[node.category]++

		f, _ := strconv.ParseFloat(node.ratings, 64)
		if f == 0 {
			sums[node.category]--
		}

		floatSums[node.category] += f
	}

	for k, _ := range sums {
		floatSums[k] = floatSums[k] / float64(sums[k])
	}

	return floatSums
}
