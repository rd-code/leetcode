package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-11 22:10
 **/
func main() {
    fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return nil
    }
    var items [][]string
    positions := []map[rune]int{}
    for _, str := range strs {
        tmp := map[rune]int{}
        for _, s := range str {
            tmp[s] += 1
        }
        var i = 0
        for i = 0; i < len(positions); i++ {
            position := positions[i]
            if len(position)!= len(tmp){
                goto Lable1
            }
            for k, v := range tmp {
                if position[k] != v {
                    goto Lable1
                }
            }
            goto Lable2
        Lable1:
        }
    Lable2:
        if i == len(positions) {
            positions = append(positions, tmp)
            items = append(items, []string{str})
        } else {
            items[i] = append(items[i], str)
        }

    }
    return items
}
