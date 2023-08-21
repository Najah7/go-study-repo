# データの記憶

## データの記憶方法
- メモリ内
- ファイルシステム上のファイル内
- サーバプログラムをフロントエンドとするデータベース内

## メモリ内での保持
- メモリ内での保持は、実行中のアプリケーション自体に保持すること。
- メモリ内のデータは通常データ構造内に格納される。
- メモリなので揮発性である。
- 主なデータ構造
    - 配列
    - スライス
    - マップ
    - 構造体：最もよく使われる。

## ファイルによる保持
- 一般的なファイルへの保存
    - csv
    - json ...etc
- gob
    - Goのバイナリ形式のデータ構造をファイルに保存するためのパッケージ
    - メモリ内のデータを高速かつ効果的にシリアル化して１つ以上のファイルに保存することができる。
    - 特徴
        - バイナリ形式なので、他の言語からは読み込めない。
        - 速度は速い。
        - ファイルサイズは小さい。
        - ファイルの中身を見ることができない。

## csvを扱うサンプルコード

```go

package main

import (
    "encoding/csv"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    data = []byte {
        "id,name,age",
        "1,山田太郎,20",
        "2,佐藤花子,30",
    }
    // ioutil.WriteFile()の第3引数は、ファイルのパーミッション
    err := ioutil.WriteFile("data.csv", data, 0666)
    if err != nil {
        panic(err)
    }

    read1, _ := ioutil.ReadFile("data.csv")
    fmt.Println(string(read1))

    // 戻り値は、*os.File型のポインタ
    file1, _ := osCreate("data2.csv")
    defer file1.Close()

    // 戻り値は、*csv.Writer型のポインタ
    write1 := csv.NewWriter(file1)
    write1.Write([]string{"1", "山田太郎", "20"})

    file2, _ := os.Open("data2.csv")
    defer file2.Close()

    read2 := csv.NewReader(file2)
    record, _ := read2.Read()
    fmt.Println(record)
}

```

### Writer.flush()
- バッファに溜まっているデータをファイルに書き込む

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, _ := os.Create("test.txt")
    defer file.Close()

    writer := bufio.NewWriter(file)
    writer.WriteString("bufio.Writer\n")
    writer.Flush()

    file2, _ := os.Create("test2.txt")
    defer file2.Close()

    file2.WriteString("os.File.WriteString\n")
}
```

### reader.FieldsPerRecord
- 1行目のフィールド数を基準に、2行目以降のフィールド数が異なる場合にエラーを返す
- -1の場合は、フィールド数のチェックを行わない
- 正の場合は、想定しているフィールド数を指定する

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func main() {
    file, _ := os.Open("data.csv")
    defer file.Close()

    reader := csv.NewReader(file)
    reader.FieldsPerRecord = 2

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        fmt.Println(record)
    }
}
```

### reader.ReadAll()
- ファイルの全ての行を読み込む

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func main() {
    file, _ := os.Open("data.csv")
    defer file.Close()

    reader := csv.NewReader(file)
    reader.FieldsPerRecord = 2

    records, _ := reader.ReadAll()
    for _, record := range records {
        fmt.Println(record)
    }
}
```

## jsonを扱うサンプルコード

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Person struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
}

func main() {
    person1 := Person{ID: 1, Name: "山田太郎", Age: 20}
    person2 := Person{ID: 2, Name: "佐藤花子", Age: 30}
    person3 := Person{ID: 3, Name: "鈴木次郎", Age: 40}
    people := []Person{person1, person2, person3}

    data1, _ := json.Marshal(people)
    fmt.Println(string(data1))

    ioutil.WriteFile("people.json", data1, 0666)

    data2, _ := ioutil.ReadFile("people.json")
    var people2 []Person
    json.Unmarshal(data2, &people2)
    fmt.Println(people2)

    file1, _ := os.Create("people2.json")
    defer file1.Close()

    encoder := json.NewEncoder(file1)
    encoder.SetIndent("", "    ")
    encoder.Encode(people)
}
```

## gobを扱うサンプルコード
- bytes.Buffer型は、バイト列を扱うための構造体。実質はバイトデータの可変バッファ。
- 流れ
    1. encoder.Encode()で、構造体をバイト列に変換する
    1. ioutil.WriteFile()で、バイト列をファイルに書き込む
    1. encoder.Decode()で、バイト列を構造体に変換する

```go
package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "io/ioutil"
    "os"
)

type Person struct {
    ID int
    Name string
    Age int
}

func main() {
    person1 := Person{ID: 1, Name: "山田太郎", Age: 20}
    person2 := Person{ID: 2, Name: "佐藤花子", Age: 30}
    person3 := Person{ID: 3, Name: "鈴木次郎", Age: 40}
    people := []Person{person1, person2, person3}

    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    encoder.Encode(people)

    ioutil.WriteFile("people.gob", buffer.Bytes(), 0666)

    data, _ := ioutil.ReadFile("people.gob")
    var people2 []Person
    decoder := gob.NewDecoder(bytes.NewReader(data))
    decoder.Decode(&people2)
    fmt.Println(people2)

    file1, _ := os.Create("people2.gob")
    defer file1.Close()

    encoder2 := gob.NewEncoder(file1)
    encoder2.Encode(people)
}
```

## GoでPostgreSQLのCRUDを行うサンプルコード
- GoでPostgreSQLを扱うためには、ドライバが必要
- ドライバは、database/sqlパッケージを使用して、データベースに接続する

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    // PostgreSQLのドライバ(勝手にinit()が実行される)
    // init()でsql.Register("postgres", &Driver{})が実行される
    _ "github.com/lib/pq"
)

type Person struct {
    ID int
    Name string
    Age int
}

var db *sql.DB

func init() {
    var err error
    // 第2引数は、データソース名（データベースドライバに固有の文字列、ドライバに接続方法を伝えるためにある）
    db, err = sql.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

fun People(limit int) ([]Person, error) {
    statement := `SELECT * FROM people`
    rows, err := db.Query(statement)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var people []Person
    for rows.Next() {
        var p Person
        err := rows.Scan(&p.ID, &p.Name, &p.Age)
        if err != nil {
            log.Fatal(err)
        }
        people = append(people, p)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(people)
}

func GetPerson(id int) (Person, error) {
    var person Person
    statement := `SELECT * FROM people WHERE id = $1`
    err := db.QueryRow(statement, id).Scan(&person.ID, &person.Name, &person.Age)
    if err != nil {
        return person, err
    }
    return person, nil
}

func Create(name string, age int) (int, error) {
    var id int
    statement := `INSERT INTO people (name, age) VALUES ($1, $2) RETURNING id`
    // Scan()：データベースから取得した値を引数に渡した変数に格納する
    // returning id：INSERT文の実行結果としてidを返す
    err := db.QueryRow(statement, name, age).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}

// 流れは、構造体を更新→DBに反映
// 下記のようなpersonを引数に渡す
// var person = GetPerson(p.ID)
//     person.Name = p.Name
//     person.Age = p.Age

func (p *Person) Update() error {
    
    statement := `UPDATE people SET name = $2, age = $3 WHERE id = $1`
    _, err := db.Exec(statement, p.ID, p.Name, p.Age)
    if err != nil {
        return err
    }
    return nil
}

func (p *Person) Delete() error {
    statement := `DELETE FROM people WHERE id = $1`
    _, err := db.Exec(statement, p.ID)
    if err != nil {
        return err
    }
    return nil
}



func main() {
    // Create()
    // People()
    // Person()
    // UpdatePerson()
    // DeletePerson()
}
```

### Scan()
- 引数に渡した変数に、データベースから取得した値を格納する
- 引数
    - 第1引数：データベースから取得した値
    - 第2引数：データベースから取得した値を格納する変数
- 戻り値
    - エラー

## rows.Next()
- 次の行があるかどうかを確認する＆次の行に移動する
- rowsは、データベースから取得した行の集合体（Query()やQueryRow()の戻り値）
- for文で回すことで、行を1行ずつ取得することができる


## GoとSQLのリレーション
- Goの構造体自体にリレーションを持たせることはできない。
- SQLでリレーションを定義するのみ。


## リレーションの種類
- 1対1(has one)
    - 1つのレコードに対して、1つのレコードが対応する関係
- 1対多
    - 1つのレコードに対して、複数のレコードが対応する関係
- 多対多
    - 複数のレコードに対して、複数のレコードが対応する関係
- 多対1(belongs to)
    - 複数のレコードに対して、1つのレコードが対応する関係

```go
if rows.Next() {
    var p Person
    err := rows.Scan(&p.ID, &p.Name, &p.Age)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(p)
}
```

## Goで使えるORM
- Sqlx
    - SQLのクエリを実行するためのパッケージ
    - database/sqlパッケージの拡張版
    - database/sqlパッケージの機能をそのまま使える
    - 付加機能
        - 構造体タグを利用することで、構造体、マップ、スライスのなどにフィールドごとにデータベースのレコードを設定する
        - プリペアドステートメントに名前付き引数を使用することができる
    - インストール：`go get github.com/jmoiron/sqlx`
- Gorm
    - Goで最も人気のあるORM
    - DataMapperパターンを採用している
    - DBの中のデータを構造体に対応づけるマッパーを提供している
    - 主な機能
        - 関係性の定義
        - マイグレーションの実行
        - チェインクエリ
        - コールバックの使用 ...etc
    - インストール：`go get github.com/jinzhu/gorm`


## Sqlxのサンプルコード

```go
package main

import (
    "fmt"
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)


var Db *sqlx.DB

func init() {
    var err error
    Db, err = sqlx.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

type Person struct {
    ID int `db:"id"`
    Name string `db:"name"`
    Age int `db:"age"`
}

func GetPerson(id int) (Person, error) {
    var person Person
    err := Db.QueryRowx("SELECT * FROM people WHERE id = $1", id).StructScan(&person)
    if err != nil {
        return person, err
    }
    return person, nil
}

func Create(name string, age int) (int, error) {
    var id int
    stmt, err := Db.Preparex("INSERT INTO people (name, age) VALUES ($1, $2) RETURNING id")
    if err != nil {
        return 0, err
    }
    defer stmt.Close()

    err = stmt.QueryRowx(name, age).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}

func (p *Person) Update() error {
    stmt, err := Db.Preparex("UPDATE people SET name = $2, age = $3 WHERE id = $1")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(p.ID, p.Name, p.Age)
    if err != nil {
        return err
    }
    return nil
}

func (p *Person) Delete() error {
    stmt, err := Db.Preparex("DELETE FROM people WHERE id = $1")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(p.ID)
    if err != nil {
        return err
    }
    return nil
}

func main() {
    // Create()
    // GetPerson()
    // UpdatePerson()
    // DeletePerson()
}
```


## Gormのサンプルコード

```go
package main

import (
    "fmt"
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
)

var Db *gorm.DB

func init() {
    var err error
    Db, err = gorm.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

type Person struct {
    gorm.Model
    Name string
    Age int
}

func GetPerson(id int) (Person, error) {
    var person Person
    err := Db.First(&person, id).Error
    if err != nil {
        return person, err
    }
    return person, nil
}

func Create(name string, age int) (int, error) {
    person := Person{Name: name, Age: age}
    Db.Create(&person)
    return person.ID, nil
}

func (p *Person) Update() error {
    Db.Save(&p)
    return nil
}

func (p *Person) Delete() error {
    Db.Delete(&p)
    return nil
}

func main() {
    // Create()
    // GetPerson()
    // UpdatePerson()
    // DeletePerson()
}
```
