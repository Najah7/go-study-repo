# マルチプレクサ
- 多重化通信の入り口になるもの

# ServerMxu
- マルチプレクサのクラス
- 要件はServer同様、ServeHTTPメソッドを持つこと
- エントリを要素とするマップを持った構造体
- ServerHTTPに要求されたURLにもっとも近いエントリを探し、そのエントリに対応するハンドラを呼び出す
- 末尾が「/」でないURLの場合、完全一致を求める
- 末尾が「/」の場合、要求されたURLの先頭が登録URLと一致するかどうかを調べる

## エントリ
- URLをハンドラに対応づけたもの

## DefaultServeMux
- デフォルトのマルチプレクサ
- httpライブラリをimportしたアプリケーションが一般に利用できるSerevrMuxのインスタンス。
- 他にも、Server内にハンドラが指定されていない場合にも、DefaultServeMuxが利用される。

## http.NewServeMux()
- マルチプレクサを生成するクラス

## その他の代表的なマルチプレクサ
- Gorilla Toolkit
    - mux
    - pat
- HttpRouter
    - 軽量で、有力なサードパーティ製のマルチプレクサ
    - URLとのパターンマッチングで変数を使用できる（:idを使って、URLの一部を変数として扱うなど）
    - サンプル
    ```go

    mux.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        fmt.Fprintf(w, "Hello, %s\n", p.ByName("id"))
    })

    ```