# 典型的なGo Web アプリのデザイン

## 典型的なGo Web アプリの構成要素
- マルチプレクサ
    - リクエストを受け取り、それを処理するハンドラーを呼び出す
- ハンドラー
    - リクエストを処理する
    - MVCのコントローラーに相当
- データモデル
    - データを表現する
    - MVCのモデルに相当
    - これを基にDBに保存したり、JSONに変換したりする
- テンプレートエンジン

## マルチプレクサのサンプルコード
```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/about", aboutHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About Page")
}
```

※Go言語ではコンパイル時にファイル内の関数を識別し、パッケージ外からアクセスすることができるので、関数の定義順序は重要でない

## ハンドラ
- ハンドラは`http.Handler`インターフェースを実装する
- `ServeHTTP`メソッドを実装する
- `ServeHTTP`メソッドは`http.ResponseWriter`と`*http.Request`を引数に取る
- `http.ResponseWriter`はレスポンスを書き込むためのインターフェース
- `*http.Request`はリクエストを表す構造体
- `http.ResponseWriter`の`Write`メソッドを使ってレスポンスを書き込む
- `http.Request`の`URL`フィールドにリクエストURLが格納されている

## ハンドラのサンプルコード
```go
package main

import (
    "fmt"
    "net/http"
)

type indexHandler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.Handle("/", &indexHandler{})
    http.ListenAndServe(":8080", nil)
}
```
