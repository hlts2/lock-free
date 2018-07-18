package gomaphore

import (
	"reflect"
	"sync"
	"testing"
)

const size = 1000

func TestGomaphore(t *testing.T) {
	semaphore := new(Gomaphore)

	cnt := 0
	nums := make([]int, 0, size)
	wg := new(sync.WaitGroup)

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore.Wait()
			nums = append(nums, cnt)
			cnt++
			semaphore.Signal()
		}(i)
	}

	wg.Wait()

	expected := make([]int, 0, size)
	for i := 0; i < size; i++ {
		expected = append(expected, i)
	}

	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("Gomaphore is wrong. expected: %v, got: %v", expected, nums)
	}
}
