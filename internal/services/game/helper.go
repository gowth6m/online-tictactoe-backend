package game

import (
	"fmt"
)

var directions = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
}

/**
 * Depth First Search (DFS) algorithm to find the winning path.
 * @param board: The game board.
 * @param row: The current row.
 * @param col: The current column.
 * @param player: The player to check.
 * @param direction: The direction to move.
 * @param count: The current count.
 * @param winningCondition: The winning condition.
 * @param seen: The map to keep track of visited cells.
 * @param path: The current path.
 * @return The winning path.
 */
func DFS(board Board, row, col int, player string, direction [2]int, count, winningCondition int, seen map[string]bool, path []Cell) []Cell {
	if count == winningCondition {
		return path
	}

	newRow := row + direction[0]
	newCol := col + direction[1]

	if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && board[newRow][newCol].Value != nil && *board[newRow][newCol].Value == player && !seen[fmt.Sprintf("%d-%d", newRow, newCol)] {
		seen[fmt.Sprintf("%d-%d", newRow, newCol)] = true
		newPath := append(path, board[newRow][newCol])
		return DFS(board, newRow, newCol, player, direction, count+1, winningCondition, seen, newPath)
	}

	return nil
}

/**
 * Check if the player is the winner by doing DFS with memorization of visited cells.
 * @param board: The game board.
 * @param player: The player to check.
 * @param winningCondition: The winning condition.
 * @return The winning path.
 */
func IsWinner(board Board, player string, winningCondition int) []Cell {
	size := len(board)
	seen := make(map[string]bool) // Keep track of visited cells (row-col)

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col].Value != nil && *board[row][col].Value == player && !seen[fmt.Sprintf("%d-%d", row, col)] {
				for _, direction := range directions {
					newSeen := make(map[string]bool)
					for k, v := range seen {
						newSeen[k] = v
					}
					newSeen[fmt.Sprintf("%d-%d", row, col)] = true
					path := DFS(board, row, col, player, direction, 1, winningCondition, newSeen, []Cell{board[row][col]})
					if path != nil {
						return path
					}
				}
			}
		}
	}
	return nil
}

/**
 * Check if the game is a draw.
 * @param board: The game board.
 * @return True if the game is a draw, false otherwise.
 */
func IsGameDraw(board Board) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell.Value == nil {
				return false
			}
		}
	}
	return true
}
