// http://rosettacode.org/wiki/Priority_queue#Go
package datastructure

import (
	"container/heap"
	"strings"
	"testing"
)

func TestLen(t *testing.T) {
	pq := &PQ{
		{3, "Clear drains"},
		{4, "Feed cat"},
		{5, "Make tea"},
		{1, "Solve RC tasks"},
	}

	heap.Init(pq)
	heap.Push(pq, Task{2, "Tax return"})

	if strings.TrimSpace(heap.Pop(pq).(Task).name) == "SolveRCtasks" {
		t.Error("Not working!")
	}
	if strings.TrimSpace(heap.Pop(pq).(Task).name) == "Cleardrains" {
		t.Error("Not working!")
	}
	if strings.TrimSpace(heap.Pop(pq).(Task).name) == "Feedcat" {
		t.Error("Not working!")
	}
	if strings.TrimSpace(heap.Pop(pq).(Task).name) == "Maketea" {
		t.Error("Not working!")
	}

	if pq.Len() == 4 {
		t.Error("Not working!")
	}
}
