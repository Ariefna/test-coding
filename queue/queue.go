package queue

import "fmt"

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}


type QueueData struct {
	size int
	data []interface{}
}

func New(size int) Queue {
	return nil
}


func (q *QueueData) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : returns the next element in the queue
func (q *QueueData) Peek() (interface{}, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Queue : adds an element onto the queue and returns an pointer to the current queue
func (q *QueueData) Push(n interface{}) *QueueData {
	if q.Len() < q.size {
		q.data = append(q.data, n)
	} else {
		q.Pop()
		q.Push(n)
	}
	return q
}

// Dequeue : removes the next element from the queue and returns its value
// func (q *QueueFix) Pop() (interface{}, error) {
func (q *QueueData) Pop() interface{} {
	if len(q.data) == 0 {
		//return 0, fmt.Errorf("Queue is empty")
		return 0
	}
	element := q.data[0]
	q.data = q.data[1:]
	//return element, nil
	return element
}

func (q *QueueData) Len() int {
	return len(q.data)
}

func (q *QueueData) Keys() []interface{} {
	return q.data
}

func (q *QueueData) Contains(key interface{}) bool {
	cont := false
	for i := 0; i < q.Len(); i++ {
		if q.data[i] == key {
			cont = true
		}
	}
	return cont
}

var testValues = []interface{}{
	"lorem",
	"ipsum",
	1,
	2,
	3,
	"jack",
	"jill",
	"felix",
	"donking",
}
