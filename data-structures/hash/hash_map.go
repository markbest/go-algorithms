package hash

import (
	"errors"
	"fmt"
)

type Entry struct {
	key   int
	value string
	next  *Entry
}

type HashMap struct {
	buckets []*Entry
}

const maxCapacity = 10

// New hash map
func NewHashMap() *HashMap {
	return &HashMap{buckets: make([]*Entry, maxCapacity)}
}

// Get key hash
func (h *HashMap) getKeyHash(key int) int {
	return key % maxCapacity
}

// delete key
func (h *HashMap) Delete(key int) (error, bool) {
	e, ok := h.Get(key)
	if ok {
		hash := h.getKeyHash(key)
		p := h.buckets[hash].next
		for p != nil {
			if p.next == e {
				p.next = p.next.next
			} else {
				p = p.next
			}
		}
		return nil, true
	} else {
		return errors.New("no key exist"), false
	}
}

// Get key
func (h *HashMap) Get(key int) (*Entry, bool) {
	hash := h.getKeyHash(key)
	p := h.buckets[hash]
	if p == nil {
		return nil, false
	}

	if p.next == nil {
		return p, true
	}

	for p.next != nil {
		if p.key == key {
			return p, true
		} else {
			p = p.next
		}
	}
	return nil, false
}

// Put key & value
func (h *HashMap) Put(key int, value string) {
	e := &Entry{key, value, nil}
	hash := h.getKeyHash(key)
	p := h.buckets[hash]
	if p == nil {
		h.buckets[hash] = e
		return
	}

	for p.next != nil {
		if p.next.key == e.key {
			p.next.value = e.value
			return
		} else {
			p = p.next
		}
	}
	p.next = e
}

// Hash map traverse
func (h *HashMap) Traverse() {
	for i := 0; i < maxCapacity; i++ {
		p := h.buckets[i]
		if p == nil {
			continue
		}

		fmt.Println(p.key, ":", p.value)
		if p.next != nil {
			for p.next != nil {
				fmt.Println("sub key: ", p.next.key, ":", p.next.value)
				if p.next == nil {
					break
				}
				p.next = p.next.next
			}
		}
	}
}
