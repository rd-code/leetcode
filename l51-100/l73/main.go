package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-20 15:02
 **/

func main() {
   /* matrix := [][]int{
        {1, 1, 1},
        {1, 0, 1},
        {1, 1, 1},
    }*/

     matrix := [][]int{
         {0, 1, 2, 0},
         {3, 4, 5, 2},
         {1, 3, 1, 5},
     }
    setZeroes(matrix)
    for _, v := range matrix {
        fmt.Println(v)
    }
}

func setZeroes(matrix [][]int) {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    row, colum := matrix[0][0], matrix[0][0]
    for i := 0; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            colum = 0
        }
    }
    for i := 0; i < len(matrix[0]); i++ {
        if matrix[0][i] == 0 {
            row = 0
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                matrix[0][j], matrix[i][0] = 0, 0
            }
        }
    }
    fmt.Println(row, colum)

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[i]); j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if row == 0 {
        for i := 0; i < len(matrix[0]); i++ {
            matrix[0][i] = 0
        }
    }
    if colum == 0 {
        for i := 0; i < len(matrix); i++ {
            matrix[i][0] = 0
        }
    }

}
