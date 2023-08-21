# ハンドラ
- ServerHTTPメソッドを持つインターフェース

# ServerHTTPメソッド
- ハンドラの実装
- 下記の引数を持つ
    - 第1引数：レスポンスを書き込むための構造体
    - 第2引数：リクエストを表す構造体

## ハンドラ関数
- ハンドラのように振る舞う関数
- ServerHTTPと同じシグネチャを持つ。（引数が同じ）

## ハンドラの登録のサンプルコード
```go

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "World")
}

func main() {
    hello := HelloHandler{}
    world := WorldHandler{}

    server := http.Server{
        Addr: "127.0.0.0:8080",
    }

    http.Handle("/hello", &hello)
    http.Handle("/world", &world)

    server.ListenAndServe()
}
```

## ハンドラ関数をハンドラに変換
- HandlerFunc型を使う
- 下記のコードはキャストしている
```go
helloHandler := HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
})
```

## `http.HandleFunc()`
- デフォルトのServeMuxにハンドラを登録するメソッド
- 引数
    - 第1引数：パス
    - 第2引数：ハンドラ関数

## `http.Handle()`
- デフォルトのServeMuxにハンドラを登録するメソッド
- 引数
    - 第1引数：パス
    - 第2引数：ハンドラ

## `ServerMux.HandleFunc()`
- ハンドラ関数を指定のマルチプレクサに登録するメソッド
- 引数
    - 第1引数：パス
    - 第2引数：ハンドラ関数


## `ServerMux.Handle()`
- ハンドラを指定のマルチプレクサに登録するメソッド
- ハンドラの引数にはResponseWriterとRequestを渡す
    - ResponseWriter
        - レスポンスを書き込むための構造体
    - Request
        - リクエストを表す構造体
- 引数
    - 第1引数：パス
    - 第2引数：ハンドラ
- 一般的には下記のように使う

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World")
})
```

## Handle v.s. HandleFunc
- Handle
    - 第2引数にはHandlerインターフェースを実装した構造体を渡す
- HandleFunc
    - 第2引数には関数を渡す
    - 関数の引数はHandlerFunc型の引数と同じ

## ハンドラとハンドラ関数のチェイン
- ハンドラとハンドラ関数をチェインすることができる

## チェイン（チェイニング）
- 関数を連鎖的に呼び出す手法やパターン
- 関数型プログラミングの原則に基づいており、コードを簡潔で読みやすく、複雑さを減少させるのに役立つ
- これにより、共通の処理などを分離させ、再利用性を高めることができる

## チェインのサンプルコード(ハンドラ関数版)
```go
func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("ハンドラ関数が呼び出されました")
        h(w, r)
    }
}

func main() {
    http.HandleFunc("/hello", log(hello))
    http.ListenAndServe(":8080", nil)
}
```

## チェインのサンプルコード(ハンドラ版)
```go
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello")
}

func log(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("ハンドラが呼び出されました")
        h.ServeHTTP(w, r)
    })
}

func main() {
    hello := HelloHandler{}
    http.Handle("/hello", log(&hello))
    http.ListenAndServe(":8080", nil)
}
```
