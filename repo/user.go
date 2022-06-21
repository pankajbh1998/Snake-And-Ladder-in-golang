package repo

import "snake-and-ladder/models"

type User struct {
	mp map[int]*models.UserPos
	size int
}

func NewUser() *User {
	return &User{
		mp: make(map[int]*models.UserPos,0),
		size: 0,
	}
}

func (g *User)AddUser(userName string) {
	g.mp[g.size] = models.NewUser(userName)
	g.size++
}

func (g *User)UpdatePosition(id int, pos int){
	user, _ := g.mp[id]
	user.CurrPos = pos
}

func (g *User)GetUser(id int) *models.UserPos {
	user, _ := g.mp[id]
	return user
}

func (g *User)GetUserCount()int{
	return g.size
}