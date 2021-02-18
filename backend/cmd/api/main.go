package main

import (
	"net/http"

	"github.com/tetsuzawa/web-spat"
)

func main() {
	http.HandleFunc("/", backend.Handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":1991", nil)
}
