# そもそもCookieの目的
- HTTPのステートレス性を補うために、クライアント側にデータを保存するための仕組み

# http.Cookie
- クッキーを表す構造体
- フィールド
    - Name
        - クッキーの名前
    - Value
        - クッキーの値
    - Path
        - クッキーが有効なパス
    - Domain
        - クッキーが有効なドメイン
    - Expires
        - クッキーの有効期限（time.Time型）
        - これが設定されていない場合、セッションCookieとなる(ブラウザが閉じられると削除される)
    - RawExpires
        - クッキーの有効期限（文字列）
    - MaxAge
        - クッキーの有効期限（秒）
    - Secure
        - HTTPSのみでクッキーを送信するかどうか
    - HttpOnly
        - HTTPのみでクッキーを送信するかどうか
    - Raw
        - クッキーの生の文字列
    - Unparsed
        - クッキーの生の文字列（パースされていない）

# http.setCookie()
- クッキーを設定する関数
- 引数
    - 第1引数：ResponseWriter
    - 第2引数：Cookien
- 例
```go
func session(w http.ResponseWriter, r *http.Request) {
    c := &http.Cookie{
        Name: "session",
        Value: "mycookie",
    }
    http.SetCookie(w, c)
}
```

## レスポンスヘッダーのSet-Cookie内にクッキーを設定する

```go
package main

import (
    "net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
    c := &http.Cookie{
        Name: "session",
        Value: "mycookie",
    }

    c2 := &http.Cookie{
        Name: "session2",
        Value: "mycookie2",
    }

    // set：ヘッダーにクッキーを設定する（上書き）
    // add：ヘッダーにクッキーを追加する（同じ名前のCookieがあってもOK
    w.Header().Set("Set-Cookie", c.String())
    w.Header().Add("Set-Cookie", c2.String())
}
```

## ブラウザからクッキーを取得する

```go
package main

import (
    "fmt"
    "net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
    c := &http.Cookie{
        Name: "session",
        Value: "mycookie",
        HttpOnly: true,
    }

    c2 := &http.Cookie{
        Name: "session2",
        Value: "mycookie2",
        HttpOnly: true,
    }

    w.Header().Set("Set-Cookie", c.String())
    w.Header().Add("Set-Cookie", c2.String())
}

func getCookie(w http.ResponseWriter, r *http.Request) {
    // Cookieを取得する
    h := r.Header["Cookie"]
    fmt.Fprintln(w, h)
}

func main() {
    server := http.Server{
        Addr: "127.0.0.0:8080",
    }

    http.HandleFunc("/set_cookie", setCookie)
    http.HandleFunc("/get_cookie", getCookie)
    server.ListenAndServe()
}
```

## r.Cookie() v.s. r.Cookies()
- r.Cookie()
    - 引数に指定した名前のクッキーを取得する
    - 引数
        - 第1引数：クッキーの名前
    - 戻り値
        - *http.Cookie
- r.Cookies()
    - クッキーを全て取得する
    - 戻り値
        - []*http.Cookie

```go
package main

import (
    "fmt"
    "net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
    c := &http.Cookie{
        Name: "session",
        Value: "mycookie",
        HttpOnly: true,
    }

    c2 := &http.Cookie{
        Name: "session2",
        Value: "mycookie2",
        HttpOnly: true,
    }

    w.Header().Set("Set-Cookie", c.String())
    w.Header().Add("Set-Cookie", c2.String())
}

func getCookie(w http.ResponseWriter, r *http.Request) {
    // Cookieを取得する
    h := r.Header["Cookie"]
    fmt.Fprintln(w, h)

    // Cookieを取得する
    c, err := r.Cookie("session")
    if err != nil {
        fmt.Fprintln(w, "Cookie session not found")
    }
    fmt.Fprintln(w, c)

    // Cookieを全て取得する
    cs := r.Cookies()
    fmt.Fprintln(w, cs)
}

func main() {
    server := http.Server{
        Addr: "127.0.0.0:8080",
    }

    http.HandleFunc("/set_cookie", setCookie)
    http.HandleFunc("/get_cookie", getCookie)
    server.ListenAndServe()
}
```

## Cookieを利用したフラッシュメッセージ

```go
package main

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
    msg := []byte("Hello World!")
    c := http.Cookie{
        Name: "flash",
        Value: base64.URLEncoding.EncodeToString(msg),
    }
    http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("flash")
    if err != nil {
        if err == http.ErrNoCookie {
            fmt.Fprintln(w, "メッセージがありません")
        }
    } else {
        rc := http.Cookie{
            Name: "flash",
            MaxAge: -1,
            Expires: time.Unix(1, 0),
        }
        http.SetCookie(w, &rc)
        val, _ := base64.URLEncoding.DecodeString(c.Value)
        fmt.Fprintln(w, string(val))
    }
}

func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    http.HandleFunc("/set_message", setMessage)
    http.HandleFunc("/show_message", showMessage)
    server.ListenAndServe()
}
```

※Encodingしている理由は、メッセージに空白や改行などの特殊文字が含まれている場合があり、Cookieに保存できないことがあるため。

※上記の処理のポイントは、Cookieはブラウザ側に保存されているということ、なので、同じ名前でCookieを上書きすることで、Cookieを削除している。