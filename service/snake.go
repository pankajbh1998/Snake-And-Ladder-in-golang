package service

import "snake-and-ladder/repo"

type Snake struct {
	g *repo.Game
	s *repo.Snake
}

func NewSnakeService(s *repo.Snake, g *repo.Game)*Snake{
	return &Snake{
		s:s,
		g:g,
	}
}

func (s *Snake)AddSnake(start int, end int) bool {
	if s.g.IsGameStarted() {
		return false
	}
	return s.s.AddSnake(start, end)
}

func (s *Snake)RemoveSnake(start int) bool {
	return false
}

func (s *Snake)GetSnake(start int) (int, bool) {
	return s.s.GetSnake(start)
}
