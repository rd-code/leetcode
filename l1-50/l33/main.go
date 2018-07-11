package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-06-26 16:15
 **/
func main() {
    nums := []int{1, 2, 3, 4, 5, 6}
    fmt.Println(search(nums, 4))
}

func search(nums []int, target int) int {
    if len(nums) < 1 {
        return -1
    }
    if len(nums) < 3 {
        for i, num := range nums {
            if num == target {
                return i
            }
        }
    }
    start, end := 0, len(nums)-1
    mid := (start + end) / 2
    for start < end {
        if nums[start] == target {
            return start
        } else if nums[end] == target {
            return end
        } else if nums[mid] == target {
            return mid
        }
        if (end - start) < 2 {
            break
        }
        if target < nums[mid] {
            if target < nums[start] && nums[mid] > nums[start] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        } else {
            if target > nums[end] && nums[mid] < nums[start] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        }
        mid = (start + end) / 2
    }
    return -1
}
