package Manager

import (
	"fmt"
	"math/rand"
	"snake-and-ladder/service"
)

type Game struct{
	g *service.Strategy
	u *service.User
}

func NewManager(g *service.Strategy, u *service.User)*Game{
	return &Game{
		g : g,
		u: u,
	}
}
func (g *Game) RunGame(){
	n := g.u.GetUserCount()
	turn := 0
	g.g.UpdateGameStarted()
	for {
		diceVal := g.GetRandomDiceVal(g.g.GetDiceNumber())
		userPos := g.u.GetUserCurrPos(turn)
		g.u.UpdateUserPos(turn, g.g.GetUpdatedPos(userPos, diceVal))
		userName := g.u.GetUserName(turn)
		if g.u.CheckForWinner(turn) {
			fmt.Printf("Congratulations %v is won\n", userName)
			break
		} else {
			fmt.Printf("%v rolled a dice got %v and moved from %v to %v\n", userName, diceVal, userPos, g.u.GetUserCurrPos(turn))
		}
		turn = (turn+1) % n
	}
}

func (g *Game)GetRandomDiceVal(n int) int {
	ret := 0
	for i:=0;i<n;i++ {
		ret += rand.Int()%6 + 1
	}
	return ret
}