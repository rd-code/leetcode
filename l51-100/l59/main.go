package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 10:53
 **/

func main() {
    res := generateMatrix(1)
    for _, v := range res {
        fmt.Println(v)
    }
}
func generateMatrix(n int) [][]int {
    if n == 0 {
        return nil
    }
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }
    count := 1
    for i := 0; i < n; i++ {
        for j := i; j < n-i; j++ {
            res[i][j] = count
            count += 1
        }
        for j := i + 1; j < n-i; j++ {
            res[j][n-i-1] = count
            count += 1
        }
        for j := n - i - 2; j >= i && i < n-i-1; j-- {
            res[n-i-1][j] = count
            count += 1
        }
        for j := n - i - 2; j > i && i < n-i-1; j-- {
            res[j][i] = count
            count += 1
        }
    }

    return res

}
