package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-25 12:37
 **/

func main() {
    // nums := []int{1, 1, 1, 2, 2, 3}
    nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
    size := removeDuplicates(nums)
    fmt.Println(size)
    fmt.Println(nums[:size])
}

func removeDuplicates(nums []int) int {
    if len(nums) < 3 {
        return len(nums)
    }
    value, position, count := nums[0], 0, 1
    for i := 1; i < len(nums); i++ {
        v := nums[i]
        if v == value && count == 2 {
            continue
        }
        position++
        nums[position], nums[i] = nums[i], nums[position]
        if v == value {
            count++
        } else {
            count = 1
            value = nums[position]
        }

    }
    return position + 1
}
