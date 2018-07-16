package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-16 17:41
 **/

func main() {
    /* matrix := [][]int{
         {1, 2, 3},
         {4, 5, 6},
         {7, 8, 9},
     }*/
    /* matrix := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }*/
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
        {10, 11, 12},
        {13, 14, 15},
    }
    //matrix := [][]int{{6, 9, 7}}
   // matrix := [][]int{{6}, {9}, {7}}
    //matrix := [][]int{{1}, {2}, {6}, {9}, {7}}
    item := spiralOrder(matrix)
    fmt.Println(item)
}

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }
    row := len(matrix)
    column := len(matrix[0])
    res := make([]int, 0, row*column)

    n := row
    if row > column {
        n = column
    }
    n = (n + 1) / 2

    for i := 0; i < n && i < row; i++ {
        for j := i; j < column-i; j++ {
            res = append(res, matrix[i][j])
        }
        for j := i + 1; j < row-i; j++ {
            res = append(res, matrix[j][column-i-1])
        }
        for j := column - i - 2; j >= i && i < row-i-1; j-- {
            res = append(res, matrix[row-i-1][j])
        }
        for j := row - i - 2; j > i && i < column-i-1; j-- {
            res = append(res, matrix[j][i])
        }
    }

    return res
}
