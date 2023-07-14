# Goの基本

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


