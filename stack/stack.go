package stack

// Stack Сам стек
type Stack struct {
	top    *Node
	length int
}

// Node элемент стека
type Node struct {
	value interface{}
	prev  *Node
}

// New Создание стека
func New() *Stack {
	return &Stack{nil, 0}
}

// Len Возвращает размер стека
func (s *Stack) Len() int {
	return s.length
}

// Pop Извлечение из стека
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	node := s.top
	s.top = node.prev
	s.length--
	return node.value
}

// Push добавление в стек
func (s *Stack) Push(n interface{}) {
	node := &Node{n, s.top}
	s.top = node
	s.length++
}
