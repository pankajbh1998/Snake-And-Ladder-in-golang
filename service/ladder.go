package service

import "snake-and-ladder/repo"

type Ladder struct {
	g *repo.Game
	l *repo.Ladder
}

func NewLadderService(l *repo.Ladder, g *repo.Game)*Ladder{
	return &Ladder{
		l:l,
		g:g,
	}
}

func (s *Ladder)AddLadder(start int, end int) bool {
	if s.g.IsGameStarted() {
		return false
	}
	return s.l.AddLadder(start, end)
}

func (s *Ladder)RemoveLadder(start int) bool {
	return false
}

func (s *Ladder)GetLadder(start int) (int, bool) {
	return s.l.GetLadder(start)
}
