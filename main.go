package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mattn/go-gtk/gtk"
)

const (
	EMPTY  = 0
	PLAYER = 1
	BOT    = 2
)

type Game struct {
	board       [3][3]int
	turn        int
	playerWins  int
	botWins     int
	buttons     [3][3]*gtk.Button
	statusLabel *gtk.Label
	scoreLabel  *gtk.Label
	window      *gtk.Window
}

func NewGame(statusLabel, scoreLabel *gtk.Label, window *gtk.Window) *Game {
	game := &Game{
		turn:        PLAYER,
		statusLabel: statusLabel,
		scoreLabel:  scoreLabel,
		window:      window,
	}
	return game
}

func (g *Game) ResetGame() {
	g.board = [3][3]int{}
	g.turn = PLAYER
	g.updateStatus("Player's turn")
	for i := range g.buttons {
		for j := range g.buttons[i] {
			g.buttons[i][j].SetLabel(" ")
		}
	}
	gtk.MainIterationDo(false)
}

func (g *Game) MakeMove(x, y int, player int) bool {
	if x < 0 || x >= 3 || y < 0 || y >= 3 || g.board[x][y] != EMPTY {
		return false
	}
	g.board[x][y] = player
	g.buttons[x][y].SetLabel(g.playerSymbol(player))
	gtk.MainIterationDo(false) // Ensure GUI update
	return true
}

func (g *Game) playerSymbol(player int) string {
	switch player {
	case PLAYER:
		return "X"
	case BOT:
		return "O"
	default:
		return " "
	}
}

func (g *Game) CheckWin() int {
	for i := 0; i < 3; i++ {
		if g.board[i][0] != EMPTY && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
			return g.board[i][0]
		}
		if g.board[0][i] != EMPTY && g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] {
			return g.board[0][i]
		}
	}
	if g.board[0][0] != EMPTY && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
		return g.board[0][0]
	}
	if g.board[0][2] != EMPTY && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
		return g.board[0][2]
	}
	return EMPTY
}

func (g *Game) IsDraw() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == EMPTY {
				return false
			}
		}
	}
	return true
}

func (g *Game) BotMove() {
	bestMove := g.findBestMove()
	g.MakeMove(bestMove[0], bestMove[1], BOT)
}

func (g *Game) findBestMove() [2]int {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == EMPTY {
				g.board[i][j] = BOT
				if g.CheckWin() == BOT {
					g.board[i][j] = EMPTY
					return [2]int{i, j}
				}
				g.board[i][j] = EMPTY
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == EMPTY {
				g.board[i][j] = PLAYER
				if g.CheckWin() == PLAYER {
					g.board[i][j] = EMPTY
					return [2]int{i, j}
				}
				g.board[i][j] = EMPTY
			}
		}
	}

	for {
		x, y := rand.Intn(3), rand.Intn(3)
		if g.board[x][y] == EMPTY {
			return [2]int{x, y}
		}
	}
}

func (g *Game) updateStatus(msg string) {
	g.statusLabel.SetText(msg)
	gtk.MainIterationDo(false) // Ensure GUI update
}

func (g *Game) updateScore() {
	g.scoreLabel.SetText(fmt.Sprintf("Player: %d | Bot: %d", g.playerWins, g.botWins))
	gtk.MainIterationDo(false) // Ensure GUI update
}

func main() {
	gtk.Init(nil)
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	setupUI()
	gtk.Main()
}

func setupUI() {
	statusLabel := gtk.NewLabel("Player's turn")
	scoreLabel := gtk.NewLabel("Player: 0 | Bot: 0")
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Tic Tac Toe")
	window.SetDefaultSize(300, 300)

	game := NewGame(statusLabel, scoreLabel, window)

	grid := gtk.NewTable(3, 3, true)
	for i := range game.buttons {
		for j := range game.buttons[i] {
			x, y := i, j
			btn := gtk.NewButtonWithLabel(" ")
			btn.Clicked(func() {
				if game.turn == PLAYER && game.MakeMove(x, y, PLAYER) {
					game.turn = BOT
					game.updateStatus("Bot's turn")
					if winner := game.CheckWin(); winner != EMPTY {
						game.handleWin(winner)
						return
					}
					if game.IsDraw() {
						game.handleDraw()
						return
					}
					game.BotMove()
					game.updateStatus("Player's turn")
					if winner := game.CheckWin(); winner != EMPTY {
						game.handleWin(winner)
						return
					}
					if game.IsDraw() {
						game.handleDraw()
						return
					}
					game.turn = PLAYER
				}
			})
			game.buttons[x][y] = btn
			grid.AttachDefaults(btn, uint(x), uint(x+1), uint(y), uint(y+1))
		}
	}

	restartButton := gtk.NewButtonWithLabel("Restart")
	restartButton.Clicked(func() {
		game.ResetGame()
	})

	vbox := gtk.NewVBox(false, 2)
	vbox.PackStart(statusLabel, false, false, 0)
	vbox.PackStart(scoreLabel, false, false, 0)
	vbox.PackStart(grid, true, true, 0)
	vbox.PackStart(restartButton, false, false, 0)

	window.Add(vbox)
	window.ShowAll()

	window.Connect("destroy", func() {
		gtk.MainQuit()
	})
}

func (g *Game) handleWin(winner int) {
	if winner == PLAYER {
		g.playerWins++
		g.updateStatus("Player wins!")
	} else {
		g.botWins++
		g.updateStatus("Bot wins!")
	}
	g.updateScore()
	messageDialog := gtk.NewMessageDialog(
		g.window,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		g.statusLabel.GetText(),
	)
	messageDialog.Run()
	messageDialog.Destroy()
	g.ResetGame()
}

func (g *Game) handleDraw() {
	g.updateStatus("It's a draw!")
	messageDialog := gtk.NewMessageDialog(
		g.window,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		g.statusLabel.GetText(),
	)
	messageDialog.Run()
	messageDialog.Destroy()
	g.ResetGame()
}
