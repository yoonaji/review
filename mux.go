package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()
	//health로 요청을 주면 func를 실행한다는 뜻인듯?
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	return mux
}
