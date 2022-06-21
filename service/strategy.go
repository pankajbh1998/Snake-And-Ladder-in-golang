package service

import "snake-and-ladder/repo"

type Strategy struct{
	g *repo.Game
	u *repo.User
	l *repo.Ladder
	s *repo.Snake
}

func NewStrategyService(u *repo.User, l *repo.Ladder, s *repo.Snake, g *repo.Game)*Strategy{
	return  &Strategy{
		u: u,
		l:l,
		s:s,
		g: g,
	}
}

func(s *Strategy)GetUpdatedPos(pos int, diceVal int)int{
	newPos := pos + diceVal
	p,ok := s.l.GetLadder(newPos)
	if pos > s.g.GetEndPos() {
		return pos
	}
	if ok {
		return p
	}
	p,ok = s.s.GetSnake(newPos)
	if ok {
		return p
	}
	return newPos
}

func (s *Strategy)GetDiceNumber()int{
	return s.g.GetDiceCount()
}

func (s *Strategy)UpdateGameStarted(){
	s.g.UpdateGameStarted()
}