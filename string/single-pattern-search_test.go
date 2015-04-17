package string

import (
	"fmt"
	"testing"
)

func Test_single_pattern_search(t *testing.T) {
	source := "This is Rinc Liu's personal website."
	pattern := "Rinc Liu"
	fmt.Printf("KMP: %d\n", KMPSearch(source, pattern))
	fmt.Printf("BM: %d\n", BMSearch(source, pattern))
}