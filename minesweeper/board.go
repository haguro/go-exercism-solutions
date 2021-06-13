package minesweeper

import (
	"bytes"
	"errors"
)

// Board represents a Minesweeper game board.
type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

// Count counts the number of mines adjacent to the board's empty cells and
// adds those counts to the board.
func (b Board) Count() error {
	if err := b.validateBorder(); err != nil {
		return err
	}
	for i := 1; i < len(b)-1; i++ {
		for j := 1; j < len(b[i])-1; j++ {
			if b[i][j] != ' ' && b[i][j] != '*' {
				return errors.New("board includes unknown character(s)")
			}

			if b[i][j] == ' ' {
				var n byte
				upIdx, downIdx := i-1, i+1
				leftIdx, rightIdx := j-1, j+1
				if i < 2 {
					upIdx = i
				}
				if i >= len(b)-2 {
					downIdx = i
				}
				if j < 2 {
					leftIdx = j
				}
				if j >= len(b[i])-2 {
					rightIdx = j
				}
				for v := upIdx; v <= downIdx; v++ {
					for w := leftIdx; w <= rightIdx; w++ {
						if b[v][w] == '*' {
							n++
						}
					}
				}

				if n > 0 {
					b[i][j] = '0' + n
				}
			}
		}
	}
	return nil
}

func (b Board) validateBorder() error {
	var err bool
	//corners
	if b[0][0] != '+' ||
		b[0][len(b[0])-1] != '+' ||
		b[len(b)-1][0] != '+' ||
		b[len(b)-1][len(b[0])-1] != '+' {
		err = true
	}
	//vertical borders
	for i := 1; i < len(b)-1; i++ {
		if b[i][0] != '|' ||
			b[i][len(b[0])-1] != '|' {
			err = true
		}
	}
	//horizontal borders
	for j := 1; j < len(b[0])-1; j++ {
		if b[0][j] != '-' ||
			b[len(b)-1][j] != '-' {
			err = true
		}
	}

	if err {
		return errors.New("invalid board border")
	}
	return nil
}
