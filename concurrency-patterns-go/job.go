package concurrencypatterns

import (
	"reflect"
	"sort"
)

type Job struct {
	intermediateState int
	finalState        []int
}

func (j *Job) MarkIntermediateState(i int) {
	j.intermediateState = i
}

func (j *Job) Merge(jj Job) {
	j.finalState = append(j.finalState, jj.intermediateState)
}

func (j *Job) Verify(count int) bool {
	sort.Ints(j.finalState)
	expectedState := make([]int, 0, count)
	for i := 0; i < count; i++ {
		expectedState = append(expectedState, i)
	}
	return reflect.DeepEqual(j.finalState, expectedState)
}
