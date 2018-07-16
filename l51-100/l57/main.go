package l57

import "sort"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-16 22:34
 **/

func main() {

}

type Interval struct {
    Start int
    End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
    intervals = append(intervals, newInterval)
    return merge(intervals)
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
