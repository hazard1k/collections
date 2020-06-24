package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()

	if s.length != 0 {
		t.Errorf("Размер созданного стека должен быть 0")
	}

	s.Push(1)

	if s.length != 1 {
		t.Errorf("Размер после добаления одного элемента должен быть 1")
	}

	s.Pop()

	if s.length != 0 {
		t.Errorf("Размер после извлечения одно элемент должен быть 0")
	}

	s.Push(1)
	s.Push(2)
	s.Push("yahoo")

	if s.top.value.(string) != "yahoo" {
		t.Errorf("верхний элемент стека должен быть \"yahoo\"")
	}

	s.Pop()
	s.Pop()
	s.Pop()

	n := s.Pop()
	if n != nil {
		t.Errorf("При извлчение из пустого стека должен вернуться nil")
	}
}
