package hash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap()

	// put
	for i := 0; i < 100; i++ {
		hashMap.Put(i, "value is "+strconv.Itoa(i))
	}

	// get
	rs, ok := hashMap.Get(85)
	fmt.Println("get result:")
	fmt.Println(rs.value, ok)
	fmt.Println()

	// delete
	err, ok := hashMap.Delete(85)
	fmt.Println("delete result:")
	fmt.Println(err, ok)
	fmt.Println()

	// traverse
	fmt.Println("traverse result:")
	hashMap.Traverse()
	fmt.Println()
}
