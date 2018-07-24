package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-23 18:29
 **/

func main() {
    fmt.Println(minWindow("ADOBECODEBANC", "ABC"))

}
func minWindow(s string, t string) string {
    requirements := make(map[byte]int, len(t))
    for i := range t {
        v := t[i]
        requirements[v] += 1
    }

    appearances := make(map[byte]int, len(t))
    counter := len(t)
    minLen := len(s) + 1
    resLeft, resRight := 0, 0
    for left, right := 0, 0; right < len(s); right++ {
        lch, rch := s[left], s[right]
        if _, ok := requirements[rch]; ok {
            appearances[rch] += 1
            if appearances[rch] <= requirements[rch] && counter != 0 {
                counter--
            }
        }
        if counter == 0 {
            for {
                if _, ok := requirements[lch]; !ok || appearances[lch] > requirements[lch] {
                    if appearances[lch] > requirements[lch] {
                        appearances[lch]--
                    }
                    left++
                    lch = s[left]
                } else {
                    break
                }
            }
            if right-left+1 < minLen {
                resLeft, resRight = left, right+1
                minLen = right - left + 1
            }
        }

    }

    return s[resLeft:resRight]
}
