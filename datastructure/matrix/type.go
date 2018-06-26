package main

import (
    "strconv"
    "strings"
    "fmt"
    "time"
)

type Matrix struct {
    Data      [][]int
    RowNum    int
    ColumnNum int
}

func New(rowNum, columnNum int) *Matrix {
    res := &Matrix{
        Data:      make([][]int, rowNum),
        RowNum:    rowNum,
        ColumnNum: columnNum,
    }
    for i := 0; i < rowNum; i++ {
        res.Data[i] = make([]int, columnNum)
    }
    return res
}

func (m *Matrix) String() string {
    res := make([]string, m.RowNum)
    for i := 0; i < m.RowNum; i++ {
        items := make([]string, m.ColumnNum)
        for j := 0; j < m.ColumnNum; j++ {
            items[j] = strconv.Itoa(m.Data[i][j])
        }
        res[i] = strings.Join(items, ", ")

    }
    return strings.Join(res, "\n")
}

//复制某个区域
func (m *Matrix) Copy(rowStart, rowNum, columnStart, columnNum int, desc *Matrix) {
    for i := 0; i < rowNum; i++ {
        for j := 0; j < columnNum; j++ {
            desc.Data[i][j] = m.Data[i+rowStart][j+columnStart]
        }
    }
}

func (m *Matrix) Add(data *Matrix) (*Matrix) {
    res := New(m.RowNum, m.ColumnNum)
    for i := 0; i < m.RowNum; i++ {
        for j := 0; j < m.ColumnNum; j++ {
            res.Data[i][j] = m.Data[i][j] + data.Data[i][j]
        }
    }
    return res
}

func Add(d1, d2 *Matrix) *Matrix {
    return d1.Add(d2)
}

func Mul(l, r *Matrix) *Matrix {
    if l.RowNum == 1 {
        res := New(1, 1)
        res.Data[0][0] = l.Data[0][0] * r.Data[0][0]
        return res
    }
    num := l.RowNum / 2

    left11 := New(num, num)
    l.Copy(0, num, 0, num, left11)

    left12 := New(num, num)
    l.Copy(0, num, num, num, left12)

    left21 := New(num, num)
    l.Copy(num, num, 0, num, left21)

    left22 := New(num, num)
    l.Copy(num, num, num, num, left22)

    right11 := New(num, num)
    r.Copy(0, num, 0, num, right11)

    right12 := New(num, num)
    r.Copy(0, num, num, num, right12)

    right21 := New(num, num)
    r.Copy(num, num, 0, num, right21)

    right22 := New(num, num)
    r.Copy(num, num, num, num, right22)

    c11 := Add(Mul(left11, right11), Mul(left12, right21))

    c12 := Add(Mul(left11, right12), Mul(left12, right22))

    c21 := Add(Mul(left21, right11), Mul(left22, right21))

    c22 := Add(Mul(left21, right12), Mul(left22, right22))

    res := New(l.RowNum, l.RowNum)

    for i := 0; i < num; i++ {
        for j := 0; j < num; j++ {
            res.Data[i][j] = c11.Data[i][j]
        }
    }
    for i := 0; i < num; i++ {
        for j := 0; j < num; j++ {
            res.Data[i][j+num] = c12.Data[i][j]
        }
    }
    for i := 0; i < num; i++ {
        for j := 0; j < num; j++ {
            res.Data[i+num][j] = c21.Data[i][j]
        }
    }
    for i := 0; i < num; i++ {
        for j := 0; j < num; j++ {
            res.Data[i+num][j+num] = c22.Data[i][j]
        }
    }

    return res
}

func main() {
    data1 := New(8, 8)
    data2 := New(8, 8)

    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            data1.Data[i][j] = i*10 + j
            data2.Data[i][j] = 100 + i*10 + j
        }
    }
    fmt.Println(data1, "\n")
    fmt.Println(data2, "\n")

    time.Sleep(time.Second)

    fmt.Println(Mul(data1, data2))

}
