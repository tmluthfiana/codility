//https://github.com/theodesp/codility-go/blob/master/stacks-queues/stack.go
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
type IntStack struct {
	size int
	data []int
}

func NewIntStack(len int) *IntStack {
	return &IntStack{
		size: 0,
		data: make([]int, len),
	}
}

func (s *IntStack) Front() int {
	return s.data[s.size-1]
}

func (s *IntStack) Push(item int) {
	if s.size < len(s.data) {
		s.data[s.size] = item
		s.size++
	}
}

func (s *IntStack) Pop() int {
	item := s.data[s.size-1]
	s.size -= 1

	return item
}

func Solution(H []int) int {
	// write your code in Go 1.4
	l := NewIntStack(len(H))
	result := 0

	for _, value := range H {
		for l.size > 0 && l.Front() > value {
			l.Pop()
		}

		if l.size == 0 || l.Front() < value {
			l.Push(value)
			result += 1
		}
	}

	return result
}
