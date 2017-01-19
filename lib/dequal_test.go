// Coverage: 94.9% 
package dequal

import "testing"

// Circular linked lists a -> b -> a and c -> c.
type link struct {
	value string
	tail  *link
}

func TestEqual(t *testing.T) {
	var tmpFunc = func() {}
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	map3 := make(map[int]int)
	map1[0] = 1
	map1[1] = 2
	map2[0] = 1
	map2[1] = 1
	map3[0] = 1

	a, b, c, d := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}, &link{value: "c"}
	a.tail, b.tail, c.tail, d.tail = b, a, c, c

	tests := []struct {
		x, y interface{}
		want bool
	}{
		{nil, 0, false},
		{1, 1, true},
		{1, "1", false},
		{true, false, false},
		{uint(1), uint(1), true},
		{tmpFunc, tmpFunc, true},
		{&tmpFunc, &tmpFunc, true},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string(nil), []string{}, true},
		{map[string]int(nil), map[string]int{}, true},
		{map1, map2, false},
		{map1, map3, false},
		{c, d, true},
		{a, a, true},
		{b, b, true},
		{c, c, true},
		{a, b, false},
		{a, c, false},
	}

	for _, test := range tests {
		if got, want := Equal(test.x, test.y), test.want; got != want {
			t.Errorf("DeepEqual(%v, %v) = %t, want %t", test.x, test.y, got, want)
		}
	}
}
