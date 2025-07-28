package iso6391

import "testing"

func TestList(t *testing.T) {
	x := make(map[string]struct{})

	for _, v := range Languages {
		if _, has := x[v]; has {
			t.Fatalf("Languages map contains duplicate values")
		}
		x[v] = struct{}{}
	}
}
