# Goの基本

## Goの特徴
- コンパイル言語（環境に依存しないバイナリを生成できる）
- 速度は、インタプリタや、JVM言語よりも速い(しかし、CやRustなどよりは少し遅い)
- コンパイルはかなり高速
- Go Runtimeにより、かなりの部分が自動的に最適化される
- メモリの使用量もインタプリタや、JVM言語よりも少ない（しかし、CやRustなどよりは多い。なぜなら、バイナリ＋Go Runtimeを含めるため）

## Go Runtimeの主な機能
- メモリ管理
- ガベージコレクション
- スケジューリング
- チャネル
- マルチスレッド


## main関数
- プログラムのエントリーポイントとして機能する関数

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## init関数
- main関数よりも先に実行される関数。
- 初期化処理などを行うのに使われる。

```go
package main

import "fmt"

func init() {
    fmt.Println("init")
}
```

## コメントアウト
- `//` でコメントアウト
- `/* */` でコメントアウト

## import
- パッケージをインポートする
- [標準パッケージサイト](https://pkg.go.dev/std)
- 基本的なimport
    - `import "fmt"` でfmtパッケージをインポートする
- 複数のパッケージをインポートする
    - `import ( "fmt" "math" )` でfmtパッケージとmathパッケージをインポートする
- 別名でインポートする
    - `import f "fmt"` でfmtパッケージをfという名前でインポートする
    - `import . "fmt"` でfmtパッケージをインポートする。この場合、fmt.Println()の代わりにPrintln()と書ける。
    - 注意👇
        - `import _ "fmt"` でfmtパッケージをインポートする。この場合、fmtパッケージの関数を使うことはできない。


