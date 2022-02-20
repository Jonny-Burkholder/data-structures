package graph

import "errors"

var errMustBePositive = errors.New("Index value must be positive")
var errVectorExists = errors.New("Vector already exists")
var errConnectionVectorNotFound = errors.New("Unable to make connection: vector not found")
var errVectorOutOfRange = errors.New("Vector out of range")
var errCannotDeleteNil = errors.New("Cannot delte: vector does not exist") //totally unnecessary, but I like to be explicit with these things
var errCannotConnectSelf = errors.New("Vector cannot be connected to itself")

type graph [][]float64

//newGraph returns a new graph function
func newGraph() graph {
	return make([][]float64, 0)
}

//addVector appends an empty vector to the end of the graph
func (g graph) addVector() {
	g = append(g, make([]float64, 0))
}

//addvectorAt takes a position argument and creates a vector at the given
//index. If the index is out of range, the function simply appends to
//the slice to pad to the given index
func (g graph) addVectorAt(a int) error {
	if a < 0 {
		return errMustBePositive
	}
	if a >= len(g) {
		g = append(g, make([]float64, a-len(g)))
	}

	if g[a] != nil {
		return errVectorExists
	}

	g[a] = make([]float64, 0)

	return nil
}

//delVector removes all connections at an index and resest the index to nil
func (g graph) delVecor(a int) error {
	if a < 0 {
		return errMustBePositive
	}
	if a >= len(g) || g[a] == nil {
		return errCannotDeleteNil //should probably have an "index out of range" error or something
	}

	g[a] = nil

	return nil

}

//Add connection takes a vector argument a, and adds to it a connection
//to vector b. If a weight is given, it adds that weight as the number
//representing the edge between a and b. Otherwise, the weight of the
//connection is simply 1. If b is out of range of vector a, the function
//appends to the vector to pad to the given index
func (g graph) addConnection(a, b int, weight ...float64) error {

	if a == b {
		return errCannotConnectSelf
	}

	if a < 0 || b < 0 {
		return errMustBePositive
	}
	if a >= len(g) || b >= len(g) {
		return errConnectionVectorNotFound
	}

	if g[a] == nil {
		return errConnectionVectorNotFound
	}

	if b >= len(g[a]) {
		pad := make([]float64, b-len(g[a]))
		g[a] = append(g[a], pad...)
	}

	if len(weight) > 0 {
		g[a][b] = weight[0]
	} else {
		g[a][b] = 1
	}

	return nil

}

//delConnection takes vector arguments a and b, and sets the connection
//from a to b to zero
func (g graph) delConnection(a, b int) error {
	if a == b {
		return errVectorOutOfRange
	}

	if a < 0 || b < 0 {
		return errMustBePositive
	}

	if a >= len(g) || b >= len(g) {
		return errConnectionVectorNotFound
	}

	if b >= len(g[a]) {
		return errVectorOutOfRange
	}

	g[a][b] = 0

	return nil

}
