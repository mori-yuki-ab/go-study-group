package chapter3

import "fmt"

// Dog is struct
type Dog struct{}

// Bark is Dog
func (d Dog) Bark() string {
	return "わんわん"
}

// Cat is struct
type Cat struct{}

// Crow is Cat
func (c Cat) Crow() string {
	return "にゃーにゃ"
}

// Kadai3 is 課題3
// この関数の引数はxの型は不定です。
// 型がDogの場合はBow()を実行した結果
// Catの場合はCrowを実行した結果
// その他の場合はerrorを返却してください。
func Kadai3(x interface{}) (string, error) {
	switch x := x.(type) {
	case Dog:
		return x.Bark(), nil
	case Cat:
		return x.Crow(), nil
	default:
		return "", fmt.Errorf("x type is %T", x)
	}
}
