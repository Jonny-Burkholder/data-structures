package graph

import "errors"

var errMustBePositive = errors.New("Index value must be positive")
var errNodeExists = errors.New("Node already exists")
var errConnectionNodeNotFound = errors.New("Unable to make connection: node not found")
var errNodeOutOfRange = errors.New("Node out of range")
var errCannotDeleteNil = errors.New("Cannot delte: node does not exist") //totally unnecessary, but I like to be explicit with these things
var errCannotConnectSelf = errors.New("Node cannot be connected to itself")

//Node is an instance or entry in the graph. It holds data
//and an adjacency list. The indeces represent the nodes held
//in the graph struct, with the float value being the weight
type node struct {
	data  []byte
	edges []float64 //probably will only need an int in most cases. I should learn generics
}

//graph basically just holds the first node in the graph
type graph struct {
	nodes []*node
}

//newGraph returns a pointer to a new graph
func newGraph() *graph {
	return &graph{
		nodes: make([]*node, 0),
	}
}

//newNode appends a node to the end of the graph
func (g *graph) newNode(data []byte) *node {
	n := &node{
		data:  data,
		edges: make([]float64, len(g.nodes)),
	}
	g.nodes = append(g.nodes, n)
	return n
}

//newNodeAt places a new node in a specific location
func (g *graph) newNodeAt(a int, data []byte) (*node, error) {

	if a < 0 {
		return nil, errMustBePositive
	}

	if a >= len(g.nodes) {
		pad := make([]*node, a-len(g.nodes))
		g.nodes = append(g.nodes, pad...)
	}

	if g.nodes[a] != nil {
		return nil, errNodeExists
	}

	n := &node{
		data:  data,
		edges: make([]float64, len(g.nodes)),
	}

	g.nodes[a] = n

	return n, nil

}

//delNode removes a node from the graph
func (g *graph) delNode(a int) ([]byte, error) {
	if a < 0 {
		return nil, errMustBePositive
	}
	if a >= len(g.nodes) {
		return nil, errNodeOutOfRange
	}

	if g.nodes[a] == nil {
		return nil, errCannotDeleteNil
	}

	d := g.nodes[a].data

	g.nodes[a] = nil //should I specifically ask/force go to garbage collect here? I don't actually know how to do that

	return d, nil

}

//newConnection takes a node and adds a connection to its adjecency list. Magic!
func (g *graph) newConnection(a, b int, weight ...float64) error {}
