package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-14 11:56
 **/
/**
1:49 2:50 3:51 4:52 5:53 6:54 7:55 8:56 9:57
 */
func main() {
    board := [][]string{
        {"5", "3", ".", ".", "7", ".", ".", ".", "."},
        {"6", ".", ".", "1", "9", "5", ".", ".", "."},
        {".", "9", "8", ".", ".", ".", ".", "6", "."},
        {"8", ".", ".", ".", "6", ".", ".", ".", "3"},
        {"4", ".", ".", "8", ".", "3", ".", ".", "1"},
        {"7", ".", ".", ".", "2", ".", ".", ".", "6"},
        {".", "6", ".", ".", ".", ".", "2", "8", "."},
        {".", ".", ".", "4", "1", "9", ".", ".", "5"},
        {".", ".", ".", ".", "8", ".", ".", "7", "9"},
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
    solveSudoku(items)
    fmt.Println("---------")
    for _, item := range items {
        fmt.Println(item)
    }

    fmt.Println(isValidSudoku(items))

}

func solveSudoku(board [][]byte) {
    solveSodokuImpl(board)
}

func solveSodokuImpl(board [][]byte) bool {
    var row, column int
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board); j++ {
            if board[i][j] == 46 {
                row = i
                column = j
                goto Lable1
            }
        }
    }
Lable1:
    if board[row][column] != 46 {
        return true
    }

    data := map[byte]struct{}{}
    //校验行
    for i := 0; i < len(board); i++ {
        v := board[row][i]
        if v != 46 {
            data[v] = struct{}{}
        }
        v = board[i][column]
        if v != 46 {
            data[v] = struct{}{}
        }
    }

    for i := row / 3 * 3; i < row/3*3+3; i++ {
        for j := column / 3 * 3; j < column/3*3+3; j++ {
            v := board[i][j]
            if v != 46 {
                data[v] = struct{}{}
            }
        }
    }
    if len(data) == len(board) {
        return false
    }

    var i uint8
    var ok bool
    for i = 49; i < 58; i++ {
        if _, ok = data[i]; !ok {
            board[row][column] = i
            if solveSodokuImpl(board) {
                return true
            }
        }
    }
    board[row][column] = 46
    return false
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
