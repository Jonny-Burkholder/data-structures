package binary

//binary basic is a simple on-disk data structure using binary and little endian encoding
//This package will seek to implement an on-disk equivalent to a double-linked list

//dntry is a struct containing a header, data, and the indexes of the previous and next entries.
//If previous or next does not apply, the index will be nil
type entry struct{
	Header *header
	Data []byte
	Previous []byte
	Next []byte
}

//Header contains the index and the length of the entry
type header struct{
	Len uint

}

type binaryFile struct{
	
}