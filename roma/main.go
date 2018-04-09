package main

/**
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.
蒸熟转罗马数字
 */
func intToRoman(num int) string {
    data := map[int]string{
        1:    "I",
        10:   "X",
        100:  "C",
        1000: "M",
        5:    "V",
        50:   "L",
        500:  "D",
    }
    res := ""
    f1 := func(v int) {
        for num >= v {
            res += data[v]
            num -= v
        }
    }
    f2 := func(v1, v2 int) {
        for num >= (v1 - v2) {
            res += data[v2]
            res += data[v1]
            num -= (v1 - v2)
        }
    }
    f1(1000)
    f2(1000, 100)
    f1(500)
    f2(500, 100)
    f1(100)
    f2(100, 10)
    f1(50)
    f2(50, 10)
    f1(10)
    f2(10, 1)
    f1(5)
    f2(5, 1)
    f1(1)

    return res
}

/**
Given a roman numeral, convert it to an integer.
Input is guaranteed to be within the range from 1 to 3999.
 */
//罗马数字转数字
func romanToInt(s string) int {
    data := map[string]int{
        "I": 1,
        "X": 10,
        "C": 100,
        "M": 1000,
        "V": 5,
        "L": 50,
        "D": 500,
    }
    res := 0
    f1 := func(key string) {
        for len(s) > 0 && s[0] == key[0] {
            res = res + data[key]
            s = s[1:]
        }
    }
    f2 := func(k1, k2 string) {
        v := data[k1] - data[k2]
        if len(s) > 1 && s[0] == k2[0] && s[1] == k1[0] {
            res = res + v
            s = s[2:]
        }
    }
    f1("M")
    f2("M", "C")
    f1("D")
    f2("D", "C")
    f1("C")
    f2("C", "X")
    f1("L")
    f2("L", "X")
    f1("X")
    f2("X", "I")
    f1("V")
    f2("V", "I")
    f1("I")
    return res
}
