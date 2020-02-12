package chapter3

import (
	"encoding/json"
)

// Master is 課題5
// とあるマスターデータのテーブルの構造をモデル化したstructを作りました。
// idやnameを変更できないようにカプセル化し、Getterのみ実装しました。
// しかしモデルをjson.Marshal()でシリアライズすると{}となってしまい、id, nameが出力されません。
// １つ関数を実装してシリアライズした結果が次のようになるようにしましょう。
// {"id": 1, "name": "hoge"}
// ヒント https://golang.org/pkg/encoding/json/#Marshal
// >> If an encountered value implements the Marshaler interface and is not a nil pointer, Marshal calls its MarshalJSON method to produce JSON
type Master struct {
	id   int
	name string
}

// ID is ...
func (m Master) ID() int {
	return m.id
}

// Name is ...
func (m Master) Name() string {
	return m.name
}

// MarshalJSON is ...
func (m Master) MarshalJSON() ([]byte, error) {
	// MyMster is struct
	type MyMster struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	return json.Marshal(MyMster{m.ID(), m.Name()})
}
