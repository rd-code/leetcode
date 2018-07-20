package main

import (
    "strings"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-19 19:11
 **/
func main() {
    path := "/home//foo/"
    fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
    if len(path) < 2 {
        return "/"
    }
    var link []string
    link = append(link, "/")
    var end int
    for i := 1; i < len(path); i++ {
        if path[i] == '/' {
            continue
        }
        for end = i; end < len(path) && path[end] != '/'; end++ {
        }
        value := path[i:end]
        if value == "." {
            continue
        }
        if value == ".." {
            if len(link) == 1 {
                continue
            }
            link = link[:len(link)-1]
        } else {
            link = append(link, value)
        }
        if end == len(path) {
            break
        }
        i = end
    }
    if len(link) == 1 {
        return "/"
    }
    sb := &strings.Builder{}
    for i := 1; i < len(link); i++ {
        sb.WriteByte('/')
        sb.WriteString(link[i])

    }
    return sb.String()
}
