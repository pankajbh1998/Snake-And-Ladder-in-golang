package service

import (
	"snake-and-ladder/repo"
)

type User struct {
	g *repo.Game
	u *repo.User
}

func NewUserService(u *repo.User, g *repo.Game)*User{
	return  &User{
		g: g,
		u: u,
	}
}

func (u *User)AddUser(userName string)bool{
	if u.g.IsGameStarted() {
		return false
	}
	u.u.AddUser(userName)
	return true
}

func (l *User)RemoveUser(userName string)bool{
	return false
}

func (l *User)GetUserCount()int{
	return l.u.GetUserCount()
}

func (l *User)GetUserName(id int)string{
	return l.u.GetUser(id).Name
}

func (l *User)GetUserCurrPos(id int)int{
	return l.u.GetUser(id).CurrPos
}

func (l *User)CheckForWinner(id int) bool {
	if l.u.GetUser(id).CurrPos == l.g.GetEndPos() {
		return true
	}
	return false
}

func (l *User)UpdateUserPos(id int, pos int){
	l.g.IsGameStarted()
	l.u.UpdatePosition(id, pos)
}



