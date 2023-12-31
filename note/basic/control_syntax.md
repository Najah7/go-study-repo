# 制御構文

## if文
- 条件式は()で囲まない

### 基本
- and: `&&`
- or: `||`
- not: `!`
```go
if 条件式 {
    処理
} else if 条件式 {
    処理
} else {
    処理
}
```

### 値の代入&条件式（ローカル変数の定義）
- if文の条件式内で変数を宣言することができる
- if文の条件式内で宣言した変数は、if文内でのみ使用可能
```go
if 変数名 := 値; 条件式 {
    処理
}
```

## for文
- ()で囲まない
- while文はない

### 基本
```go
for 初期化; 条件式; 更新処理 {
    処理
}
```

### while代用としての使い方
- 初期化と更新処理を省略することで、while文のように使うことができる
```go
for ; 条件式; {
    処理
}

sum := 1
for sum < 1000 {
    sum += sum
}
```

### range
- スライスやマップの要素を1つずつ取り出す
```go
for インデックス, 要素 := range スライス {
    処理
}

// valueのみ取り出す場合
for _, value := range スライス {
    処理
}

// mapの場合
for キー, 値 := range マップ {
    処理
}

// valueのみ取り出す場合(mapの場合は省略可能)
for value := range マップ {
    処理
}
```
※keyはmapの場合は省略可能だが、スライスの場合はエラーになるので注意。

## switch文
- 条件式は()で囲まない
- defaultは省略可能

### 基本
```go
switch 条件式 {
case 値:
    処理
case 値:
    処理
default:
    処理
}
```

### 代入&条件式（ローカル変数の定義）
- switch文の条件式内で変数を宣言することができる
- switch文の条件式内で宣言した変数は、switch文内でのみ使用可能
```go
switch 変数名 := 値; 条件式 {
case 値:
    処理
case 値:
    処理
default:
    処理
}
```

### 条件式の省略
- switch文の条件式を省略することができる
- 条件式を省略した場合は、caseの条件式がtrueの場合に処理が実行される
```go
switch {
case 条件式:
    処理
case 条件式:
    処理
default:
    処理
}
```

## defer文
- 関数の終了時に実行される処理を登録する
- defer文で登録した処理は、登録した順番の逆順で実行される(スタックに積まれるから)
- 関数の終了時に実行される処理は、return文の後に実行される
- defer文で登録した処理は、関数の終了時に実行されるため、関数内でエラーが発生した場合でも実行される
- ファイルのクローズ処理などに使用する
```go
defer 処理
```

## log
- ログを出力する
- ログの出力先は標準エラー出力
- ログのフォーマットはデフォルトで`2009/01/23 01:23:23 message`の形式
- 代表的な関数
    - Print：ログを出力する
    - Printf：フォーマットを指定してログを出力する
    - Println：ログを出力し、最後に改行を追加する
    - Fatal：ログを出力し、os.Exit(1)を実行する
    - Fatalf：フォーマットを指定してログを出力し、os.Exit(1)を実行する
    - Fatalln：ログを出力し、最後に改行を追加し、os.Exit(1)を実行する
    - Panic：ログを出力し、panicを実行する
    - Panicf：フォーマットを指定してログを出力し、panicを実行する
    - Panicln：ログを出力し、最後に改行を追加し、panicを実行する

### ログの出力
```go
log.Print("message")
```

## エラーハンドル
- Goでは、関数の戻り値としてエラーを返すことが多い
- エラーを返す関数を呼び出す際は、戻り値のエラーをチェックする必要がある
- エラーをチェックする際は、if文を使用する
- エラーをチェックする際は、変数の再定義とエラーのチェックを同時に行うことができる

### 基本
```go
tmp, err := 関数()
if err != nil {
    log.Fatal(err)
}
```

### １行での書き方
```go
if tmp, err := 関数(); err != nil {
    log.Fatal(err)
}
```

## panic & recover
- panic：エラーを発生させる
- recover：panicで発生したエラーを取得する
- panicで発生したエラーをrecoverで取得することで、エラーをハンドルすることができる
- panicは基本的には使用しない。しっかりerroをハンドルすることが重要。
```go
func main() {
    defer func() {
        if x := recover(); x != nil {
            log.Fatal(x)
        }
    }()

    panic("runtime error")
}
```
