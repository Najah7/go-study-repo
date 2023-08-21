# テスト

## テストの作成方法
1. 対象の関数名_test.goというファイルを作成（同じ階層に）
2. テスト関数を作成
3. `go test`コマンドでテストを実行(`go test ./...`)

## 詳細を表示する
- `go test -v`コマンドで詳細を表示する

## テスト関数の作成
- テスト関数は、`Test対象の関数名`で始める
- テスト関数は、引数にテストパッケージの`*testing.T`を取る

```go
func Test対象の関数名(t *testing.T) {
    // テストの内容
}
```

## テストのサンプル

```go
// add.go

package main

func Add(x, y int) int {
    return x + y
}
```

```go
// add_test.go

package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    if Add(1, 2) != 3 {
        t.Fatal("Add(1, 2) should be 3, but doesn't match")
    }
}
```

## Goのテスティングフレームワーク
- Ginkgo：BDDスタイルのテストフレームワーク
- Gomega：Ginkgoのマッチャー