package goutils

import (
	"sync"
	"testing"
)

func TestSyncList_Add(t *testing.T) {
	list := NewSyncList[int]()
	numGoroutines := 100
	itemsPerGoroutine := 100

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(base int) {
			for j := 0; j < itemsPerGoroutine; j++ {
				list.Add(base + j)
			}
			wg.Done()
		}(i * itemsPerGoroutine)
	}
	wg.Wait()

	expectedLen := numGoroutines * itemsPerGoroutine
	if list.Len() != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, list.Len())
	}
	//for _, v := range list.ToSlice() {
	//	println(v)
	//}
}
