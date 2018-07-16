package _2

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-16 14:59
 **/

func main() {
    items := solveNQueens(4)
    for _, v := range items {
        for _, s := range v {
            fmt.Println(s)
        }
        fmt.Println("==============")
    }
}

func totalNQueens(n int) int {
    if n < 2 {
        return n
    }
    positions := make([][2]int, n) //每个元素代表一行，记录该行皇后的位置
    for i := 0; i < n; i++ {
        positions[i] = [2]int{-1, -1}
    }
    res := make([][][2]int, 0, n)

    fill(positions, 0, &res)

    items := make([][]string, 0, len(res))
    for _, rect := range res {
        item := make([]string, 0, n)
        for _, v := range rect {
            tmp := make([]byte, 0, n)
            for j := 0; j < n; j++ {
                if v[1] == j {
                    tmp = append(tmp, 'Q')
                } else {
                    tmp = append(tmp, '.')
                }
            }
            item = append(item, string(tmp))
        }
        items = append(items, item)
    }
    return len(items)
}

func cp(positions [][2]int) [][2]int {
    res := make([][2]int, len(positions))
    for i, p := range positions {
        res[i] = [2]int{p[0], p[1]}
    }
    return res
}

func fill(positions [][2]int, n int, res *[][][2]int) {
    if n == len(positions) {
        *res = append(*res, positions)
        return
    }
    columns := getColumns(positions, n)
    if len(columns) == 0 {
        return
    }
    for _, column := range columns {
        p := cp(positions)
        p[n] = [2]int{n, column}
        fill(p, n+1, res)
    }
}

func getColumns(flags [][2]int, n int) []int {
    columns := make(map[int]struct{})
    size := len(flags)
    for i := 0; i < n; i++ {
        flag := flags[i]
        row, column := flag[0], flag[1]
        if row == -1 || column == -1 {
            continue
        }
        columns[column] = struct{}{}
        row = n - row
        column = column + row
        if column < size {
            columns[column] = struct{}{}
        }
        row, column = flag[0], flag[1]
        row = n - row
        column = column - row
        if column >= 0 {
            columns[column] = struct{}{}
        }

    }
    var res []int
    for i := 0; i < size; i++ {
        if _, ok := columns[i]; !ok {
            res = append(res, i)
        }
    }
    return res
}
