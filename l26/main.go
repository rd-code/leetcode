/**
Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the new length.
 */
package main

import "fmt"

func main() {
    t := []int{-3, -1, -1, 0, 0, 0, 0, 0, 2}
    removeDuplicates(t)
    fmt.Println(t)
}

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[count] != nums[i] {
            count++
            nums[count] = nums[i]
        }
    }
    return count + 1
}
