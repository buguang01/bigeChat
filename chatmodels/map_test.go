package ChatModels_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	a := make(map[int]struct{})

	for i := 0; i <= 100; i++ {
		a[i] = struct{}{}
	}

	for k, v := range a {
		_ = v
		delete(a, k)
	}
}

// func TestFilterModel(t *testing.T) {

// 	a := ChatModels.FilterModel{"", make(map[string]ChatModels.FilterModel)}
// 	a.Subli["abc"] = ChatModels.FilterModel{"abc", make(map[string]ChatModels.FilterModel)}
// 	a.Subli["dd"] = ChatModels.FilterModel{"dd", make(map[string]ChatModels.FilterModel)}
// 	s := a.Subli["dd"]
// 	s.NodeStr = "ff"
// 	fmt.Println(a)
// 	fmt.Println(s)
// }
