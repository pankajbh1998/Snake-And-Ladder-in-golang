package repo


type Snake struct {
	mp map[int]int
	size int
}

func NewSnake()*Snake{
	return &Snake{
		mp : make(map[int]int,0),
		size: 0,
	}
}

func (s *Snake)AddSnake(start int, end int) bool {
	_, ok := s.mp[start]
	if ok {
		return false
	}
	s.mp[start] = end
	s.size++
	return true
}

func (s *Snake)RemoveSnake(start int) bool {
	_, ok := s.mp[start]
	if !ok {
		return false
	}
	delete(s.mp,start)
	s.size--
	return false
}

func (s *Snake)GetSnake(start int) (int, bool) {
	val, ok := s.mp[start]
	if ok {
		return val, true
	}
	return -1, false
}

