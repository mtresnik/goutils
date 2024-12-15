package mrutil

import "testing"

func TestToSet(t *testing.T) {
	tempSet := ToSet([]int{1, 2, 3, 4, 5})
	println(tempSet)
}

func TestSetToString(t *testing.T) {
	tempSet := ToSet([]int{1, 2, 3, 4, 5})
	println(SetToArray(tempSet))
}
