package main

import (
    "fmt"
    "strings"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-12 18:30
 **/

func main() {
    nums := []int{1, 2, 1}
    items := permuteUnique(nums)
    for _, item := range items {
        fmt.Println(item)
    }
}

func permuteUnique(nums []int) [][]int {
    if len(nums) < 2 {
        return [][]int{nums}
    }
    if len(nums) == 2 {
        if nums[0] == nums[1] {
            return [][]int{nums}
        }
        return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
    }

    items := permuteUnique(nums[1:])

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

    data := make(map[string]int, len(res))
    for i, array := range res {
        sb := strings.Builder{}
        for _, v := range array {
            sb.WriteRune(rune(v))
            sb.WriteString(",")
        }
        data[sb.String()] = i
    }
    items = make([][]int, 0, len(data))
    for _, v := range data {
        items = append(items, res[v])
    }

    return items
}
