# Request

## ParseForm
- フォームをパースするための関数
- フォームのデータを解析してr.Formにデータを格納する
- URLの値もフォームの値もr.Formに格納される

### Form
- フォームのデータを格納するためのmap
- mapのキーはフォームのname属性の値
- mapの値はフォームのvalue属性の値
- フォームとURLに同値のキーがある場合両方の値をスライスに入れてくれるが、mapなので並び順は保証されない

## PostForm
- POSTメソッドで送信されたフォームの値を取得するための関数
- 常にフォームの値を優先してURLよりも前に配置してくれる
- 使い方はParseFormと同じ

### サンプル

#### フォーム全体
```go

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        fmt.Fprintf(w, "Hello, %s!", r.Form)
    })
    http.ListenAndServe(":8080", nil)
}

```

#### 特定のフォームの値
```go

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        fmt.Fprintf(w, "Hello, %s!", r.Form["name"])
    })
    http.ListenAndServe(":8080", nil)
}
```

## ParseMultiPartForm
- マルチパートフォームをパースするための関数
- URLの値は入っていない。URLの値を取得する場合はParseFormを使う
- マルチパートフォームのデータを解析してr.Formにデータを格納する
- 引数
    - maxMemory int64
        - フォームの最大メモリ容量(バイト)
        - フォームのデータがこの値を超える場合はディスクに書き込まれる
        - デフォルトは32MB

## MultiPartForm
- マルチパートフォームのデータを格納するためのmap
- URLの値は入っていない。URLの値を取得する場合はParseFormを使う

### サンプル

#### フォーム全体
```go

func main () {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultiPartForm(32 << 20)
        fmt.Fprintf(w, "Hello, %s!", r.Form)
    })
    http.ListenAndServe(":8080", nil)
}

```

#### 特定のフォームの値
```go

func main () {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultiPartForm(32 << 20)
        fmt.Fprintf(w, "Hello, %s!", r.Form["name"])
    })
    http.ListenAndServe(":8080", nil)
}

```

## FormValue
- 特定のフォームの値を取得するための関数
- FormParseやParseMultiPartFormを呼び出さなくてもURLの値やフォームの値を取得できる(裏で呼び出してくれている）

### サンプル（FormValue）
```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
    })
    http.ListenAndServe(":8080", nil)
}
```

## PostFormValue
- 特定のフォームの値を取得するための関数
- FormParseやParseMultiPartFormを呼び出さなくてもURLの値やフォームの値を取得できる(裏で呼び出してくれている）

### サンプル（PostFormValue）
```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", r.PostFormValue("name"))
    })
    http.ListenAndServe(":8080", nil)
}
```

## FormValueとPostFormValueの違い
- FormValueはURLの値もフォームの値も取得できる
- PostFormValueはフォームの値のみ取得できる

## FormValueとPostFormValueの注意点
- MultiPartFormの値は取得できない

## フォーム系の関数のまとめ
| フィールド | 呼び出すべきメソッド | URL | フォーム | URLエンコード | マルチパートフォーム |
| --- | --- | --- | --- | --- | --- |
| Form | ParseForm | ○ | ○ | ○ | × |
| PostForm | ParseForm | × | ○ | ○ | × |
| MultiPartForm | ParseMultiPartForm | × | o | × | ○ |
| FormValue | ParseForm | ○ | ○ | ○ | × |
| PostFormValue | ParseForm | × | ○ | ○ | × |

## フォームでのファイルのアップロード

```html
<form action="/process" method="post" enctype="multipart/form-data">
    <input type="file" name="uploaded">
    <input type="submit" value="Upload">
</form>
```

```go

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
    r.ParseMultiPartForm(32 << 20)
    fileHeader := r.MultiPartForm.File["uploaded"][0]
    file, err := fileHeader.Open()
    if err == nil {
        data, err := ioutil.ReadAll(file)
        if err == nil {
            fmt.Fprintln(w, string(data))
        }
    }
}


func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/process", process)
    server.ListenAndServe()
}
```

## FormFile
- フォームでアップロードされたファイルを取得するための関数
- 引数
    - key string
        - フォームのname属性の値
- 戻り値
    - multipart.File
        - ファイルのデータ
    - multipart.FileHeader
        - ファイルのヘッダー情報

### サンプル
```go
func process(w http.ResponseWriter, r * http.Request) {
    file, fileHeader, err := r.FormFile("uploaded")
    if err == nil {
        data, err := ioutil.ReadAll(file)
        if err == nil {
            fmt.Fprintln(w, string(data))
        }
    }
}
```