// Valid Sudoku
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// 53- -7- ---
// 6-- 195 ---
// -98 --- -6-
// 8-- -6- --3
// 4-- 8-3 --1
// 7-- -2- --6
// -6- --- 28-
// --- 419 -5-
// --- -8- -79

// Example 1:
// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true

// Example 2:
// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1,
// except with the 5 in the top left corner being modified to 8.
// Since there are two 8's in the top left 3x3 sub-box, it is invalid.

package medium

import "fmt"

func isValidSudoku(board [][]byte) bool {

	mapRow := make(map[string]bool, 81)
	mapCol := make(map[string]bool, 81)
	mapSqr := make(map[string]bool, 270)

	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 46 {
				continue
			}

			cellRow := fmt.Sprintf("%v-%v", i, board[i][j])

			// check row
			if mapRow[cellRow] {
				return false
			}
			// add to row
			mapRow[cellRow] = true

			cellColumn := fmt.Sprintf("%v-%v", j, board[i][j])
			// check col

			if mapCol[cellColumn] {
				return false
			}

			// add to col
			mapCol[cellColumn] = true

			// in which squre is ?
			cellSquare := fmt.Sprintf("%v-%v-%v", i/3, j/3, board[i][j])

			if mapSqr[cellSquare] {
				return false
			}

			// add to square
			mapSqr[cellSquare] = true

		}

	}

	return true

}
