package main

import (
	"snake-and-ladder/Manager"
	"snake-and-ladder/repo"
	"snake-and-ladder/service"
)

func main(){
	startPos := 0
	endPos := 100
	diceCont := 1
	game := repo.NewGame(startPos, endPos,diceCont)
	snake := repo.NewSnake()
	ladder := repo.NewLadder()
	user := repo.NewUser()
	snakeService := service.NewSnakeService(snake, game)
	ladderService := service.NewLadderService(ladder, game)
	userService := service.NewUserService(user, game)
	getSnakes(snakeService)
	getLadders(ladderService)
	addUsers(userService)
	man := Manager.NewManager(service.NewStrategyService(user, ladder, snake, game), userService)
	man.RunGame()
}

func getSnakes(snake *service.Snake){
	//9
	//62 5
	//33 6
	//49 9
	//88 16
	//41 20
	//56 53
	//98 64
	//93 73
	//95 75
	snake.AddSnake(62, 5)
	snake.AddSnake(33,6)
	snake.AddSnake(49, 9)
	snake.AddSnake(88, 16)
	snake.AddSnake(41, 20)
	snake.AddSnake(56, 53)
	snake.AddSnake(98, 64)
	snake.AddSnake(93, 73)
	snake.AddSnake(95, 75)
}

func getLadders(l *service.Ladder){
	//8
	//2 37
	//27 46
	//10 32
	//51 68
	//61 79
	//65 84
	//71 91
	//81 100
	l.AddLadder(2, 37)
	l.AddLadder(27, 46)
	l.AddLadder(10, 32)
	l.AddLadder(51, 68)
	l.AddLadder(61, 79)
	l.AddLadder(65, 84)
	l.AddLadder(71, 91)
	l.AddLadder(81, 100)
}

func addUsers(service *service.User){
	users := []string{
		"Pankaj",
		"Shubham",
		"Shivam",
	}
	for _, val := range users {
		service.AddUser(val)
	}
}