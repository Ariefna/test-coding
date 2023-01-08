package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type myQueue struct {
	data  []interface{}
	count int
}

// New is a function to make new queue
func New(size int) Queue {
	return &myQueue{
		data: make([]interface{}, size),
	}
}

func (m *myQueue) Push(key interface{}) {
	if m.count == len(m.data) {
		m.Pop()
	}

	m.data[m.count] = key
	m.count++
}

func (m *myQueue) Pop() interface{} {
	if m.count <= 0 {
		panic("queue: Remove() called on empty queue")
	}
	popedData := m.data[0]
	m.data[0] = nil

	for i := 0; i < m.count-1; i++ {
		m.data[i] = m.data[i+1]
	}

	m.count--

	return popedData
}

func (m *myQueue) Contains(key interface{}) bool {
	contain := false
	for i := 0; i < m.Len(); i++ {
		if m.data[i] == key {
			contain = true
		}
	}

	return contain
}

func (m *myQueue) Len() int {
	return m.count
}

func (m *myQueue) Keys() []interface{} {
	return m.data[0:m.count]
}
