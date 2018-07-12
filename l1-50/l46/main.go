package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-12 18:02
 **/

func main() {
    nums := []int{1, 2, 3}
    items := permute(nums)
    for _, item := range items {
        fmt.Println(item)
    }
}

func permute(nums []int) [][]int {
    if len(nums) < 2 {
        return [][]int{nums}
    }
    if len(nums) == 2 {
        return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
    }

    items := permute(nums[1:])

    var res [][]int
    for _, item := range items {
        for i := 0; i < len(item)+1; i++ {
            tmp := make([]int, 0, len(item)+1)
            tmp = append(tmp, item[:i]...)
            tmp = append(tmp, nums[0])
            tmp = append(tmp, item[i:]...)
            res = append(res, tmp)
        }
    }

    return res
}
