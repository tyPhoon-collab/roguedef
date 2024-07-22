package ds

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Push(t T) {
	q.data = append(q.data, t)
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.IsEmpty() {
		var t T
		return t, false
	}
	t := q.data[0]
	q.data = q.data[1:]
	return t, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var t T
		return t, false
	}
	return q.data[0], true
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Data() []T {
	return q.data
}

func (q *Queue[T]) Insert(i int, t T) {
	q.data = append([]T{t}, q.data...)
}

func (q *Queue[T]) Clear() {
	q.data = []T{}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func NewQueueFrom[T any](data []T) *Queue[T] {
	return &Queue[T]{data: data}
}
