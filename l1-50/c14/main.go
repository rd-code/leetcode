package main

func main() {

}
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    s := strs[0]
    var i int
    for i = 0; i < len(s); i++ {
        for _, str := range strs {
            if len(str) == i || str[i] != s[i] {
                goto Label
            }
        }
    }
Label:

    return s[:i]
}
