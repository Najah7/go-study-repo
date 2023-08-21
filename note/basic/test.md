# テスト

## 前提
- テストをソフトウェア開発で非常に大きな役割を果たす

## Goのテスティングフレームワーク
- 組み込み
    - testing
    - net/http/httptest
- サードパーティ
    - Ginkgo：BDDスタイルのテストフレームワーク
    - Gomega：Ginkgoのマッチャー

##　testing
- Goの組み込みのテストフレームワーク
- メインのGoの組み込みのメインのテストライブラリ
- プログラムファイルと同じパッケージにテストファイルを作成する
- テストの作成方法
    1. 対象の<関数名 or ファイル名>_test.goというファイルを作成（同じ階層に）
    2. テスト関数を作成
    3. `go test`コマンドでテストを実行(`go test ./...`)
- `testing.T`型：パッケージが提供する２つの主な構造体の1つであり、失敗やエラーを通知するのに使う主な構造体

## testing.T
- テストの失敗やエラーを通知するのに使う
- メソッド
    - `Name`：テスト関数内で実行中のテストの名前を取得することができます
    - `Parallel`：テスト関数を並列実行可能な形式にするために使用されます。このメソッドを呼び出すことで、テストランナーはテスト関数内で並列実行を許可し、複数のテスト関数を同時に実行することが可能になります
    - `Error`：テストが失敗したことを報告し、メッセージを出力しますが、テストの続行は行われます。
    - `Errorf`：テストが失敗したことを報告し、フォーマット付きメッセージを出力しますが、テストの続行は行われます。
    - `Fail`：テストを失敗させますが、メッセージは出力されません。テストの続行は行われます。
    - `FailNow`：テストを失敗させ、即座にテストの実行を中止します
    - `Fatal`：テストを失敗させ、メッセージを出力し、テストの実行を中止します。
    - `Fatalf`：テストを失敗させ、フォーマット付きメッセージを出力し、テストの実行を中止します。
    - `Log`：メッセージを出力しますが、テストが失敗しているかどうかに関わらず、テストの続行が行われます。
    - `Logf`：フォーマット付きメッセージを出力しますが、テストが失敗しているかどうかに関わらず、テストの続行が行われます。
    - `Skip`：テストをスキップし、メッセージを出力します。このメソッドを呼び出すと、テストはスキップされますが、同一スコープのテストは続行されません。
    - `SkipNow`：テストをスキップし、テストの実行を即座に中止しま
    - `Skipf`：フォーマット付きメッセージを出力し、テストをスキップします。このメソッドを呼び出すと、テストはスキップされますが、同一スコープのテストは続行されません。
    - `Skipped`：テストが実行されずにスキップされたことを示すテストの状態です。テストがスキップされると、そのテストの結果は skipped としてマークされ、テスト結果の集計に含まれます

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

## net/http/httptest
- Goの組み込みのテストフレームワーク
- HTTPサーバーのテストを行うためのライブラリ
- Webサーバーをシミュレーションする機能を提供する
- 仕組み
    1. マルチプレクサを作成
    1. テスト対象のハンドラをマルチプレクサに登録
    1. レコーダを作成
    1. リクエストを作成
    1. テスト対象のハンドラにリクエストを送信し、レコーダに記録
    1. レコーダにより、結果をチェック
- 代表的なメソッド
    - `NewRecorder`：レスポンスを記録するためのレコーダを作成する

## テストのサンプル
```go
package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    shutdown()
    os.Exit(code)
}

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func setup() *mux.Router {
    mux := http.NewServeMux()
    mux.HandleFunc("/post/", handleRequest(&FakePost{}))
    writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
    request, _ := http.NewRequest("GET", "/post/1", nil)
    mux.ServeHTTP(writer, request)
    if writer.Code != 200 {
        t.Errorf("Response code is %v", writer.Code)
    }
    var post Post
    json.Unmarshal(writer.Body.Bytes(), &post)
    if post.Id != 1 {
        t.Errorf("Cannot retrieve JSON post")
    }
}

func TestHandlePut(t *testing.T) {
    json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
    request, _ := http.NewRequest("PUT", "/post/1", json)
    mux.ServeHTTP(writer, request)
    if writer.Code != 200 {
        t.Errorf("Response code is %v", writer.Code)
    }

    request, _ = http.NewRequest("GET", "/post/1", nil)
    writer = httptest.NewRecorder()
    mux.ServeHTTP(writer, request)
    if writer.Code != 200 {
        t.Errorf("Response code is %v", writer.Code)
    }

    var post Post
    json.Unmarshal(writer.Body.Bytes(), &post)
    if post.Content != "Updated post" {
        t.Errorf("Post content is %v", post.Content)
    }
}

```

## テストタブルと依存性の注入
- テストを実行するために必要な環境を提供するためのコード(タブル：代役、影武者)
- 依存性の注入
    - デザインパターンの一種
    - ソフトウェアの複数のレイヤの依存性を分離するもの。
    - 方法：オブジェクト、構造体、関数などを呼び出す際に「依存性の源」を渡すことで、依存性を注入する
        - グローバル変数で呼び出すではなく、関数の引数として渡していく。（関数を入れ子にする場合のこと。深いレイヤーでグローバル変数を呼び出すのでなく、バケツリレー方式で渡していく。大体はオブジェクトや構造体で抽象化するので、必要ない関数では意識する必要はない。）
        - Goでは、インタフェース型がよく使われます
    - テストタブルを使って、テストの実行に必要な環境を提供すること
- 例：メールを送信する関数のテストで、メール送信をシミュレーションするための、テストダブルを作成する

## テストダブル v.s. モック
- テストダブル
    - テスト対象のコードを単独でテストするために、依存するコンポーネントや外部リソースの代わりに使用される置換物を指す
    - テストダブルの種類
        - モック
            - テスト中に特定の動作を呼び出しを予期した通りに動作されるために使われるテストダブル
            - テスト中に特定の振る舞いを制御するために使用されるダブルの一種
        - スタブ
            - 特定のメソッドや関数が呼び出されたときに返すべき返り値を事前に設定するテストダブル
            - 主にテスト中に外部サービスや依存関係を模倣し、テストの際に特定の結果を提供するために使用される
        - フェイク
            - 本物のコンポーネントや外部リソースの代わりに使われるテストダブルで、実際の実装を持っていることがある
            - 本物よりも簡易な実装で、テスト中にコントロール可能な振る舞いを持たせることを目的としいる
            - 本物のデータベースを使用する代わりに、メモリ内でデータを保持するフェイクデータベースを使用することが多い
        - スパイ
            - 特定のメソッドや関数が呼び出されたときに、その呼び出しに関する情報を収集するテストダブル
            - メソッドの呼び出し回数や引数、戻り値などを監視するために使用される
            - モックやスタブと組み合わせて使用されることが多く、メソッドの呼び出しを追跡しながらテスト中のコードの振る舞いを検証するのに役立つ

## サードパーティ製のテスト用Goライブラリ

### gocheck
- testingを利用して構築されている
- 主な機能
    - テストのグループ化によるスイートの作成
    - テストスイートあるいはテストケースごとのフィクスチャ(テストを実行するための環境)の設定
    - 拡張性のあるチェック用インタフェースのついたアサーション機能
    - より有用なエラー報告機能
- インストール：`go get gopkg.in/check.v1`
- 詳細を表示する：`go test -gocheck.vv`
- 代表的なメソッド
    - `Suite`：テストスイートを作成する
    - `Test`：テストケースを作成する
    - `SetUpSuite`：テストスイートのセットアップを行う
    - `TearDownSuite`：テストスイートの後処理を行う
    - `SetUpTest`：テストケースのセットアップを行う
    - `TearDownTest`：テストケースの後処理を行う
    - `Assert(result, <Operator>, expected)`：アサーションを行う。失敗した時点で呼び出しもとに戻る
        - Equal：`c.Assert(result, Equals, expected)`
        - NotEqual：`c.Assert(result, Not(Equals), expected)`
        - DeepEqual：`c.Assert(result, DeepEquals, expected)`
        - NotDeepEqual：`c.Assert(result, Not(DeepEquals), expected)`
        - IsNil：`c.Assert(result, IsNil)`
        - NotNil：`c.Assert(result, NotNil)`
    - `Check`：アサーションを行う。実行はテストケースの最後まで続く
    - `CheckEqual`：アサーションを行う
    - `CheckClose`：アサーションを行う
    - `CheckDeepEqual`：アサーションを行う


### サンプルコード
```go
package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    . "gopkg.in/check.v1"
)


type PostTestSuite struct{
    post *Post
}

func init() {
    Suite(&PostTestSuite{})
}

func Test(t *testing.T) {
    TestingT(t)
}

// フィクスチャの設定
func (t *testing.T) setUpTest(c *C) {
    s.post = &FakePost{}
    s.mux = http.NewServeMux()
    s.mux.HandleFunc("/post/", handleRequest(s.post))
    s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TestHandleGet(c *C) {
    request, _ := http.NewRequest("GET", "/post/1", nil)
    s.mux.ServeHTTP(s.writer, request)
    c.Check(s.writer.Code, Equals, 200)

    var post Post
    json.Unmarshal(s.writer.Body.Bytes(), &post)
    c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestHandlePut(c *C) {
    json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
    request, _ := http.NewRequest("PUT", "/post/1", json)
    s.mux.ServeHTTP(s.writer, request)
    c.Check(s.writer.Code, Equals, 200)

    request, _ = http.NewRequest("GET", "/post/1", nil)
    s.writer = httptest.NewRecorder()
    s.mux.ServeHTTP(s.writer, request)
    c.Check(s.writer.Code, Equals, 200)

    var post Post
    json.Unmarshal(s.writer.Body.Bytes(), &post)
    c.Check(post.Content, Equals, "Updated post")
}
```

### Ginkgo
- BDD(Behavior Driven Development)スタイルのテストフレームワーク
- 簡単に表すとBDDはテスト駆動開発の拡張版。テスト手法ではなく、開発技法の1つ。
- BDDのユーザストーリ（エンドユーザの言語と視点で書かれた要件定義）は振る舞いの観点から書かれている。
- TDDと同じくBDDでも先にテストを書くが、BDDではテストを書く前にユーザストーリを書く。
- 実際の現場では、高レベルのユーザストーリを書いた後に、そのユーザストーリを２、３階層、詳細化していく（高レベルのユーザストーリがテストスイートと対応する）
- ユーザーストーリーをテストケースに割り当てることを可能にする文法構造があり、testingと上手く統合されている。
- インストール：`go get github.com/onsi/ginkgo/ginkgo`　&& `go get github.com/onsi/gomega/...`
- テストの実行：`ginkgo`コマンドでテストを実行
- testingのテストをGinkgoのテストに変換：`ginkgo convert <テストファイル名>`
- 初期化(テストのセットアップ)：`ginkgo bootstrap`
- テストの作成：`ginkgo generate <テストファイル名>`
- 主なオプション
    - `-v`：詳細を表示する
    - `-r`：サブディレクトリを再帰的に検索する
    - `-skipPackage`：指定したパッケージをスキップする
    - `-focus`：指定したテストだけを実行する
    - `-nodes`：並列実行するノード数を指定する
    - `-randomizeAllSpecs`：テストの実行順序をランダムにする
    - `-randomizeSuites`：テストスイートの実行順序をランダムにする
    - `-seed`：ランダムなシード値を指定する
    - `-succinct`：テストの実行結果を簡潔に表示する
    - `-slowSpecThreshold`：テストの実行時間の閾値を指定する
    - `-cover`：カバレッジを表示する
    - `-coverpkg`：カバレッジを表示するパッケージを指定する
- 名前空間に`main`は使わない（mainから独立させるため）

### BDDのユーザストーリーの例
```
Story: "投稿を取得する"
In order to 投稿をユーザに提示する"
As a "呼び出し側のプログラム"
I want to "投稿を取得する"

Scenario 1: "idを使う"
Given "投稿のidが1"
when "そのidのGETリクエストを送信した"
Then "投稿が取得する"

Scenario 2: "非整数のidを使う"
Given "投稿の文字列がhello"
when "そのidのGETリクエストを送信した"
Then "HTTP 500のレスポンスを取得する"
```

### Ginkgoのテストのサンプル
```go
package main_test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    . "<path>/ginkgo_sample" // テスト対象のパッケージをインポート

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Get a post", func() {
    post = &FakePost{}
    var mux *http.ServeMux
    mux.HandleFunc("/post/", handleRequest(post))
    var writer *httptest.ResponseRecorder

    BeforeEach(func() {
        post = &FakePost{}
        mux = http.NewServeMux()
        mux.HandleFunc("/post/", handleRequest(post))
        writer = httptest.NewRecorder()
    })


    context("Get a post using on id", func() {
        It("should get a post", func() {
            request, _ := http.NewRequest("GET", "/post/1", nil)
            mux.ServeHTTP(writer, request)
            Expect(writer.Code).To(Equal(200))

            Expect(writer.Code).To(Equal(200))

            var post Post
            json.Unmarshal(writer.Body.Bytes(), &post)
            Expect(post.Id).To(Equal(1))
        })
    })

    context("Get an error if post id is not an integer", func() {
        It("should get a 500", func() {
            request, _ := http.NewRequest("GET", "/post/hello", nil)
            mux.ServeHTTP(writer, request)
            Expect(writer.Code).To(Equal(500))
        })
    })

})

```





## ベンチマークテスト
- 機能をテストするのではなく、作業単位のパフォーマンスを測定するためのテスト
- ベンチマークテストの作成方法
    1. 対象の<関数名 or ファイル名>_test.goというファイルを作成（同じ階層に）
    2. ベンチマーク関数を作成（`Benchmark_対象の関数名`）
    3. `go test -bench=.`コマンドでベンチマークテストを実行(`go test -bench=. ./...`)
- ベンチマークテストで使えるオプション
    - `-bench`：ベンチマークテストを実行する
    - `-cover`：カバレッジを表示する
    - `-run <正規表現> | <テスト名>`：指定の関数のテストのみ実行(存在しないテスト名 + -bench、でベンチマークテストだけを実行することができる)。複数指定する場合は、「|」で区切る
    - `-short`：テストを短縮する
    - `-benchmem`：ベンチマークテストの実行時にメモリの割り当て量を表示する
    - `-benchtime`：ベンチマークテストの実行時間を指定する
    - `-count`：ベンチマークテストの実行回数を指定する
    - `-cpu`：ベンチマークテストの実行時に使用するCPUの数を指定する
    - `-cpuprofile`：ベンチマークテストの実行時にCPUプロファイルを取得する
    - `-memprofile`：ベンチマークテストの実行時にメモリプロファイルを取得する
    
## 並行性（Concurrency）と並列性（Parallelism）

### 実行方法の整理
- 下記の2つの用語はどちらも複数のタスクが同時に行われることを指す
- 並行実行（プロセスが並んで（同じ時間に）重なり、実行する）
    - 複数のタスクが同時に進行しているように見える状態を指す（実行時間に重なりがあるだけ）
    - 多くのモノを同時に扱うこと
    - 主にタスクの進行と切り替えによって実現している。
    - プログラム内の異なる部分が同時に実行されているように見えるが、実際には一度に1つのタスクしか実行されていない場合も多い。
    - タスク間の相互作用がが可能（同じリソースが共有される）
    - タスク間の切り替えは、プログラムが実行中にタイムスライスを切り替えて行う
- 並列実行（プロセスが並列して重ならない）
    - 複数のタスクが同時に実際に同時に実行される状態
    - 多くのモノを同時に行うこと。
    - 複数のプロセッサやコアを使用して、異なるタスクが同時に処理されることが可能。
    - それぞれのタスクは完全に独立して実行され、タスクごとに別々の実行コンテキストが存在する。（独立したリソースが必要）
    - 大きな問題を細分化して、それぞれのタスクを並列に実行することで、処理速度を向上させる用途に使用する。
    - 複数のタスクを同時に処理して処理速度を向上させるために使用する
    
## GOMAXPROCS
- Goのランタイムが使用するスレッド数を制御するための環境変数
- デフォルトはCPUのコア数
- バージョン1.5未満は、デフォルトが１なので注意。

## Goと並列実行
- Go言語で、並列処理を作成することはできるが、実際に想定されいるのは、並列性でなく、並行性。
- 構成技術
    - goroutine
        - 複数が同時に実行される関数のこと
        - Go言語のランタイムによって管理される軽量なスレッドのようなもの（goroutine）を生成する(スレッドより軽量)
        - スタックサイズは2KBからスタートし、必要に応じて拡張される
        - gorutineは、軽量ではあるもの、起動するコストはそれなりにかかるので、たくさん作りすぎたら、逆にパフォーマンスが落ちる（順次実行のコストとgoroutineの起動コストのバランスを考える必要がある）<br>※👆CPUなどの独立したリソースの数にも依存する（しかし、複数のリソースのスケジューリングと実行にも別途コストがかかるので注意）
        - 並行処理するかの判断は、ベンチマークテストで決めるべき
    - waitgroup
        - goroutineの完了を待つための機能
        - syncパッケージに含まれる
        - sync.WaitGroup型のを使って作成
        - `Add`：WaitGroupに追加するgoroutineの数を指定する
        - `Done`：WaitGroupに追加されたgoroutineの数を減らす
    - channel
        - goroutine間でデータをやり取りするための機能
        - `ch := make(chan <型>)`でバッファなしのチャネルを作成
        - `ch := make(chan <型>, <バッファサイズ>)`でバッファありのチャネルを作成
            - バッファ：チャネルに格納できるデータの数
        - `ch <- <データ>`でチャネルにデータを送信
        - `<変数> := <- ch`でチャネルからデータを受信
        - `close(ch)`でチャネルをクローズ
            - チャネルのクローズは、チャネルの送信が終了したことを示す
        - `for <変数> := range ch`でチャネルからデータを受信する
            - チャネルの送受信は、データが送受信されるまでブロックされる
        - `<変数>, <OK> := <- ch`でチャネルからデータを受信し、チャネルがクローズされているかどうかを確認する
            - チャネルがクローズされている場合、OKはfalseになる
        - `select`：複数のチャンネルの中から受け渡しに使うチャンネルを選択できる
            ```go
            select {
                case <変数> := <- ch1:
                    // ch1からデータを受信
                case <変数> := <- ch2:
                    // ch2からデータを受信
                case ch3 <- <データ>:
                    // ch3にデータを送信
                default:
                    // どのチャネルも受信や送信ができない場合
            }
            ```
            - デッドロックには注意：goroutineがあるチャンネルから値を取り出すと、そのチャンネルから値を取り出す他のgoroutineがすべてブロックされスリーブしてしまう。
            - select文のdefault節を使って、デッドロックを回避することができる
        

    
