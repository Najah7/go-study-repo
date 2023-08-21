# テンプレートエンジン

## テンプレートエンジンとは
- テンプレートエンジンとは、テンプレートをもとにして、データを埋め込んで、出力する仕組みのこと
- テンプレートエンジンを使うことで、HTMLのコードを簡潔に記述できる
- テンプレートエンジンは、HTMLだけでなく、メールのテンプレートなどにも使われる

## テンプレートエンジンを使ってHTMLを返す流れ
1. データ＆テンプレート→テンプレートエンジン
1. テンプレートエンジンがHTMLを生成
1. HTMLをブラウザに返す

## html/template

### template.ParseFiles
- template.ParseFilesは、テンプレートをパースする関数
- テンプレートをパースすると、テンプレートエンジンがHTMLを生成するためのデータ構造が作られる
- 引数
    - テンプレートファイルのパス
- 戻り値
    - テンプレートエンジン
```go
template := template.ParseFiles("index.html")
```

### template.Must
- template.Mustは、テンプレートエンジンのエラーをチェックする関数
- テンプレートエンジンのエラーは、テンプレートをパースするときに発生する
- 引数
    - テンプレートエンジン
- 戻り値
    - テンプレートエンジン
```go
template := template.Must(template.ParseFiles("index.html"))
```

### template.Execute
- template.Executeは、テンプレートエンジンがHTMLを生成する関数
- テンプレートエンジンがHTMLを生成するときに、テンプレートエンジンが生成したHTMLを、io.Writerに書き込む
- 引数
    - io.Writer
    - データ
- 戻り値
    - エラー
```go
template.Execute(w, data)
```

### template.ExecuteTemplate
- template.ExecuteTemplateは、テンプレートエンジンがHTMLを生成する関数
- template.ExecuteTemplateは、テンプレートエンジンがHTMLを生成するときに、テンプレートエンジンが生成したHTMLを、io.Writerに書き込む
- 引数
    - io.Writer
    - テンプレート名
    - データ
- 戻り値
    - エラー
```go
template.ExecuteTemplate(w, "index.html", data)
```

### template.New
- template.Newは、テンプレートエンジンを生成する関数
- テンプレートエンジンを生成するときに、テンプレートエンジンが生成したHTMLを、io.Writerに書き込む
- 引数
    - テンプレート名
- 戻り値
    - テンプレートエンジン
```go
template := template.New("index.html")
```

### template.Parse
- template.Parseは、テンプレートエンジンをパースする関数
- テンプレートエンジンをパースすると、テンプレートエンジンがHTMLを生成するためのデータ構造が作られる
- 引数
    - テンプレート文字列
- 戻り値
    - テンプレートエンジン
```go
template := template.Parse("{{.Name}}")
```

### テンプレートの書式
- テンプレートの書式は、HTMLの書式と同じ
- テンプレートの書式には、テンプレートエンジンがHTMLを生成するためのデータ構造を埋め込むことが可能
- 基本的には、{{}}で囲むことで、テンプレートエンジンがHTMLを生成するためのデータ構造を埋め込むことが可能

```html
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>テンプレートエンジン</title>
</head>
<body>
    <h1>{{.Name}}</h1>
</body>
</html>
```



