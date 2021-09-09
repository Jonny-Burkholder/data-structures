package bloom

//Filter will hold a slice of bytes
type Filter struct {
	Filter []byte
}

//Search checks to see if a byte slice is contained in the filter. This is done with an and operator
//Search returns a non-nil error if the byte slice argument is not found to be matching
func (f *Filter) Search(b []byte) error {
	//check to see if the bits in the byte slice are contained in the filter
	//If not, return an error
	return nil
}

//Put adds an entry to the filter, using some kind of bitwise operator. I forget, gotta double-check
func (f *Filter) Put(b []byte) {
	//Use the fancy binary filter thing. It's late and I can't remember which operator to use
}
