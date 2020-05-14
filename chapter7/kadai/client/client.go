package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/apbgo/go-study-group/chapter7/kadai/model"
)

func main() {
	if err := userFotunr(); err != nil {
		log.Fatal(err)
	}
}

func userFotunr() error {
	// クライアントを作成（タイムアウトを指定）
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	// タイムアウト・キャンセル用のコンテキストを作成
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// リクエストjsonデータの作成
	request := &model.Request{
		UserID: 123456789,
		Name:   "森",
	}
	var body bytes.Buffer
	enc := json.NewEncoder(&body)
	if err := enc.Encode(request); err != nil {
		return err
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", "http://localhost:8080/user_fortune", &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// リクエストにタイムアウトを設定したコンテキストを持たせる
	req = req.WithContext(ctx)

	// リクエスト投げる
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var r model.Response
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&r); err != nil {
		return err
	}

	// 返ってきたレスポンスの内容を表示
	fmt.Println(r.Data)

	return nil
}
