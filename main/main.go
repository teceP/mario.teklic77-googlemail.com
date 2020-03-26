package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"tictactoe_go/utils"
)

const PLAYER_NO = 0
const PLAYER_A = 1
const PLAYER_B = 2
const PLAYER_A_SIGN = 'X'
const PLAYER_B_SIGN = 'O'

type Board struct {
	coords [3][3]Coordinate
}

type Coordinate struct {
	x      int
	y      int
	player int
}

type Player struct {
	sign rune
}

type Game struct {
	moves  int
	active int
	board  Board
}

func main() {
	fmt.Println("Welcome to TicTacToe!")
	var game = Game{}

	//Preps
	game.board.prepareBoard()
	game.chooseBeginner()

	//Game loop
	for {
		c := game.nextMove()
		game.makeMove(c)
		game.printBoard()
		if game.checkWinner() {
			break
		}
		game.nextPlayer()
	}

	fmt.Println(game.activeSing(), " has won!\nGoodbye.")
}

func (b *Board) prepareBoard() {
	for i, f := range b.coords {
		for j, fj := range b.coords[i] {
			b.coords[i][j].x = i
			b.coords[i][j].y = j
			b.coords[i][j].player = PLAYER_NO
			_ = fj
		}
		_ = f
	}
}

func (g *Game) chooseBeginner() {
	i := rand.Intn(2) //.. 1 or 2 -> PLAYER_A or PLAYER_B
	i = i + 1
	g.active = i

	if i == 1 {
		fmt.Println("X begins!")
	} else if i == 2 {
		fmt.Println("O begins!")
	} else {
		fmt.Println("error!")

	}
}

func (g *Game) nextPlayer() {
	if g.active == PLAYER_A {
		g.active = PLAYER_B
	} else {
		g.active = PLAYER_A
	}
}

func (g *Game) nextMove() Coordinate {
	var c = Coordinate{-1, -1, g.active}
	good := false

	for !good {
		fmt.Print(g.activeSing(), ", please insert coords (Example: A/1): ")

		input := utils.ReadInput()
		//Convert input like "A/1" to 0/0 or "B/1" to 1/0
		c.x, c.y = utils.ConvertInput(input)

		good = utils.CheckInputLength(input) && g.checkAvailability(c)
	}

	fmt.Println("Shot will be perfomed!")
	return c
}

func (g *Game) makeMove(c Coordinate) {
	g.board.coords[c.x][c.y].player = c.player
}

func (g *Game) checkAvailability(c Coordinate) bool {

	if checkRange(c) {
		if g.board.coords[c.x][c.y].player == PLAYER_NO {
			return true
		}
	}

	fmt.Println("Field is not available.")
	return false
}

func checkRange(c Coordinate) bool {
	if c.x < 0 || c.x > 2 {
		return false
	} else if c.y < 0 || c.y > 2 {
		return false
	}
	return true
}

func (g *Game) activeSing() string {
	if g.active == PLAYER_A {
		return strconv.QuoteRune(PLAYER_A_SIGN)
	} else {
		return strconv.QuoteRune(PLAYER_B_SIGN)
	}
}

func (b *Board) coordSign(x int, y int) string {
	if b.coords[x][y].player == PLAYER_A {
		return strconv.QuoteRune(PLAYER_A_SIGN)
	} else if b.coords[x][y].player == PLAYER_B {
		return strconv.QuoteRune(PLAYER_B_SIGN)
	} else {
		return "~"
	}
}

func (g *Game) printBoard() {
	fmt.Println("=========================")

	for i, f := range g.board.coords {
		fmt.Print("|  ")

		for j, fj := range g.board.coords[i] {
			fmt.Print(g.board.coordSign(i, j), "  |  ")

			_ = fj
		}

		fmt.Println(" \n-------------------------")
		_ = f
	}

}

func (g *Game) checkWinner() bool {
	won := false

	for i, f := range g.board.coords {
		for j, fj := range g.board.coords[i] {

			//  -
			if g.board.coords[i][0].player == g.active &&
				g.board.coords[i][1].player == g.active &&
				g.board.coords[i][2].player == g.active {
				won = true
			}

			//  |
			if g.board.coords[0][j].player == g.active &&
				g.board.coords[1][j].player == g.active &&
				g.board.coords[2][j].player == g.active {
				won = true
			}
			_ = fj
		}
		_ = f
	}

	//  \
	if g.board.coords[0][0].player == g.active &&
		g.board.coords[1][1].player == g.active &&
		g.board.coords[2][2].player == g.active {
		won = true
	}

	//  /
	if g.board.coords[0][2].player == g.active &&
		g.board.coords[1][1].player == g.active &&
		g.board.coords[2][0].player == g.active {
		won = true
	}

	return won
}
