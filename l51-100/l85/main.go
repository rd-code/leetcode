package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-11-25 18:12
 **/
func main() {
    items := [][]byte{
        {'1', '0', '1', '0', '0'},
        {'1', '0', '1', '1', '1'},
        {'1', '1', '1', '1', '1'},
        {'1', '0', '0', '1', '0'},
    }
    fmt.Println(maximalRectangle(items))
}

func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    data := make([][]int, len(matrix))
    for i := 0; i < len(data); i++ {
        data[i] = make([]int, len(matrix[0]))
    }
    var max int
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            max = cal(matrix, data, i, j)
        }
    }
    return max
}

func cal(matrix [][]byte, data [][]int, x, y int) int {
    var max = 0
    var cMin int
    for row := x; row >= 0 && data[row][y] != '0'; row-- {
        for column := y; column >= cMin; column-- {
            if matrix[row][column] == '0' {
                cMin = column + 1
                break
            }
            tmp := (x - row + 1) * (y - column + 1)
            if max < tmp {
                max = tmp
            }
        }
    }
    if x > 0 && max < data[x-1][y] {
        max = data[x-1][y]
    }
    if y > 0 && max < data[x][y-1] {
        max = data[x][y-1]
    }
    data[x][y] = max
    return max
}

/**
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
 */
