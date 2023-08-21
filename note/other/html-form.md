# HTMLフォームとGo言語

# HTMLフォームの仕組み
- HTMLフォームのデータは常にkey-valueのペアで送信される

## post body内での形式
- HTMLフォームのコンテンツタイプで指定する
- よく使われるコンテンツタイプ
    - application/x-www-form-urlencoded
        - URLエンコーディングで送信する場合に使う
        - `&`でkey-valueを区切る
        - `=`でkeyとvalueを区切る
        - デフォルト
        - key=value&key=value
    - multipart/form-data（MIME）
        - ファイルを送信する場合に使う
        - それぞれのペアが、コンテンツタイプとコンテンツディスポジションを持つ
            - コンテンツタイプ
                - コンテンツの形式を指定できる
                - text/plain
                - image/jpeg
                - image/png
                - application/pdf
                - application/json ...etc
            - コンテンツディスポジション
                - サーバーがクライアントに対してレスポンスを返す際に、どのようにコンテンツ（ファイルなどのデータ）を表示または処理すべきかを指示しするためのモノ
                - ファイル名などを指定できる。
        - ファイルのバイナリデータを送信するため、base64エンコードを行う
    - text/plain
        - テキストのみを送信する場合に使う
        - key=value&key=value

## application/x-www-form-urlencoded v.s. multipart/form-data
- application/x-www-form-urlencoded
    - 単純なテキストの送信に使う
    - 単純で効率がよく処理が少なく済むから
- multipart/form-data
    - ファイルのアップロードのように大量のデータを送信する場合に使う
    - ファイルのバイナリデータを送信するため、base64エンコードを行ったりする

## MIME
- Multipurpose Internet Mail Extensions
- 電子メールやHTTPなどのプロトコルで、さまざまな種類のデータを正しく表現・伝達するための仕組み
- メディアタイプと呼ばれる種別とサブタイプの組み合わせで表現される
    - text/plain
    - text/html
    - image/jpeg
    - image/png
    - application/pdf
    - application/json ...etc
- 構成要素
    - Content-Type
        - 主にデータを表すための型
    - Content-Transfer-Encoding
        - データのエンコード方法を示すためのヘッダー
        - Base64やQuoted-Printableなどがある