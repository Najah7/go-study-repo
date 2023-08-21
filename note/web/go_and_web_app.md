# Go言語とWebアプリ

## Go言語について
- バックエンドスステムのプログラムを単純かつ効率的に書くために開発された言語
- Googleが開発した言語
- 先進的な技術が詰まっている

## Goの特徴
- 手続き型に慣れているプログラマにとって馴染みやすい
- 関数型言語の特徴も兼ね備えている
- 並行実行を標準でサポート
- モダンなパッケージ作成システム
- ガベージコレクション
- 広範囲で強力な標準ライブラリが組み込まれている

## Webアプリの開発言語が必要とする性質
- スケーラビリティ
    - 垂直スケーラビリティ
        - Goは並行処理（Goroutine） により、こっちに優れている
    - 水平スケーラビリティ
        - 1つのGoアプリを多数のインスタンスの上にプロキシを置いて分散させることができる
        - 動的な依存関係のない静的なバイナリとしてコンパイルされるので、Go言語が組み込まれていないシステムにも分散できる。
- モジュール性
    - 変更に柔軟なアプリを作るうえで必須
    - インタフェースにより、動的な型付けを行うことができる
    - 関数はインタフェースを引数としてとれるので、既存の関数に対して、そのインタフェースが要求するメソッドを実装することで、その関数をそのまま使い続けることができる
    - 引数に空のインタフェースを渡すことで、引数にどんな型でも渡すことができる
- 保守性
    - ソフトウェア工学のグットプラクティスを推進するように、デザインされている。
    - go doc
    - go fmt
- 高い実行効率
    - GoはC言語の実行効率を目指している

## Webアプリ v.s. Webサーバー v.s. Webサービス
- Webアプリケーション
    - アプリケーション：ユーザとのイントラクション（やり取り）によるユーザーが何かの活動をするのを援助するソフトウェア
    - Webアプリ：Web経由で配信され使用されるアプリケーション
    - Webアプリケーションの2つの条件
        - 呼び出したクライアントにHTMLを返す。クライアントはHTMLをレンダリングし、ユーザーに提示する
        - データはHTTPを使ってクライアントに送信される
- Webサーバー
    - 要求されたファイルのみを返す特殊なタイプのWebアプリケーション
- Webサービス
    - プルグラムがユーザーに対してHTMLを返さず、別のプログラムに他のフォーマットでデータを返すWebアプリケーション

## Webアプリの3本の柱
- データの保持
- テンプレート
- ハンドラ


## HTTPの単純な特徴
- テキストベースのリクエスト / レスポンス型のプロトコル
- クライアントサーバー型のコンピューティングモデルで使われる
- ステートレスなプロトコル
    - リクエストとレスポンスの間に状態を保持しない
    - リクエストとレスポンスの間に状態を保持するために、Cookieを使う
    - 利点
        - 目的は、プロトコルアナライザなしで確認できる
        - トラブルシューティングが容易
- リクエストとレスポンスの構造
    - リクエスト(リクエストヘッダーとリクエストボディの間に空行がある)
        - リクエストライン
            - リクエストメソッド
            - リクエストURI
            - HTTPバージョン
        - リクエストヘッダー
            - リクエストヘッダー名
            - リクエストヘッダー値
        - リクエストボディ
    - レスポンス
        - ステータスライン
            - ステータスコード
            - ステータスメッセージ
        - レスポンスヘッダー
            - レスポンスヘッダー名
            - レスポンスヘッダー値
        - レスポンスボディ

## Webアプリケーションの誕生
- 動的なコンテンツにしたいという要求から
- カスタマイズされた動的コンテンツをHTTP経由でユーザーに配信することから始まりました。
- そこで出てきたのが、CGI(Common Gateway Interface)
- CGI
    - 単純なインタフェース
    - Webサーバーが外部プロセスで実行されているプログラムと通信することを可能に。<br>上記を満たすものは、すべてCGI。どんな言語でもかける。
    - プログラムの入力は、環境変数と標準入力
    - プログラムの出力は、サーバ経由でクライアントに送信される
- SSI(Server-side includes)
    - HTMLファイルに含めることが可能な「ディレクティブ（命令）」（タグ）

## HTTP1.0

## HTTPメソッド
|メソッド|説明|
|:--|:--|
|GET|リソースの取得|
|HEAD|リソースのヘッダーの取得|
|POST|リソースの作成|
|PUT|リソースの更新|
|DELETE|リソースの削除|
|TRACE|リクエストのループバックテスト|
|OPTIONS|サポートされているメソッドの取得|
|CONNECT|プロキシ経由でのリソースの接続|
|PATCH|リソースの部分的な更新|

### 安全なリクエスト（サーバーのステートを変更しないリクエスト）
- GET
- HEAD
- OPTIONS
- TRACE

### 冪等なリクエストメソッド
- 同じリクエストを何度も送信しても、同じ結果になるメソッド
- 安全なメソッドは、定義的に、冪等なメソッドでもある
- 安全でない冪等なメソッド
    - PUT
    - DELETE

## HTMLがサポートしているメソッド
- GET
- POST

## ブラウザ（クライアント）がサポートしているメソッド
- XHR(XMLHttpRequest)を使用することで、すべてのメソッドをサポートしている

## よく使われるHTTPリクエストヘッダー
|ヘッダー|説明|
|:--|:--|
|Accept|クライアントがサーバーに送信できるメディアタイプを指定する|
|Accept-Charset|クライアントがサーバーに送信できる文字セットを指定する|
|Accept-Encoding|クライアントがサーバーに送信できるエンコーディングを指定する|
|Authorization|クライアントがサーバーに送信する認証情報を指定する|
|Cookie|クライアントがサーバーに送信するCookieを指定する|
|Content-Length|クライアントがサーバーに送信するリクエストボディの長さを指定する|
|Content-Type|クライアントがサーバーに送信するリクエストボディのメディアタイプを指定する|
|Host|クライアントがサーバーに送信するホスト名とポート番号を指定する|
|Referer|クライアントがサーバーに送信するリクエストのリファラーを指定する|
|User-Agent|クライアントがサーバーに送信するユーザーエージェントを指定する|

## ステータスコード
|ステータスコード|説明|
|:--|:--|
|1xx|情報提供。すでにリクエストは受理されており、処理を始めていることを伝える|
|2xx|成功。リクエストは受理され、処理が成功したことを伝える|
|3xx|リダイレクト。リクエストは受理され、クライアントが追加の処理が必要であることを伝える|
|4xx|クライアントエラー。リクエストは受理されなかったことを伝える|
|5xx|サーバーエラー。サーバーの不具合でリクエストは受理されたが、サーバーがリクエストを処理できなかったことを伝える|

## よく使われるHTTPレスポンスヘッダー
|ヘッダー|説明|
|:--|:--|
| Allow | サーバーがサポートするメソッドを指定する |
| Content-Encoding | サーバーがレスポンスボディに適用したエンコーディングを指定する |
| Content-Length | サーバーがレスポンスボディに含めたバイト数を指定する |
| Content-Type | サーバーがレスポンスボディに含めたメディアタイプを指定する |
| Date | サーバーがレスポンスを作成した日時を指定する |
| Location | サーバーがリダイレクト先のURLを指定する |
| Server | サーバーが使用しているソフトウェアを指定する |
| Set-Cookie | サーバーがクライアントに送信するCookieを指定する |
| WWW-Authenticate | サーバーがクライアントに送信する認証情報を指定する |

## URI
- Uniform Resource Identifier
- Uniform Resource Locator(URL)とUniform Resource Name(URN)の総称
- URL
    - リソースの場所を示す
- URN
    - リソースの名前を示す
- 書式：`スキーム名://<ユーザー名>:<パスワード>@<ホスト名>:<ポート番号>/<パス>?<クエリ>#<フラグメント>`<br>※階層部=パス+クエリ+フラグメント

## URLの構造体
- `<scheme>://[userinfor@]host/path[?query][#fragment]`
```go
type URL struct {
    Scheme   string
    Opaque   string
    User     *Userinfo
    Host     string
    Path     string
    RawQuery string
    Fragment string
}
```

※リクエストがブラウザからの場合、フラグメントはブラウザによって除去される。

## HTTP/2
- パフォーマンスに重点を置いたプロトコル
- バイナリプロトコル
    - テキストベースのプロトコルよりも、パフォーマンスが高い
- 多重化
    - 1つのTCP接続で複数のリクエストを送信できる
- サーバープッシュ
    - サーバーがクライアントのリクエストに対して、リソースをプッシュできる
- ヘッダー圧縮
    - ヘッダーを圧縮することで、パフォーマンスを向上させる

## Webアプリの構成
1. HTTPを介して、HTTPリクエストメッセージの形でクライアントから入力を受け取る
2. HTTPリクエストメッセージをを処理し、必要な作業を行う
3. HTMLを生成し、HTTPレスポンスメッセージを入れて返す

## Webアプリの構成要素
- ハンドラ
    - HTTPリクエストを受け取り、HTTPレスポンスを返す
    - MVCでいうところのコントローラー
    - 場合によっては、テンプレートエンジンを呼び出して、HTMLを生成し返す
- テンプレートエンジン
    - HTMLに変換できるコード
    - クライアントにHTMLを返すときに使用される
    - 2種類のデザイン哲学
        - 静的テンプレート（ロジックレステンプレート）
            - プレースホルダとなるトークンをHTMLに埋め込む
        - アクティブテンプレート
            - プレースホルダ・トークンのみでなく、条件式、繰り返し文、変数などのロジックも埋め込むことができる


## Goでのハンドラ
- ServerHTTPというメソッドを持ったインタフェースのこと。

## ServerHTTP
- ハンドラのインタフェース
- 引数には、ResponseWriterとRequestを受け取る
    - 第１引数：ResponseWriter
        - レスポンスヘッダーを設定する
        - レスポンスボディを書き込む
    - 第２引数：Request
        - リクエストヘッダーを取得する
        - リクエストボディを読み込む