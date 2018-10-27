package queue

import (
	"sync"
)

type Queue interface {
	Enqueue(int)
	Dequeue() (int, bool)
	HasNext() bool
}

type queueImpl struct {
	lock   *sync.Mutex
	values []int
}

func NewQueue() Queue {
	return &queueImpl{
		lock:   &sync.Mutex{},
		values: make([]int, 0),
	}
}

func (q *queueImpl) Enqueue(x int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.values = append(q.values, x)

}

func (q *queueImpl) Dequeue() (int, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	result := 0
	if q.HasNext() {
		result = q.values[0]
		q.values = q.values[1:]
		return result, true
	}
	return 0, false
}

func (q *queueImpl) HasNext() bool {
	return len(q.values) > 0
}
