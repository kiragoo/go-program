package difstruct

import "testing"

func TestInfFn(t *testing.T) {
	// fa := SA{
	// 	Name:  "a",
	// 	Value: 1,
	// }
	fb := SB{
		Name:  "b",
		Value: 2,
		Age:   3,
	}

	// InfFn(&fa)
	InfFn(&fb)
}
