# Header 
- HTTP Headerの構造体
- HTTP Headerはkey-valueの形式のmap
- key：文字列型。string
- value：文字列型のスライス。[]string。スライスの最初の要素が新しいヘッダの値になる。

# Headerの操作
- 値の追加：既存の文字列のスライスに追加
- 値の取得
    - すべてを取得：`Header()`メソッド
    - 特定のキーの値を取得：`Get("<key>")`メソッド
- 値の削除
- 値の設定

# Body
- HTTP Bodyフィールドの構造体
- Bodyはio.ReadCloserインターフェースを満たす
- io.ReadCloserインターフェースはio.Readerとio.Closerインターフェースを満たす
- io.ReaderインターフェースはRead()メソッドを持つ
- io.CloserインターフェースはClose()メソッドを持つ