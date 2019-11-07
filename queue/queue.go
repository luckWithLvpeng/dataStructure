package queue

type Queue struct {
	list   []interface{}
	length int
}

// Init 初始化或清空队列
func (q *Queue) Init() *Queue {
	q.list = []interface{}{}
	q.length = 0
	return q
}

// Putin 入队
func (q *Queue) Putin(v interface{}) {
	q.list = append(q.list, v)
	q.length++
}

// Putout 出队
func (q *Queue) Putout(v interface{}) {
	q.list = q.list[1:]
	q.length--
}

// Length 数量
func (q *Queue) Length() int {
	return q.length
}

func NewQueue() *Queue {
	return new(Queue).Init()
}
