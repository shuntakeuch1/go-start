package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}
// ServeHttpはHTTPリクエストを処理しまままままます
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func () {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil) // 本当は戻り値をチェックすべき
}

func main() {
	http.Handle("/", &templateHandler{filename:"chat.html"})
	// Webサーバーを開始します
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("LitenAdServe:", err)
	}
}
