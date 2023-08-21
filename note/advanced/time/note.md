## timeパッケージ
- Goの標準ライブラリに含まれるパッケージ
- 時間を扱うためのパッケージ
- timeパッケージでは、日付フォーマットを指定する際に特定の数値を使用する
- フォーマットで使う値
    - 2006：年（西暦）
    - 01：月
    - 02：日
    - 03：時（12時間表記）
    - 04：分
    - 05：秒
    - 06：10進数のミリ秒
    - -07：タイムゾーン
    - Mon: 短縮形の曜日
    - Monday：曜日（完全）
    - .000：ミリ秒
    - AM：午前午後
    - PM：午前午後

## timeモジュールを使ったサンプルコード

```go
package main

import (
    "html/template"
    "net/http"
    "time"
)

func process(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl.html")
    t.Execute(w, time.Now())
}

func main() {
    server := http.Server{
        Addr: "127.0.0.0:8080",
    }
    http.HandleFunc("/process", process)
    server.ListenAndServe()
}
```