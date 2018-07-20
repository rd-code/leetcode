package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-20 17:28
 **/
func main() {
    //nums := []int{2, 0, 2, 1, 1, 0}
    //nums := []int{1, 0}
    //nums := []int{1, 0, 0}
    //nums := []int{2, 1, 2}
       nums := []int{1, 2, 1}
    sortColors(nums)
    fmt.Println(nums)

}
func sortColors(nums []int) {
    left := 0
    right := len(nums) - 1
    for i := 0; i < right+1; i++ {
        v := nums[i]
        if v == 0 {
            nums[left], nums[i] = nums[i], nums[left]
            left++
        }
        if v == 2 {
            nums[right], nums[i] = nums[i], nums[right]
            i--
            right--
        }
    }
}
