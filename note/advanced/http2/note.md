# HTTP/2
- http2.ConfigureServerを設定するのみ
    - HTTP/2のサーバーを設定する
    - この関数を呼び出すと、サーバーはHTTP/2のクライアントを受け入れるようになる
    - この関数を呼び出す前に、TLSの設定を行う必要がある
    - この関数を呼び出すと、サーバーはHTTP/2のクライアントを受け入れるようになる
    - この関数を呼び出す前に、TLSの設定を行う必要がある

## インストール

```bash
go get -u golang.org/x/net/http2
```

## サンプルコード

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func main() {
    handler := &MyHandler{}

    server := &http.Server{
        Addr:    ":8080",
        Handler: handler,
    }

    http2.ConfigureServer(server, &http2.Server{})

    log.Println("Listening on http://example.com:8080")
}
```

## cURL
- ユーザーがURLを指定して、ファイルを送受信できるコマンドラインツール
- HTTPやHTTPSなど、多数の一般的なインターネットプロトコルに対応している。
- --http2を指定すれば、HTTP/2で通信できる

```bash
curl --http2 -v https://localhost:8080
```