# エラー

## Goのエラーの型
- Goのエラーの型は`error`型
- `error`型は組み込みインターフェース
- `error`型は`Error() string`メソッドを持つ

## エラーの定義
- エラーを定義するには、`errors`パッケージの`New()`関数を使う
- `New()`関数は、`error`型を返す
- `New()`関数の引数には、エラーの内容を文字列で渡す
- `New()`関数の戻り値は、`error`型の値
- `error`型の値は、`Error() string`メソッドを持つ
- `Error() string`メソッドは、エラーの内容を文字列で返す

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("エラーです")
    fmt.Println(err.Error())
}
```
