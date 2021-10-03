package bloom

//There are two types of bloom filters that I"m primarily familar with. The first is one that uses a byte slice as the filter, and
//binary operations to determine whether the bytes of an entry are filled on the filter. This seems really complex and outside of
//the current realm of my purview. The second method is basically just a hash table that stores boolean values instead of entries.
//That seems way easier and more doable, so I'm just going to use that

//Filter is a struct containing a byte slice, which we will manipulate to work with bits
type Filter struct {
	Filter [1052]bool
}

//Put adds an entry to the byte slice of filter f by performing a bitwise & operation
func (f *Filter) Put(e []byte) {
	//hash using go's crypto library
	//res := crypto.Hash(e)
	//f.Filter[res] = true
}

//Contains checks to see if an entry is in the hash map
func (f *Filter) Contains(e []byte) bool {
	//return f.Filter[crypto.Hash(e)]
	return false
}
