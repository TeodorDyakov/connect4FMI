package main

import (
	"fmt"
)

type Board struct {
	board [][]string
	col   []int
}

const (
	BOARD_WIDTH  = 7
	BOARD_HEIGHT = 6
	EMPTY_SPOT   = "_"
)

func NewBoard() *Board {
	var b *Board
	b = new(Board)

	b.col = make([]int, BOARD_WIDTH)
	//initialize the connect 4 b.board
	for i := 0; i < BOARD_HEIGHT; i++ {
		row := make([]string, BOARD_WIDTH)

		for i := 0; i < len(row); i++ {
			row[i] = EMPTY_SPOT
		}
		b.board = append(b.board, row)
	}
	return b
}

func (b *Board) printBoard() {
	for i := 0; i < len(b.board[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[0]); j++ {
			fmt.Print(b.board[i][j] + " ")
		}
		fmt.Println()
	}
}

func (b *Board) undoDrop(column int) {
	b.col[column]--
	b.board[5-b.col[column]][column] = EMPTY_SPOT
}

func (b *Board) drop(column int, player string) bool {
	if column < len(b.board[0]) && (column >= 0) && b.col[column] < len(b.board) {
		b.board[5-b.col[column]][column] = player
		b.col[column]++
		return true
	}
	return false
}

func (b *Board) areFourConnected(player string) bool {
	// horizontalCheck
	for j := 0; j < len(b.board[0])-3; j++ {
		for i := 0; i < len(b.board); i++ {
			if b.board[i][j] == player &&
				b.board[i][j+1] == player &&
				b.board[i][j+2] == player &&
				b.board[i][j+3] == player {
				return true
			}
		}
	}
	// verticalCheck
	for i := 0; i < len(b.board)-3; i++ {
		for j := 0; j < len(b.board[0]); j++ {
			if b.board[i][j] == player &&
				b.board[i+1][j] == player &&
				b.board[i+2][j] == player &&
				b.board[i+3][j] == player {
				return true
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < len(b.board); i++ {
		for j := 0; j < len(b.board[0])-3; j++ {
			if b.board[i][j] == player &&
				b.board[i-1][j+1] == player &&
				b.board[i-2][j+2] == player &&
				b.board[i-3][j+3] == player {
				return true
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < len(b.board); i++ {
		for j := 3; j < len(b.board[0]); j++ {
			if b.board[i][j] == player &&
				b.board[i-1][j-1] == player &&
				b.board[i-2][j-2] == player &&
				b.board[i-3][j-3] == player {
				return true
			}
		}
	}
	return false
}
