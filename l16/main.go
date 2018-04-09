package main

func main() {

}

/**
Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

    For example, given array S = {-1 2 1 -4}, and target = 1.

    The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

 */
func threeSumClosest(nums []int, target int) int {
    if len(nums) < 3 {
        return 0
    }
    f := func(a, b int) int {
        if a > b {
            return a - b
        }
        return b - a
    }
    var result int
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                tmp := nums[i] + nums[j] + nums[k]
                t1 := f(target, result)
                t2 := f(target, tmp)
                if t2 < t1 {
                    result = tmp
                }
            }
        }
    }

    return result
}
