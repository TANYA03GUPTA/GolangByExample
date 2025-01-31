package main

import "fmt"

type game struct {
    board           *board
    firstPlayer     iPlayer
    secondPlayer    iPlayer
    firstPlayerTurn bool
    moveIndex       int
    gameStatus      GameStatus
}

func initGame(b *board, p1, p2 iPlayer) *game {
    game := &game{
        board:           b,
        firstPlayer:     p1,
        secondPlayer:    p2,
        firstPlayerTurn: true,
        gameStatus:      GameInProgress,
    }
    return game
}
func (g *game) play() error {
    var win bool
    var symbol Symbol
    for {
        if g.firstPlayerTurn {
            x, y, err := g.firstPlayer.getNextMove()
            if err != nil {
                return err
            }
            win, symbol, err = g.board.markSymbol(x, y, g.firstPlayer.getSymbol())
            if err != nil {
                return err
            }
            g.firstPlayerTurn = false
            g.printMove(g.firstPlayer, x, y)
        } else {
            x, y, err := g.secondPlayer.getNextMove()
            if err != nil {
                return err
            }
            win, symbol, err = g.board.markSymbol(x, y, g.secondPlayer.getSymbol())
            if err != nil {
                return err
            }
            g.firstPlayerTurn = true
            g.printMove(g.secondPlayer, x, y)
        }
        g.moveIndex = g.moveIndex + 1
        g.setGameStatus(win, symbol)
        if g.gameStatus != GameInProgress {
            break
        }
    }
	return nil
}

func (g *game) setGameStatus(win bool, symbol Symbol) {
    if win {
        if g.firstPlayer.getSymbol() == symbol {
            g.gameStatus = FirstPlayerWin
            return
        } else if g.secondPlayer.getSymbol() == symbol {
            g.gameStatus = SecondPlayerWin
            return
        }
    }
    if g.moveIndex == g.board.dimension*g.board.dimension {
        g.gameStatus = GameDraw
        return
    }
	g.gameStatus = GameInProgress
}

func (g *game) printMove(player iPlayer, x, y int) {
	symbolString := ""
	symbol := player.getSymbol()
	if symbol == Cross {
		symbolString = "*"
	} else if symbol == Circle {
		symbolString = "o"
	}
	fmt.Printf("Player %d Move with Symbol %s at Position X:%d Y:%d\n", player.getID(), symbolString, x, y)
	g.board.printBoard()
	fmt.Println("")
}
func (g *game) printResult() {
    switch g.gameStatus {
    case GameInProgress:
        fmt.Println("Game in Between")
    case GameDraw:
        fmt.Println("Game Drawn")
    case FirstPlayerWin:
        fmt.Println("First Player Win")
    case SecondPlayerWin:
        fmt.Println("Second Player Win")
    default:
        fmt.Println("Invalid Game Status")
    }
    g.board.printBoard()
} 