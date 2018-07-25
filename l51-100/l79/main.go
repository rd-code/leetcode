package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-25 09:51
 **/

func main() {
    board := [][]byte{
        {'A', 'B', 'C', 'E'},
        {'S', 'F', 'C', 'S'},
        {'A', 'D', 'E', 'E'},
    }
    fmt.Println(exist(board, "ABCCED"))
    //fmt.Println(exist(board, "SEE"))
    //fmt.Println(exist(board, "ABCB"))

}

/**

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
 */

func exist(board [][]byte, word string) bool {
    if len(word) == 0 {
        return true
    }
    for i, vs := range board {
        for j, v := range vs {
            if v == word[0] && find(board, i, j, word[1:]) {
                return true
            }
        }
    }
    return false
}

func find(board [][]byte, x, y int, word string) bool {
    if len(word) == 0 {
        return true
    }
    v := board[x][y]
    board[x][y] = 0

    if x != 0 && board[x-1][y] == word[0] && find(board, x-1, y, word[1:]) {
        return true
    }
    if x < len(board)-1 && board[x+1][y] == word[0] && find(board, x+1, y, word[1:]) {
        return true
    }
    if y != 0 && board[x][y-1] == word[0] && find(board, x, y-1, word[1:]) {
        return true
    }
    if y < len(board[x])-1 && board[x][y+1] == word[0] && find(board, x, y+1, word[1:]) {
        return true
    }
    board[x][y] = v

    return false
}
