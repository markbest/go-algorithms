package search

import (
	"fmt"
	"testing"
)

func TestSearchString(t *testing.T) {
	target := "why every programming language use the hello world as the first test???"
	search := "hello world"
	rs := NormalSearchString([]byte(target), []byte(search))
	fmt.Println(rs)

	rs = KMPStringSearch([]byte(target), []byte(search))
	fmt.Println(rs)
}
