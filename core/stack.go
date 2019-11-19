package core


type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return -1, nil
	}

	lastOne := s.items[len(s.items)-1]
	s.items = append(s.items[:len(s.items)-1])
	return lastOne, nil
}