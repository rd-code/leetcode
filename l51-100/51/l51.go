package main

import (
    "fmt"
)

/**
 * DESCRIPTION: n皇后问题，解题思路为使用一个array记录每行皇后的位置，然后计算出来目标行可以使用的位置,对某些算法进行了优化
 *
 * @author rd
 * @create 2018-07-15 19:08
 **/

func main() {
    items := solveNQueens(15)
    for _, v := range items {
        for _, s := range v {
            fmt.Println(s)
        }
        fmt.Println("==============")
    }
}

func solveNQueens(n int) [][]string {
    if n < 1 {
        return nil
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
    return items
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
        *res = append(*res, cp(positions))
        return
    }
    columns := getColumns(positions, n)
    if len(columns) == 0 {
        return
    }
    for _, column := range columns {
        positions[n] = [2]int{n, column}
        fill(positions, n+1, res)
        positions[n] = [2]int{-1, -1}
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
