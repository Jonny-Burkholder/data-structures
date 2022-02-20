package graph

import "errors"

var errMustBePositive = errors.New("Index value must be positive")
var errVectorExists = errors.New("Vector already exists")
var errConnectionVectorNotFound = errors.New("Unable to make connection: vector not found")
var errVectorOutOfRange = errors.New("Vector out of range")
var errCannotDeleteNil = errors.New("Cannot delte: vector does not exist") //totally unnecessary, but I like to be explicit with these things
var errCannotConnectSelf = errors.New("Vector cannot be connected to itself")

type graph struct {
	Data    [][]byte
	Vectors [][]float64
}

//newGraph returns a new graph function
func newGraph() *graph {
	return &graph{
		Data:    make([][]byte, 0),
		Vectors: make([][]float64, 0),
	}
}

//addVector appends an empty vector to the end of the graph
func (g graph) addVector(data []byte) {
	g.Data = append(g.Data, data)
	g.Vectors = append(g.Vectors, make([]float64, 0))
}

//addvectorAt takes a position argument and creates a vector at the given
//index. If the index is out of range, the function simply appends to
//the slice to pad to the given index
func (g graph) addVectorAt(a int, data []byte) error {
	if a < 0 {
		return errMustBePositive
	}
	if a >= len(g.Data) {
		dataPad := make([][]byte, a-len(g.Data))
		g.Data = append(g.Data, dataPad...)
		pad := make([][]float64, a-len(g.Data))
		g.Vectors = append(g.Vectors, pad...)
	}

	if g.Vectors[a] != nil {
		return errVectorExists
	}

	g.Data[a] = data
	g.Vectors[a] = make([]float64, 0)

	return nil
}

//delVector removes all connections at an index and resets the index to nil
func (g graph) delVecor(a int) error {
	if a < 0 {
		return errMustBePositive
	}
	if a >= len(g.Vectors) || g.Vectors[a] == nil {
		return errCannotDeleteNil //should probably have an "index out of range" error or something
	}

	g.Data[a] = nil
	g.Vectors[a] = nil

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
	if a >= len(g.Vectors) || b >= len(g.Vectors) {
		return errConnectionVectorNotFound
	}

	if g.Vectors[a] == nil {
		return errConnectionVectorNotFound
	}

	if b >= len(g.Vectors[a]) {
		pad := make([]float64, b-len(g.Vectors[a]))
		g.Vectors[a] = append(g.Vectors[a], pad...)
	}

	if len(weight) > 0 {
		g.Vectors[a][b] = weight[0]
	} else {
		g.Vectors[a][b] = 1
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

	if a >= len(g.Vectors) || b >= len(g.Vectors) {
		return errConnectionVectorNotFound
	}

	if b >= len(g.Vectors[a]) {
		return errVectorOutOfRange
	}

	g[a][b] = 0

	return nil

}

//So, yeah. This searches for stuff. Starts at index 0 probably
func (g graph) depthSearch(data []byte) (int, error) {
	//make a slice of bools for all the vectorss. If it's
	//already been searched, it will be flipped to true
	searched := make([]bool, len(g.Vectors))
	//make a variable for the next vector to search, and keep
	//track of how many vectors we've searched, so we can stop
	//when we've searched them all... yeah? idk.
	next, count := 0, 0
	for {
		//look at the next vector to see if it contains the data
		//if not, pick the first connectio not already searched
		//and perform depth search on it. Which is fine and all,
		//but we need to know which vectos *that* function searches.
		//Might have to keep track of the searched nodes on the struct
		//level
	}
}
