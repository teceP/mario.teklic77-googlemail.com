package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Consts

const PLAYER_NO = 0
const PLAYER_A = 1
const PLAYER_B = 2
const PLAYER_A_SIGN = 'X'
const PLAYER_B_SIGN = 'O'

// Structs

//Represents the fields of the board
type Board struct {
	coords [3][3]Coordinate
}

//Each field can be owned by a player (PLAYER_A (1) or PLAYER_B (2)) or can be free (PLAYER_NO (0))
type Coordinate struct {
	x      int
	y      int
	player int
}

//The game struct, which contains several needed variables:
//moves = number of moves
//active = current player
//board = board
//computer = if true: the human plays against the computer
type Game struct {
	moves    int
	active   int
	board    Board
	computer bool
}

func main() {
	fmt.Println("Welcome to TicTacToe!")
	var game = Game{}

	//Let the user decide if he wants to play against the computer or not:
	fmt.Print("Do you want to play against the computer? Y/n ")
	usersDec := ReadInput()
	var err error
	game.computer, err = EvaluateUserDec(usersDec)

	for err != nil {
		fmt.Println(err)
		fmt.Print("Do you want to play against the computer? Y/n ")
		usersDec = ReadInput()
		game.computer, err = EvaluateUserDec(usersDec)
	}

	//Preps
	game.board.prepareBoard()
	game.chooseBeginner()

	//Game loop
	for {
		c := game.nextMove()
		game.makeMove(c)
		game.printBoard()

		if game.checkWinner() {
			fmt.Println(game.activeSing(), " has won!")
			break
		}

		if game.moves == 9 {
			fmt.Println("Tie!")
		}
		game.nextPlayer()
	}

	fmt.Println("Goodbye.")

	duration := time.Duration(10) * time.Second
	time.Sleep(duration)
}

// Board actions

//Sets PLAYER_NO on each field (init func)
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

//Chooses which player will begin
func (g *Game) chooseBeginner() {
	i := rand.Intn(2) //.. 1 or 2 -> PLAYER_A or PLAYER_B
	i = i + 1
	g.active = i

	if i == 1 {
		fmt.Println("X begins!")
	} else if i == 2 {
		fmt.Println("O begins!")
	}
}

//Changes the current player to the next player
func (g *Game) nextPlayer() {
	if g.active == PLAYER_A {
		g.active = PLAYER_B
	} else {
		g.active = PLAYER_A
	}
}

// Moves

//Next move either by human, or redirects to computers automatic-move
func (g *Game) nextMove() Coordinate {

	//If computer has to choose the next move
	if g.computer && g.active == PLAYER_B {
		return g.bestMove()
	}

	var c = Coordinate{-1, -1, g.active}
	good := false

	for !good {
		fmt.Print(g.activeSing(), ", please insert coords (Example: A/1): ")

		input := ReadInput()
		//Convert input like "A/1" to 0/0 or "B/1" to 1/0
		c.x, c.y = ConvertInput(input)

		good = CheckInputLength(input) && g.checkAvailability(c)
	}

	fmt.Println("Shot will be perfomed!")
	return c
}

//Evaluates a board by its placed moves
//Retuns 0 if no one won
func evaluate(b Board) int {
	for i := range b.coords {
		// -
		if b.coords[i][0].player == b.coords[i][1].player &&
			b.coords[i][1].player == b.coords[i][2].player {
			if b.coords[i][0].player == PLAYER_B {
				return +10
			} else if b.coords[i][0].player == PLAYER_A {
				return -10
			}
		}

		// |
		if b.coords[0][i].player == b.coords[1][i].player &&
			b.coords[1][i].player == b.coords[2][i].player {
			if b.coords[0][i].player == PLAYER_B {
				return +10
			} else if b.coords[0][i].player == PLAYER_A {
				return -10
			}
		}

		// \
		if b.coords[0][0].player == b.coords[1][1].player &&
			b.coords[1][1].player == b.coords[2][2].player {
			if b.coords[0][0].player == PLAYER_B {
				return +10
			} else if b.coords[0][0].player == PLAYER_A {
				return -10
			}
		}

		// /
		if b.coords[0][2].player == b.coords[1][1].player &&
			b.coords[1][1].player == b.coords[2][0].player {
			if b.coords[0][2].player == PLAYER_B {
				return +10
			} else if b.coords[0][2].player == PLAYER_A {
				return -10
			}
		}
	}
	return 0
}

//Algorithms which calculates the next smartest move
func minimax(b Board, depth int, isMax bool) int {
	score := evaluate(b)

	// If max has won game, return his evaluated score (10) || if min has won game, return his e... (-10)
	if score == 10 || score == -10 {
		return score
	}

	//if no more moves & no winner == tie
	if len(emptyCells(b)) == 0 {
		return 0
	}

	//if its max's move
	if isMax {
		best := -10000.0

		for i := range b.coords {
			for j := range b.coords[i] {
				//check if empty
				if b.coords[i][j].player == PLAYER_NO {
					//make move
					b.coords[i][j].player = PLAYER_B

					//call minimax recur.
					best = math.Max(float64(best), float64(minimax(b, depth+1, !isMax)))

					//undo move
					b.coords[i][j].player = PLAYER_NO
				}
			}
		}
		//test, use with care..
		return int(best)
	} else {
		// min's move
		best := 10000.0

		//trav all cells

		for i := range b.coords {
			for j := range b.coords[i] {
				//check emptyness
				if b.coords[i][j].player == PLAYER_NO {
					//make (hoooomans) move
					b.coords[i][j].player = PLAYER_A

					//call minimax recurs., choose min val
					best = math.Min(float64(best), float64(minimax(b, depth+1, !isMax)))

					//undo move
					b.coords[i][j].player = PLAYER_NO
				}
			}
		}
		return int(best)
	}
}

//Receives the next smartest move
func (g *Game) bestMove() Coordinate {
	//gc == gameCopy
	//gc := g.copyGame()
	bestMove := Coordinate{
		x:      -1,
		y:      -1,
		player: g.active,
	}

	bestVal := -10000

	for i := range g.board.coords {
		for j := range g.board.coords[i] {
			if g.board.coords[i][j].player == PLAYER_NO {
				currentMove := Coordinate{
					x:      i,
					y:      j,
					player: g.active,
				}

				//Make move
				g.board.coords[currentMove.x][currentMove.y].player = g.active

				//compute elevation func for this move
				moveVal := minimax(g.board, 0, false)

				//undo move
				g.board.coords[i][j].player = PLAYER_NO

				//if value of current move better than best, update
				if moveVal > bestVal {
					bestMove.x, bestMove.y = i, j
					bestVal = moveVal
				}
			}
		}
	}
	return bestMove
}

//Returns all leftover (free) cells of the board
func emptyCells(b Board) (cells []Coordinate) {
	for i := range b.coords {
		for j := range b.coords[i] {
			if b.coords[i][j].player == PLAYER_NO {
				cells = append(cells, b.coords[i][j])
			}
		}
	}
	return
}

//Makes a move and increments the move var
func (g *Game) makeMove(c Coordinate) {
	g.board.coords[c.x][c.y].player = c.player
	g.moves = g.moves + 1
}

// Signs

//Returns the sign of the active player
func (g *Game) activeSing() string {
	if g.active == PLAYER_A {
		return strconv.QuoteRune(PLAYER_A_SIGN)
	} else {
		return strconv.QuoteRune(PLAYER_B_SIGN)
	}
}

//Returns the sign of a specific field, depending on the status of the field
func (b *Board) coordSign(x int, y int) string {
	if b.coords[x][y].player == PLAYER_A {
		return strconv.QuoteRune(PLAYER_A_SIGN)
	} else if b.coords[x][y].player == PLAYER_B {
		return strconv.QuoteRune(PLAYER_B_SIGN)
	} else {
		return "~"
	}
}

// Prints

//Prints the board
func (g *Game) printBoard() {
	fmt.Println("\n=========================")
	for i := range g.board.coords {
		fmt.Print("|  ")
		for j := range g.board.coords[i] {
			fmt.Print(g.board.coordSign(i, j), "  |  ")
		}
		fmt.Println(" \n-------------------------")
	}
}

// Proofs

//Checks if the field is available
//Returns false, if field is busy
func (g *Game) checkAvailability(c Coordinate) bool {

	if checkRange(c) {
		if g.board.coords[c.x][c.y].player == PLAYER_NO {
			return true
		}
	}

	fmt.Println("Field is not available.")
	return false
}

//Checks if the coordinate x and y is in the range of the board
func checkRange(c Coordinate) bool {
	if c.x < 0 || c.x > 2 {
		return false
	} else if c.y < 0 || c.y > 2 {
		return false
	}
	return true
}

//Checks if there is a winner
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
