## 一般的なWebサービス

### SOAPベース(※元はSimple Object Access Protocol。名が体を表さなくなったので、現在は、ただのSOAP)
    - XMLで定義された構造化データを交換するためのプロトコル
    - 主に企業の内部システム連携などに使われることが多い
    - 機能駆動型（機能を選択することでアプリケーションが動作する）
    - メッセージのフォーマットのみに関する考え方
    - RPC(Remote Procedure Call)をベースにしたプロトコル
    - 利点
        - 歴史が長く、ドキュメントも豊富。
        - W3Cのワーキンググループによって標準化されている。
        - 企業によるサポートが充実している。
        - 拡張機能も豊富。
        - 堅牢でWSDL（Web Service Description Language）を用いて、サービスの仕様を記述することができる。
        - エラー処理も組み込まれている。
        - UDDI(Universal Description, Discovery and Integration)という標準化されたサービスディレクトリがある。
    - 欠点
        - 仕様が大きく、不必要に複雑になっている。
        - XMLのメッセージは項目が多い。
        - トラブルシューティングが難しい。なので、解析ツールを必要とすることもしばしば。
        - メッセージのサイズが大きい。なので、処理が重くなりがち。
        - 変更があるたびにWSDLを更新する必要がある。これにより、変更を嫌がり、バージョンを固定しがちになったり。。。


## 構造体タグ
- 構造体のフィールドにタグを付けることができる
- タグは、構造体のフィールドのメタデータを表す
- タグの書き方：key:"value"
    - `json:"name"`
    - `db:"name"`
- タグの使い方
    - データベースのカラム名を指定する
    - JSONのキー名を指定する
    - keyには、「空白」、「カンマ」、「二重引用符」を含めることができない
    - 構造体タグに「`」を使う理由、構造体タグで「"」使う＆「'」はint32型の多バイト文字列にしか使われないので。
        - 「"」を使った文字はstring型
        - 「'」を使った文字はrune型
        - 「`」を使った文字はstring型(row string)
- タグの使い方の例
    - `json:"name"`：JSONのキー名をnameにする
    - `db:"name"`：データベースのカラム名をnameにする
    - `json:"name" db:"name"`：JSONのキー名をnameにし、データベースのカラム名もnameにする

## XML構造体タグが従うべき規則
- XML要素各自体を保存するには（通常は構造体が要素名になる）、XMLNameフィールドという名前でxml.Name型のフィールドを追加する。要素名がこのフィールドに保存される。
- XML要素の属性を保存するには、属性と同じ名前のフィールドを定義し、`xml:<属性名>,attr`というタグを付ける。
- XML要素の文字データを保存するには、XML要素タグと同様の名前のフィールドを定義し、`xml:",chardata"`というタグを付ける。
- XML要素から未処理のままで（生の）XMLを得るには、フィールドを定義し（名前は任意）、`xml:",innerxml"`というタグを付ける。
- モードフラグ（,attr、,chardata、,innerxml）がない場合、構造体のフィールドは構造体と同じ名前のXML要素にマッピングされる。
- 木構造を追ってXML要素に到達するのではなく、要素を直接取得したい場合は、構造体タグに`xml:"a>b>c"`を使う。

#### SOAPのサンプルコード(Commentを取得する場合)
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope
    xmlns:soap="http://www.w3.org/2003/05/soap-envelope/"
    soap:encodingStyle="http://www.w3.org/2003/05/soap-encoding">
    <soap:Body>
        
        <m:GetCommentRequest xmlns:m="http://example.com/comment">
            <m:CommentId>1</m:CommentId>
        </m:GetCommentRequest>
    </soap:Body>
</soap:Envelope>
```

- WSDL
```xml
<?xml version="1.0" encoding="UTF-8"?>
<definitions name="sampleApp"
    targetNamespace="http://example.com/comment"
    xmlns:tns="http://example.com/comment"
    xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/"
    xmlns="http://schemas.xmlsoap.org/wsdl/">
<message name="GetCommentRequest">
    <part name="CommentId" type="xsd:int"/>
</message>
<protType name="GetCommentPortType">
    <operation name="GetComment">
        <input message="tns:GetCommentRequest"/>
        <output message="tns:GetCommentResponse"/>
    </operation>
</protType>
<binding name="GetCommentBinding" type="tns:GetCommentPortType">
    <soap:binding style="document" transport="http://schemas.xmlsoap.org/soap/http"/>
    <operation name="GetComment">
        <soap:operation soapAction="http://example.com/comment/GetComment"/>
        <input>
            <soap:body use="literal"/>
        </input>
        <output>
            <soap:body use="literal"/>
        </output>
    </operation>
</binding>
<service name="GetCommentService">
    <documentation>Returns a comment</documentation>
    <port name="GetCommentPort" binding="tns:GetCommentBinding">
        <soap:address location="http://example.com/comment"/>
    </port>
</service>
</definitions>
```

### Go言語によるXMLのパースと生成

#### XMLのパース
```go
package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

type Post struct {
    XMLName xml.Name `xml:"post"`　// XMLの要素名を指定
    Id string `xml:"id,attr"`
    Content string `xml:"content"`
    Author Author `xml:"author"`
    Xml string `xml:",innerxml"`　// XMLの文字列を取得。innerxmlは、xmlのタグのテキストに対応している
    Comments []Comment `xml:"comments>comment"`　// comments要素の子要素のcomment要素を取得。>は、子要素を取得することを表している
}

type Author struct {
    Id string `xml:"id,attr"` // 属性を取得。attrは、xmlのタグの属性に対応している
    Name string `xml:",chardata"`　// 要素のテキストを取得。chardataは、xmlのタグのテキストに対応している
}

type Comment struct {
    Id string `xml:"id,attr"`
    Content string `xml:"content"`
    Author Author `xml:"author"`
}

func main() {
    xmlFile, err := os.Open("post.xml")
    if err != nil {
        fmt.Println("Error opening XML file:", err)
        return
    }
    defer xmlFile.Close()

    xmlData, err := ioutil.ReadAll(xmlFile)
    if err != nil {
        fmt.Println("Error reading XML data:", err)
        return
    }

    var post Post
    xml.Unmarshal(xmlData, &post)　// XMLを構造体に変換
    fmt.Println(post)
}
```


#### XMLの生成
```go

package main

import (
    "encoding/xml"
    "fmt"
    "os"
)

type Post struct {
    XMLName xml.Name `xml:"post"`
    Id string `xml:"id,attr"`
    Content string `xml:"content"`
    Author Author `xml:"author"`
    Xml string `xml:",innerxml"`
    Comments []Comment `xml:"comments>comment"`
}

type Author struct {
    Id string `xml:"id,attr"`
    Name string `xml:",chardata"`
}

type Comment struct {
    Id string `xml:"id,attr"`
    Content string `xml:"content"`
    Author Author `xml:"author"`
}

func main() {
    post := Post{
        Id: "1",
        Content: "Hello World!",
        Author: Author{
            Id: "2",
            Name: "Sau Sheong",
        },
        Comments: []Comment{
            Comment{
                Id: "1",
                Content: "Have a great day!",
                Author: Author{
                    Id: "3",
                    Name: "Adam",
                },
            },
            Comment{
                Id: "2",
                Content: "How are you today?",
                Author: Author{
                    Id: "4",
                    Name: "Betty",
                },
            },
        },
    }

    xmlFile, err := os.Create("post.xml")
    if err != nil {
        fmt.Println("Error creating XML file:", err)
        return
    }
    encoder := xml.NewEncoder(xmlFile)
    encoder.Indent("", "    ")　// 第1引数：インデントの文字列(文頭に追加するテキスト)、第2引数：インデントの階層(インデントの際に使用する空白文字列)
    err = encoder.Encode(&post)
    if err != nil {
        fmt.Println("Error encoding XML to file:", err)
        return
    }
}
```


### RESTベース（REpresentational State Transfer）
    - 相互通信するプログラムをデザインするときに使われる原則
    - デザインにのみ関わる。送られるメッセージには無関係。
    - 標準化された少数のアクションを使ってリソースを操作することで機能を実現する
    - リソースと呼ばれるモデルを公開して、それに対して小数のアクションを許可する
    - WADL(Web Application Description Language)を用いて、サービスの仕様を記述することができる。
        - OpenAPI(Swagger)
        - RAML(Restful API Modeling Language)
        - JSON-home
    - 主に一般に公開されているWebサービスはRESTベースが多い
    - データ駆動型（データ中心にアプリケーションが組み立てられる）
    - HTTPをベースにしたプロトコル(HTTPメソッドを使って、リソースを操作する)
    - 特徴
        - 柔軟
        - 軽量
        - シンプル
    - RESTで複雑な処理を実装する方法
        - 処理を具体化する。アクションを名詞に変え、リソースにする（例：/comment/1/like）
        - アクションをリソースの属性にする（例：/comment/1?like=true or postでkey-valueで送る）

### GoによるJSONのパースと生成

#### JSONのパース
```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Post struct {
    Id int `json:"id"`
    Content string `json:"content"`
    Author Author `json:"author"`
    Comments []Comment `json:"comments"`
}

type Author struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

type Comment struct {
    Id int `json:"id"`
    Content string `json:"content"`
    Author string `json:"author"`
}

func main() {
    jsonFile, err := os.Open("post.json")
    if err != nil {
        fmt.Println("Error opening JSON file:", err)
        return
    }
    defer jsonFile.Close()

    jsonData, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        fmt.Println("Error reading JSON data:", err)
        return
    }

    var post Post
    json.Unmarshal(jsonData, &post)　// JSONを構造体に変換
    fmt.Println(post)
}
```

#### JSONの生成
```go

package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Post struct {
    Id int `json:"id"`
    Content string `json:"content"`
    Author Author `json:"author"`
    Comments []Comment `json:"comments"`
}

type Author struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

type Comment struct {
    Id int `json:"id"`
    Content string `json:"content"`
    Author string `json:"author"`
}

func main() {
    post := Post{
        Id: 1,
        Content: "Hello World!",
        Author: Author{
            Id: 2,
            Name: "Sau Sheong",
        },
        Comments: []Comment{
            Comment{
                Id: 3,
                Content: "Have a great day!",
                Author: Author{
                    Id: 3,
                    Name: "Adam",
                },
            },
            Comment{
                Id: 4,
                Content: "How are you today?",
                Author: Author{
                    Id: 4,
                    Name: "Betty",
                },
            },
        },
    }

    jsonFile, err := os.Create("post.json")
    if err != nil {
        fmt.Println("Error creating JSON file:", err)
        return
    }
    encoder := json.NewEncoder(jsonFile)
    err = encoder.Encode(&post)
    if err != nil {
        fmt.Println("Error encoding JSON to file:", err)
        return
    }
}
```