package queue

type (
	//Queue queue data type
	Queue []string
	//Impl Queue Interface
	Impl interface {
		Enqueue(string) string
		Dequeue() string
		Peek() string
	}
)

//NewQueue create a new Queue
func NewQueue() Queue {
	return Queue{}
}

//NewQueueFromArray create a new queue from an array
func NewQueueFromArray(arr []string) Queue {
	return arr
}

//Dequeue remove an element from the front of the queue
func (queue *Queue) Dequeue() string {
	if len(*queue) > 0 {
		index := 0
		element := (*queue)[index]
		*queue = (*queue)[index:]
		return element
	}
	return ""

}

//Enqueue add an element to the back of the queue
func (queue *Queue) Enqueue(value string) string {
	*queue = append(*queue, value)
	return value
}

//Peek view the element at the front of the queue
func (queue Queue) Peek() string {
	if len(queue) > 0 {
		return queue[0]
	}
	return ""
}
