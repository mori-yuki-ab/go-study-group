package chapter2

import (
	"fmt"
)

/*
Calc is ...
引数のスライスsliceの要素数が
0の場合、0とエラー
2以下の場合、要素を掛け算
3以上の場合、要素を足し算
を返却。正常終了時、errorはnilでよい
*/
func Calc(slice []int) (int, error) {
	/*
		TODO Q1
		ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
		fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
		https://golang.org/pkg/fmt/#Errorf
	*/
	length := len(slice)
	switch {
	case length == 0:
		return 0, fmt.Errorf("Error: %s", "sliceの要素数が0です。")
	case length == 1:
		return slice[0], nil
	case length == 2:
		return slice[0] * slice[1], nil
	case length >= 3:
		var multiply int
		for _, value := range slice {
			multiply += value
		}
		return multiply, nil
	default:
		return 0, nil
	}
}

// Number is struct
type Number struct {
	index int
}

/*
Numbers is ...
構造体Numberを3つの要素数から成るスライスにして返却
3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
*/
func Numbers() []Number {
	// TODO Q2
	slice := make([]Number, 0, 3)
	slice = append(slice, Number{1}, Number{2}, Number{3})
	return slice
}

/*
CalcMap is ...
引数mをforで回し、「値」部分だけの和を返却
キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
*/
func CalcMap(m map[string]int) int {
	// TODO Q3
	var sum int
	for key, value := range m {
		if key != "yon" {
			sum += value
		}
	}
	return sum
}

// Model is struct
type Model struct {
	Value int
}

/*
Add is ...
与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
*/
func Add(models []Model) {
	// TODO  Q4
	for key, value := range models {
		models[key] = Model{value.Value + 5}
	}
}

/*
Unique is ...
引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
順序はスライスに格納されている順番のまま返却すること
ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
*/
func Unique(slice []int) []int {
	// TODO Q5
	results := make([]int, 0, len(slice))
	encountered := make(map[int]bool)
	for i := 0; i < len(slice); i++ {
		if !encountered[slice[i]] {
			encountered[slice[i]] = true
			results = append(results, slice[i])
		}
	}
	return results
}

/*
Fibonacci is ...
連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
*/
func Fibonacci() func() int {
	var n int
	return func() int {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			x, y = y, x+y
		}
		n++
		return x
	}
}
