package hashtable

import (
	"fmt"
	"sync"
)

//Entry is what we put into the hash table
type Entry struct {
	Key  string //to double-check for collisions
	Data []byte //maybe a bit jenky, but with slices it don't even matter. I mean it does but I'm not bothered
	Next *Entry
}

//HashTable is where the data is stored and what we implement methods on. It will take a string as a
//key and a byte slice as data
type HashTable struct {
	Len  int
	Data []*Entry
	Mux  sync.Mutex
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
	h.Mux.Lock()
	//check if there is an entry in the table
	res := h.Data[hash]
	if res.Data != nil { //if entry is not empty
		//check if the key matches
		for res.Key != key { //go to next entry while key doesn't match and next entry is available
			if res.Next != nil {
				res = res.Next
			} else {
				break
			}
		}
		res.Next = &Entry{
			Key:  key,
			Data: data,
		}
	}
	h.Data[hash] = &Entry{
		Key:  key,
		Data: data,
	}
	h.Mux.Unlock()
}

//Search taeks a string argument for a key and returns a pointer to an entry and a bool value. If the
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
	h.Mux.Lock()
	res := h.Data[hash]
	for res.Key != key {
		if res.Next != nil {
			res = res.Next
		} else {
			return fmt.Errorf("Error: Entry for %q does not exist", key)
		}
	}
	res = &Entry{} //this is pretty lousy, but hey what the heck
	h.Mux.Unlock()
	return nil
}
