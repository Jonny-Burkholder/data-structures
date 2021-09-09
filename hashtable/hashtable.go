package hashtable

import (
	"fmt"
	"runtime"
	"sync"
)

//Entry is what we put into the hash table
type Entry struct {
	Key      string //to double-check for collisions
	Data     []byte //maybe a bit jenky, but with slices it don't even matter. I mean it does but I'm not bothered
	Next     *Entry
	Previous *Entry
}

//String fulfills stringer for type Entry
func (e *Entry) String() string {
	if e == nil {
		return `Entry{"","",nil}`
	}
	ss := fmt.Sprintf("Entry{%q, %q, %p}", e.Key, e.Data, e.Next)
	return ss
}

//HashTable is where the data is stored and what we implement methods on. It will take a string as a
//key and a byte slice as data
type HashTable struct {
	//I'd love to have both len and cap here, but cap is basically infinite because of separate chaining
	Len  int
	Data []*Entry
	Mu   sync.Mutex
}

//NewHashTable returns an empty hash table of a given length
func NewHashTable(len int) *HashTable {
	return &HashTable{
		Len:  len,
		Data: make([]*Entry, len),
	}
}

//Destroy sets all entries to nil and garbage collects the hash table
func Destroy(ht *HashTable) {
	ht.Mu.Lock()
	defer ht.Mu.Unlock()

	for i := range ht.Data {
		ht.Data[i] = nil // free entry pointers
	}
	ht = nil

	runtime.GC() // force collection
}

//String is a stringer function for printing the table to st.Out
func (h *HashTable) String() string {
	if h.Len < 1 {
		return "empty hash table"
	}
	ss := fmt.Sprintf("Len=%d\n", h.Len)
	for i, e := range h.Data {
		ss += fmt.Sprintf("[%d]%s\n", i, e)
	}
	return ss
}

//badHash is just what it sounds like - a bad hashing function
func (h *HashTable) badHash(s string) int {
	var res int
	for _, char := range s {
		res += int(char)
	}
	return res % h.Len
}

//Put takes a string argument for a key, and a byte slice argument for data. If the key already exists,
//the entry is updated. If the key does not exist, a new entry is created
func (h *HashTable) Put(key string, data []byte) {
	hash := h.badHash(key)
	h.Mu.Lock()
	defer h.Mu.Unlock()
	entry := &Entry{
		Key:  key,
		Data: data,
	}
	//check if there is an entry in the table
	res := h.Data[hash]

	if res == nil {
		h.Data[hash] = entry
		return
	}

	//check if the key matches
	for res.Key != key { //go to next entry while key doesn't match and next entry is available
		if res.Next != nil {
			res = res.Next
		} else {
			entry.Previous = res.Previous.Next //Jenked up as can be, but I can just use res, so I gotta be fancy with pointers
			res.Next = entry
		}
	}
	//If entry is found, update the data in the entry
	if res.Previous == nil {
		h.Data[hash] = entry
	} else {
		res.Previous.Next.Data = data //More poiner dancing
	}
}

//Search takes a string argument for a key and returns a pointer to an entry and a bool value. If the
//entry is found, it is returned with the value true. Otherwise, a pointer to an empty entry is returned,
//along with the value false
func (h *HashTable) Search(key string) (*Entry, bool) {
	hash := h.badHash(key)
	res := h.Data[hash]
	for res.Key != key {
		if res.Next != nil {
			res = res.Next
		} else {
			return &Entry{}, false
		}
	}
	return res, true
}

//Delete - pretty straightforward stuff here
func (h *HashTable) Delete(key string) error {
	hash := h.badHash(key)
	h.Mu.Lock()
	defer h.Mu.Unlock()
	res := h.Data[hash]
	for res.Key != key {
		if res.Next != nil {
			res = res.Next
		} else {
			return fmt.Errorf("Error: Entry for %q does not exist", key)
		}
	}
	h.Data[hash] = &Entry{} //kind of the nuclear option, I'll make it more specific later
	return nil
}
