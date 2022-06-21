package repo

type Game struct{
	startPos int
	endPos int
	isGameStarted bool
	diceCount int
}

func NewGame(s int, e int, diceCount int) *Game{
	return &Game{
		startPos:      s,
		endPos:        e,
		diceCount:  diceCount,
		isGameStarted: false,
	}
}

func (g *Game)GetEndPos()int{
	return g.endPos
}

func (g *Game)GetStartPos()int{
	return g.startPos
}

func (g *Game)IsGameStarted()bool{
	return g.isGameStarted
}

func (g *Game)GetDiceCount()int{
	return g.diceCount
}

func (g *Game)UpdateGameStarted(){
	g.isGameStarted = true
}
