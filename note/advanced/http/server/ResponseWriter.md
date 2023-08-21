## responseWriter
- レスポンスへのインタフェース
- レスポンスへの書き込みを行うためのもの
- メソッド
    - Write
        - バイト列を書き込む
        - レスポンスボディに書き込む
    - WriteHeader
        - ステータスコードを書き込む
        - 引数の型はint
        - エラーメッセージを書く場合は`fmt.Fprint(w, "エラーメッセージ")`とする
    - Header
        - ヘッダーを書き込む
        - `w.Header().Set("Content-Type", "application/json")`

## HeaderとWriteHeaderの注意点
- WriteHeaderを呼びだした後のヘッダーの変更を防ぐので、WriteHeaderを呼び出した後にHeaderを変更することはできない。

## Jsonのレスポンスを返す
- `w.Header().Set("Content-Type", "application/json")`
- `json.Marshal()`：Goのデータ構造をJSON形式の文字列に変換する。戻り値は[]byte型