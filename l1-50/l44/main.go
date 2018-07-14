package main

import (
    "fmt"
)

//使用二维数组记录上一次匹配的信息，然后根据上一次匹配的情况处理当前的情况
/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-12 22:14
 **/

func main() {
    //s := "abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb"
    //p := "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"
    //s := "ab"
    //p := "*?*?*"
    // s := ""
    //p := "*a*"
    // s := "adceb"
    //p := "*a*b"
    // true
    /* s := "a"
     p := "a*"*/
    // s = "babbba"
    //p = "*a***a"
    /* //false
     s := "b"
     p := "?*?"*/
    //fmt.Println(isMatch(s, p))
}
func isMatch(s string, p string) bool {
    p = formatPattern(p)
    data := make([][]bool, len(p))
    for i := 0; i < len(p); i++ {
        data[i] = make([]bool, len(s)+1)
    }

    if len(p) == 0 {
        return len(s) == 0
    }
    if len(s) == 0 {
        return p[0] == '*' && len(p) == 1
    }
    if len(p) == 1 && (p[0] == '*') {
        return true
    }
    if len(s) == 1 {
        if len(p) == 1 && (p[0] == s[0] || p[0] == '?') {
            return true
        }
    }
    if p[0] == '*' {
        for i := 0; i < len(s); i++ {
            data[0][i] = true
        }
    }
    if p[0] == '?' || p[0] == s[0] {
        data[0][0] = true
    }

    _ = match(s, p, data, 0)
    for _, d := range data {
        fmt.Println(d)
    }
    return match(s, p, data, 0)
}

func formatPattern(p string) string {
    s := make([]byte, 0, len(p))
    if len(p) > 0 {
        s = append(s, p[0])
        for i := 1; i < len(p); i++ {
            if s[len(s)-1] == '*' && p[i] == '*' {
                continue
            }
            s = append(s, p[i])
        }
    }

    return string(s)

}

func match(s string, p string, data [][]bool, position int) bool {
    if position == len(p)-1 {
        return data[len(p)-1][len(s)-1]
    }
    for i := 0; i < len(s); i++ {
        if data[position][i] {
            goto Lable1
        }
    }
    return false
Lable1:
    for i := 0; i < len(s); i++ {
        if data[position][i] == true {
            if p[position] == '*' { //上一个是型号，特殊处理
                for j := i; j < len(s); j++ {
                    if p[position+1] == '?' || p[position+1] == s[j] {
                        if position < 1 {
                            data[position+1][j] = true
                        } else {
                            for k := 0; k < j; k++ {
                                if data[position-1][k] {
                                    data[position+1][j] = true
                                    break
                                }
                            }
                        }

                    }
                }
            } else if p[position+1] == '*' {
                for j := i; j < len(s); j++ {
                    data[position+1][j] = true
                }
                goto Lable2
            } else if len(s) > (i+1) && (p[position+1] == '?' || p[position+1] == s[i+1]) {
                data[position+1][i+1] = true
            }
        }
    }
Lable2:
    return match(s, p, data, position+1)
}
