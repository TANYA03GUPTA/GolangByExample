package main

import "fmt"

type board struct {
	square    [][]Symbol
	dimension int
}

func (b *board) printBoard() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.square[i][j] == Dot {
				fmt.Print(".")
			} else if b.square[i][j] == Cross {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}

		}
		fmt.Println("")
	}
}
func (b *board) markSymbol(i, j int, symbol Symbol) (bool, Symbol, error) {
	if i > b.dimension || j > b.dimension {
		return false, Dot, fmt.Errorf("index input is greater than dimension")
	}
	if b.square[i][j] != Dot {
		return false, Dot, fmt.Errorf("input square already marked")
	}
	if symbol != Cross && symbol != Circle {
		return false, Dot, fmt.Errorf("incorrect Symbol")
	}
	b.square[i][j] = symbol
	win := b.checkWin(i, j, symbol)

	return win, symbol, nil
}

func (b *board) checkWin(i, j int, symbol Symbol) bool {
	//Check Row
	rowMatch := true
	for k := 0; k < b.dimension; k++ {
		if b.square[i][k] != symbol {
			rowMatch = false
		}
	}

	if rowMatch {
		return rowMatch
	}

	//Check Row
	columnMatch := true
	for k := 0; k < b.dimension; k++ {
		if b.square[k][j] != symbol {
			columnMatch = false
		}
	}

	if columnMatch {
		return columnMatch
	}
	diagonalMatch := false
	if i == j {
		diagonalMatch = true
		for k := 0; k < b.dimension; k++ {
			if b.square[k][k] != symbol {
				diagonalMatch = false
			}
		}
	}

	return diagonalMatch

}