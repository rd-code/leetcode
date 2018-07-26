package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-26 17:02
 **/
func main() {
    //nums = [2,5,6,0,0,1,2], target = 0

    nums := []int{2, 5, 6, 0, 0, 1, 2}
    fmt.Println(search(nums, 0))
    fmt.Println(search(nums, 3))
    nums = []int{1, 3, 5}
    fmt.Println(search(nums, 5))
    nums = []int{1, 1, 3, 1}
    fmt.Println(search(nums, 3))
    nums = []int{1, 3, 1, 1, 1}
    fmt.Println(search(nums, 3))
    nums = []int{0, 0, 1, 1, 2, 0}
    fmt.Println(search(nums, 2))

}
func search(nums []int, target int) bool {
    start := 0
    end := len(nums) - 1
    for start <= end {
        mid := (start + end) / 2
        v := nums[mid]
        if v == target  || target==nums[start]||target==nums[end]{
            return true
        }
        if v < target {
            if nums[start] == nums[end] {
                if target == nums[start] {
                    return true
                }
                start = start + 1
                continue
            }
            if target > nums[end] && v <= nums[end] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        } else {
            if nums[start] == nums[end] {
                if target == nums[start] {
                    return true
                }
                start += 1
                continue
            }
            if v > nums[end] && target < nums[end] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
    }
    return false
}
