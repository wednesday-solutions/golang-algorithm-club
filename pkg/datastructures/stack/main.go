package stack

type (
	//Stack ...
	Stack []string
	//Impl Stack Interface
	Impl interface {
		Peek() string
		Pop() string
		Push(value string)
	}
)

//NewStack create a new stack
func NewStack() Stack {
	return Stack{}
}

//NewStackFromArray create a new stack from an existing array
func NewStackFromArray(arr []string) Stack {
	return arr
}

//Peek see the element at the top of the stack
func (stack Stack) Peek() string {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return ""
}

//Pop remove the element from the top of the stack
func (stack *Stack) Pop() string {
	if len(*stack) > 0 {
		index := len(*stack) - 1
		element := (*stack)[index]
		*stack = (*stack)[:index]
		return element
	}
	return ""
}

//Push add an element to the top of the stack
func (stack *Stack) Push(value string) {
	*stack = append(*stack, value)
}
