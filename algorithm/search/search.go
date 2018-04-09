package main

import "fmt"

func main() {
    a := []int{1, 2}
    fmt.Println(binSearch(a, 0))
}

//二分查找
func binSearch(nums []int, v int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == v {
            return mid
        }
        if nums[mid] < v {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
