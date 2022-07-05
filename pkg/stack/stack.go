package stack

import "log"

type Stack struct {
	data   []string
	length int
}

func (s Stack) GetLen() int {
	return s.length
}

func (s *Stack) Push(element string) {
	s.length++
	s.data = append(s.data, element)
}

func (s *Stack) Pop() string {
	if s.length == 0 {
		log.Fatal("pop empty Stack")
	}
	returnValue := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return returnValue
}

func (s Stack) Top() string {
	if s.length == 0 {
		log.Fatal("top empty Stack")
	}
	return s.data[s.length-1]
}
