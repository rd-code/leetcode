package main

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    data := map[string][]string{
        "2": {"a", "b", "c"},
        "3": {"d", "e", "f"},
        "4": {"g", "h", "i"},
        "5": {"j", "k", "l"},
        "6": {"m", "n", "o"},
        "7": {"p", "q", "r", "s"},
        "8": {"t", "u", "v"},
        "9": {"w", "x", "y", "z"},
    }

    var items [][]string
    {
        tmp := []rune(digits)
        items = make([][]string, 0, len(tmp))
        for _, t := range tmp {
            items = append(items, data[string(t)])
        }
    }

    return merge(items)
}

func merge(items [][]string) []string {
    if len(items) < 2 {
        return items[0]
    }
    list := items[1:]
    item := items[0]

    var result []string
    for _, s := range item {
        tmp := merge(list)
        for _, ss := range tmp {
            result = append(result, s+ss)
        }
    }
    return result
}
