package repo


type Ladder struct {
	mp map[int]int
	size int
}

func NewLadder()*Ladder{
	return &Ladder{
		mp : make(map[int]int,0),
		size: 0,
	}
}

func (s *Ladder)AddLadder(start int, end int) bool {
	_, ok := s.mp[start]
	if ok {
		return false
	}
	s.mp[start] = end
	s.size++
	return true
}

func (s *Ladder)RemoveLadder(start int) bool {
	_, ok := s.mp[start]
	if !ok {
		return false
	}
	delete(s.mp,start)
	s.size--
	return false
}

func (s *Ladder)GetLadder(start int) (int, bool) {
	val, ok := s.mp[start]
	if ok {
		return val, true
	}
	return -1, false
}

