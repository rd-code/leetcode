package main

import "fmt"

/**
 * DESCRIPTION: 给予折半查找
 *
 * @author rd
 * @create 2018-07-20 16:30
 **/

func main() {
    matrix := [][]int{
        {1, 3, 5, 7},
        {10, 11, 16, 20},
        {23, 30, 34, 50},
    }
    fmt.Println(searchMatrix(matrix, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    rowNum, columNum := len(matrix), len(matrix[0])
    start, end := 0, rowNum*columNum-1
    var mid int
    for start <= end {
        mid = (start + end) / 2
        v := getValue(matrix, columNum, mid)
        if v == target {
            return true
        } else if v < target {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    return false
}

func getValue(matrix [][]int, columnNum, position int) int {
    row := position / columnNum
    column := position % columnNum
    return matrix[row][column]
}
