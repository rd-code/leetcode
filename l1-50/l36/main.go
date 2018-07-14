package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-13 17:47
 **/

func main() {
    board := [][]string{
        {".", ".", ".", ".", ".", ".", "5", ".", "."},
        {".", ".", ".", ".", ".", ".", ".", ".", "."},
        {".", ".", ".", ".", ".", ".", ".", ".", "."},
        {"9", "3", ".", ".", "2", ".", "4", ".", "."},
        {".", ".", "7", ".", ".", ".", "3", ".", "."},
        {".", ".", ".", ".", ".", ".", ".", ".", "."},
        {".", ".", ".", "3", "4", ".", ".", ".", "."},
        {".", ".", ".", ".", ".", "3", ".", ".", "."},
        {".", ".", ".", ".", ".", "5", "2", ".", "."},
    }

    var items [][]byte
    for _, b := range board {
        var item []byte
        for _, s := range b {
            item = append(item, byte(s[0]))
        }
        items = append(items, item)
    }
    for _, item := range items {
        fmt.Println(item)
    }
    fmt.Println(isValidSudoku(items))

}

func isValidSudoku(board [][]byte) bool {
    data1, data2 := map[byte]struct{}{}, map[byte]struct{}{}
    var ok bool
    var v byte
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board); j++ {
            v = board[i][j]
            if v == 46 {

            } else if _, ok = data1[v]; ok {
                return false
            } else {
                data1[v] = struct{}{}
            }

            v = board[j][i]
            if v == 46 {
                continue
            }
            if _, ok = data2[v]; ok {
                return false
            } else {
                data2[v] = struct{}{}
            }
        }
        for v = range data1 {
            delete(data1, v)
        }
        for v = range data2 {
            delete(data2, v)
        }
    }
    for i := 0; i < len(board)-2; i += 3 {
        for j := 0; j < len(board)-2; j += 3 {
            for row := 0; row < 3; row++ {
                for column := 0; column < 3; column++ {
                    v = board[i+row][j+column]
                    if v == 46 {
                        continue
                    } else if _, ok = data1[v]; ok {
                        return false
                    } else {
                        data1[v] = struct{}{}
                    }
                }
            }
            for v = range data1 {
                delete(data1, v)
            }
        }
    }

    return true
}
