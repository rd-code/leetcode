package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 19:45
 **/

func main() {
    data := [][]int{
        {1, 3, 1},
        {1, 5, 1},
        {4, 2, 1},
    }
    fmt.Println(minPathSum(data))
}

func minPathSum(grid [][]int) int {
    if len(grid) == 1 {
        sum := 0
        for _, v := range grid[0] {
            sum += v
        }
        return sum
    }
    if len(grid[0]) == 1 {
        sum := 0
        for _, v := range grid {
            sum += v[0]
        }
        return sum
    }
    data := make([][]int, len(grid))
    for i, v := range grid {
        items := make([]int, len(v))
        for j := range items {
            items[j] = -1
        }
        data[i] = items
    }
    return walk(grid, data, 0, 0)
}

func walk(grid, data [][]int, row, column int) int {
    if data[row][column] != -1 {
        return data[row][column]
    }
    if row == len(grid)-1 && column == len(grid[row])-1 {
        data[row][column] = grid[row][column]
        return grid[row][column]
    }

    min := -1
    if row < len(grid)-1 {
        v1 := grid[row][column] + walk(grid, data, row+1, column)
        min = v1
    }
    if column < len(grid[row])-1 {
        v2 := grid[row][column] + walk(grid, data, row, column+1)
        if min == -1 || min > v2 {
            min = v2
        }
    }
    data[row][column] = min

    return min
}
