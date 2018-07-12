package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-12 18:43
 **/

func main() {
    nums := []int{2, 3, 1, 1, 4}
    fmt.Println(jump(nums))
}

func jump(nums []int) int {
    steps := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        steps[i] = -1
    }
    n := jumps(nums, steps)
    return n
}

func jumps(nums []int, steps []int) int {
    if len(nums) <= 1 {
        if len(nums) == 1 {
            steps[0] = 0
        }
        return 0
    }
    if nums[0]+1 >= len(nums) {
        if steps[0] == -1 {
            steps[0] = 1
        }
        return 1
    }

    min := 1 << 30
    for i := 0; i < nums[0]; i++ {
        if steps[i+1] != -1 {
            if steps[i+1] < min {
                min = steps[i+1]
            }
        } else {
            t := jumps(nums[i+1:], steps[i+1:])
            if t < min {
                min = t
            }
        }
    }

    steps[0] = min + 1
    return min + 1
}
