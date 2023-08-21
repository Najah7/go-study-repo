# インターフェース

## インターフェースとは
- クラスの機能を定義したものです。
- 使い方は型と同じイメージ。ただ、型と違い、振る舞いを定義してる

### インターフェースの定義
```go
type インターフェース名 interface {
    メソッド名(引数の型, 引数の型) 戻り値の型
    メソッド名(引数の型, 引数の型) 戻り値の型
}
```

### インターフェースの実装
- インターフェースを実装するには、インターフェースで定義されているメソッドを全て実装する必要がある
- インターフェースを実装するには、構造体を定義する
- 構造体にインターフェースで定義されているメソッドを実装する
- 構造体にインターフェースを実装するには、構造体のポインタ型を指定する
```go
func (レシーバー *構造体名) メソッド名(引数 引数の型, 引数 引数の型) 戻り値の型 {
    処理
    return 戻り値
}
```

### インターフェースの呼び出し
- インターフェースを呼び出すには、インターフェースの変数を定義する
- インターフェースの変数に構造体のポインタ型を代入する

```go
var 変数名 インターフェース名 = 構造体のポインタ型
```

## 空のインターフェース
- 空のインターフェースは、全ての型を代入できる
- 空のインターフェースは、型アサーションを使って、型を取得できる
- インターフェースをどんな型でも受け取ることができる変数として扱える

### 定義
```go
var 変数名 interface{} = 値
```

## タイプアサーション
- インターフェース変数内の具体的な値を明らかにする
- 上記の空のインターフェースの型を指定するためのモノ

### 定義
```go
変数名 := インターフェース名.(型)
```

### switch文でのタイプアサーション
- switch文でのタイプアサーションは、インターフェースの変数に構造体のポインタ型を代入するときに使う

```go
switch 変数名 := インターフェース名.(type) {
case 構造体のポインタ型:
    処理
case 構造体のポインタ型:
    処理
default:
    処理
}
```

## Stringerインターフェース
- Stringerインターフェースは、fmtパッケージで定義されているインターフェース
- Stringerインターフェースは、String() stringメソッドを定義している
- Stringerインターフェースは、fmtパッケージのPrintln()関数で使われている
- pythonの__str__メソッドのようなもの

### 定義
```go
type Stringer interface {
    String() string
}
```

## インターフェースの埋め込み
- インターフェースを埋め込むことで、インターフェースのメソッドを使うことができる
- インターフェースを埋め込むことで、インターフェースのメソッドを実装する必要がなくなる
```go
func (レシーバー *構造体名) String() string {
    処理
    return fmt.Sprintf("構造体の値")
}
```