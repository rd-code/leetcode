package main

import (
    "math/rand"
    "time"
    "fmt"
)

func main() {
    rand := rand.New(rand.NewSource(time.Now().UnixNano()))
    var nums []int
    for i := 0; i < 100; i++ {
        nums = append(nums, rand.Intn(100))
    }
    shellSort(nums)
    for i := 0; i < 99; i++ {
        if nums[i] > nums[i+1] {
            fmt.Println("not succeed")
        }
    }
    fmt.Println(len(nums))

    fmt.Println(nums)

}

//冒泡排序
func bubblSort(nums []int) {
    for i := 0; i < len(nums)-1; i++ {
        for j := 0; j < len(nums)-i-1; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}

//选择排序
func selectSort(nums []int) {
    var pmin int
    for i := 0; i < len(nums)-1; i++ {
        pmin = len(nums) - 1
        for j := len(nums) - 2; j >= i; j-- {
            if nums[j] < nums[pmin] {
                pmin = j
            }
        }
        nums[i], nums[pmin] = nums[pmin], nums[i]

    }
}

//插入排序
func insertSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        var index int
        for index = 0; index < i; index++ {
            if nums[index] <= nums[i] {
                continue
            }
            break
        }
        t := nums[i]
        for j := i; j > index; j-- {
            nums[j] = nums[j-1]
        }
        nums[index] = t

    }
}

//快速排序
func quickSort(nums []int) {
    if len(nums) < 3 {
        bubblSort(nums)
        return
    }
    i, j := 0, len(nums)-1
    for i < j {
        for ; i < j; j-- {
            if nums[j] < nums[i] {
                break
            }
        }
        if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
        for ; i < j; i++ {
            if nums[i] > nums[j] {
                break
            }
        }
        if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    quickSort(nums[:i])
    quickSort(nums[i+1:])

}

//归并排序
func mergeSort(nums []int) {
    if len(nums) < 2 {
        return
    }
    mid := len(nums) / 2
    list1 := nums[:mid]
    list2 := nums[mid:]
    mergeSort(list1)
    mergeSort(list2)
    merge(nums, list1, list2)
}

func merge(nums, num1 []int, num2 []int) {
    t := make([]int, len(nums))
    p1, p2 := 0, 0
    for i := 0; i < len(nums); i++ {
        if p1 >= len(num1) {
            t[i] = num2[p2]
            p2++
        } else if p2 >= len(num2) {
            t[i] = num1[p1]
            p1++
        } else if num1[p1] < num2[p2] {
            t[i] = num1[p1]
            p1++
        } else {
            t[i] = num2[p2]
            p2++
        }
    }
    for i := range nums {
        nums[i] = t[i]
    }
}

//希尔排序
func shellSort(nums []int) {
    gap := 5
    for gap > 0 {
        for i := 0; i < gap; i++ {
            for j := i + gap; j < len(nums); j += gap {
                t := nums[j]
                var k int
                for k = i; k < j; k += gap {
                    if nums[j] < nums[k] {
                        break
                    }
                }
                for p := j; p > k; p -= gap {
                    nums[p] = nums[p-gap]
                }
                nums[k] = t
            }
        }
        gap /= 2
    }
}
