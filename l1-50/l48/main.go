package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-11 22:38
 **/
func main() {
    matrics := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    /* matrics := [][]int{
         {5, 1, 9, 11},
         {2, 4, 8, 10},
         {13, 3, 6, 7},
         {15, 14, 12, 16},
     }*/
    rotate(matrics)
    for _, matric := range matrics {
        fmt.Println(matric)
    }
}

func rotate(matrix [][]int) {
    num := len(matrix)

    for i := 0; i < num/2; i++ {
        for j := i; j < num-i-1; j++ {
            matrix[i][j], matrix[j][num-1-i], matrix[num-1-i][num-1-j], matrix[num-1-j][i] =
                matrix[num-1-j][i], matrix[i][j], matrix[j][num-1-i], matrix[num-1-i][num-1-j]
        }

    }

}
