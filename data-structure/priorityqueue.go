// http://rosettacode.org/wiki/Priority_queue#Go
package datastructure

type Task struct {
	priority int
	name     string
}

type PQ []Task

func (self PQ) Len() int {
	return len(self)
}

func (self PQ) Less(i, j int) bool {
	return self[i].priority < self[j].priority
}

func (self PQ) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self *PQ) Push(x interface{}) {
	*self = append(*self, x.(Task))
}

func (self *PQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}
