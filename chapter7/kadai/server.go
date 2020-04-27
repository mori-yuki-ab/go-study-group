package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apbgo/go-study-group/chapter7/kadai/model"
)

// 処理ハンドラ
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, server.")
}

// 処理ハンドラ
func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("p") == "cheat" {
		fmt.Fprint(w, "大吉")
	} else {
		fmt.Fprint(w, DrawFortune())
	}
}

func userFortuneHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディの取得
	defer r.Body.Close()
	var req model.Request
	// リクエストパラメタ取得
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := fmt.Sprintf("ID:%dの%sさんの運勢は%sです！", req.UserID, req.Name, DrawFortune())
	response := &model.Response{
		Status: http.StatusOK,
		Data:   data,
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(response); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, buf.String())
}

func main() {
	mux := http.NewServeMux()
	// ハンドラをエントリポイントと紐付け
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/fortune", fortuneHandler)
	mux.HandleFunc("/user_fortune", userFortuneHandler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// OSからのシグナルを待つ
	go func() {
		// SIGTERM: コンテナが終了する時に送信されるシグナル
		// SIGINT: Ctrl+c
		sigCh := make(chan os.Signal, 1)
		// 受け取るシグナルを指定
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		// チャネルでの待受、シグナルを受け取るまで以降は処理されない
		<-sigCh

		log.Println("start graceful shutdown server.")
		// タイムアウトのコンテキストを設定
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
			// 接続されたままのコネクションも明示的に切る
			srv.Close()
		}
		log.Println("HTTPServer shutdown.")
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Print(err)
	}
}

// DrawFortune is ...
// おみくじを引く
func DrawFortune() string {
	fortune := []string{"大吉", "中吉", "吉", "凶"}
	rand.Seed(time.Now().UnixNano())
	return fortune[rand.Intn(len(fortune))]
}
