package main

import "sort"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-16 21:16
 **/

func main() {

}

type Interval struct {
    Start int
    End   int
}

func merge(intervals []Interval) []Interval {
    if len(intervals) == 0 {
        return nil
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    res := make([]Interval, 0, len(intervals))
    res = append(res, intervals[0])
    position := 0
    for i := 1; i < len(intervals); i++ {
        if res[position].Start <= intervals[i].Start && res[position].End >= intervals[i].Start ||
            res[position].Start <= intervals[i].End && res[position].End >= intervals[i].End ||
            res[position].Start >= intervals[i].Start && res[position].Start <= intervals[i].End ||
            res[position].End >= intervals[i].Start && res[position].End <= intervals[i].End {
            if res[position].End < intervals[i].End {
                res[position].End = intervals[i].End
            }
            if res[position].Start > intervals[i].Start {
                res[position].Start = intervals[i].Start
            }
        } else {
            res = append(res, intervals[i])
            position += 1
        }
    }
    if len(res) == len(intervals) {
        return res
    }
    return merge(res)
}
