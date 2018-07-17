package main

import (
    "sort"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 15:32
 **/

func main() {
    res := getPermutation(4, 9)
    fmt.Println(res)
}

func getPermutation(n int, k int) string {
    if n == 1 {
        return "1"
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    array := make([]byte, n)
    if k == 1 {
        for i, n := range nums {
            array[i] = byte(n) + '0'
        }
        return string(array)
    }

    generatePermutation(nums, k, 1)
    for i, n := range nums {
        array[i] = byte(n) + '0'
    }
    return string(array)
}

func generatePermutation(nums []int, sum, k int) {
    if k == sum {
        return
    }
    nextPermutation(nums)
    generatePermutation(nums, sum, k+1)

}

func nextPermutation(nums []int) {
    if len(nums) < 2 {
        return
    }
    generaeGreater(nums)
}

func generaeGreater(nums []int) bool {
    if len(nums) == 2 {
        if nums[0] < nums[1] {
            nums[0], nums[1] = nums[1], nums[0]
            return true
        } else {
            return false
        }
    }
    if generaeGreater(nums[1:]) {
        return true
    }
    min := 1
    for i := 1; i < len(nums); i++ {
        if nums[min] > nums[i] && nums[i] > nums[0] {
            min = i
        }
    }
    if nums[0] < nums[min] {
        nums[0], nums[min] = nums[min], nums[0]
        sort.Ints(nums[1:])
        return true
    }

    return false
}
