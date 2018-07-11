package main

func main() {

}

func strStr(haystack string, needle string) int {
    if len(haystack) == 0 && len(needle) == 0 {
        return 0
    }
    if len(haystack) == 0 || len(needle) > len(haystack) {
        return -1
    }
    for i := 0; i < len(haystack)-len(needle)+1; i++ {
        if needle == haystack[i:i+len(needle)] {
            return i
        }
    }
    return -1
}
