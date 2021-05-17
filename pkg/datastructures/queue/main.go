package queue

type (
	//Queue queue data type
	Queue []interface{}

	//Impl Queue Interface
	Impl interface {
		Enqueue(interface{}) interface{}
		Dequeue() interface{}
		Peek() interface{}
	}
)

//NewQueue create a new Queue
func NewQueue() Queue {
	return Queue{}
}

//NewQueueFromArray create a new queue from an array
func NewQueueFromArray(arr []interface{}) Queue {
	return arr
}

//Dequeue remove an element from the front of the queue
func (queue *Queue) Dequeue() interface{} {
	if len(*queue) > 0 {
		index := 0
		element := (*queue)[index]
		*queue = (*queue)[1:]
		return element
	}
	return ""
}

//Enqueue add an element to the back of the queue
func (queue *Queue) Enqueue(value interface{}) interface{} {
	*queue = append(*queue, value)
	return value
}

//Peek view the element at the front of the queue
func (queue Queue) Peek() interface{} {
	if len(queue) > 0 {
		return queue[0]
	}
	return ""
}
