# Server
- Serverの構造体
- この構造体を使ってサーバーを設定する
```go
type Server struct {
    Addr string
    Handler Handler
    ReadTimeout time.Duration
    WriteTimeout time.Duration
    MaxHeaderBytes int
    TLSConfig *tls.Config
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```

## server.ListenAndServe()
- サーバーを起動する関数
- 引数（Optional）
    - 第1引数：ポート番号。デフォルトは、80。
    - 第2引数：ハンドラ。デフォルトは、DefaultServeMux。
- 例
    - `server.ListenAndServe(":3000", nil)`

## server.ListenAndServeTLS()
- サーバーを起動する関数
- httpsを使う場合に使用する
- 引数
    - 第1引数：ポート番号。デフォルトは、443。
    - 第2引数：証明書のパス
    - 第3引数：秘密鍵のパス
    - 第4引数：ハンドラ。デフォルトは、DefaultServeMux。
- 例
    - `server.ListenAndServeTLS(":3000", "server.crt", "server.key", nil)`

## http.FileServer()
- 静的ファイルを配信するハンドラを生成する関数
- 引数
    - 第1引数：ファイルのルートディレクトリ
    - 第2引数：http.Dir型の変数
- 例
    - http.FileServer(http.Dir("public"))