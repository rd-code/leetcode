package main

import (
    "strings"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-19 11:44
 **/
func main() {

    /*words := []string{"This", "is", "an", "example", "of", "text", "justification."}

    items := fullJustify(words, 16)*/

    /* words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}

     items := fullJustify(words, 16)*/

    words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
        "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}

    items := fullJustify(words, 20)

    for _, item := range items {
        fmt.Println(item)
    }

}

func fullJustify(words []string, maxWidth int) []string {
    if len(words) == 0 {
        return words
    }
    count := len(words[0]) + 1
    start := 0
    var res []string
    for i := 1; i < len(words); i++ {
        if count+len(words[i]) > maxWidth {
            res = append(res, compose(words[start:i], maxWidth))
            count = 0
            start = i
        }
        count += len(words[i]) + 1

    }
    t := compose(words[start:], maxWidth)
    array := make([]byte, 0, maxWidth)
    array = append(array, t[0])

    for i := 1; i < maxWidth; i++ {
        if array[len(array)-1] == ' ' && t[i] == ' ' {
            continue
        }
        array = append(array, t[i])
    }
    for i := len(array); i < maxWidth; i++ {
        array = append(array, ' ')
    }

    res = append(res, string(array))
    return res
}

func compose(words []string, maxWidth int) string {
    sb := strings.Builder{}
    if len(words) < 2 {
        sb.WriteString(words[0])
        for i := 0; i < maxWidth-len(words[0]); i++ {
            sb.WriteByte(' ')
        }
        return sb.String()
    }
    var allSize int
    for _, v := range words {
        allSize += len(v)
    }
    allSize = maxWidth - allSize
    avgSize := allSize / (len(words) - 1)
    levalSize := allSize % (len(words) - 1)
    for i := 0; i < len(words)-1; i++ {
        sb.WriteString(words[i])
        for i := 0; i < avgSize; i++ {
            sb.WriteByte(' ')
        }
        if levalSize > 0 {
            levalSize -= 1
            sb.WriteByte(' ')
        }
    }
    sb.WriteString(words[len(words)-1])
    return sb.String()
}
