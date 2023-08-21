# セッションを実装する

## Goでセッションを実装する方法
- 自分で実装する
- サードパーティのライブラリを使う
    - gorilla/sessions
    - beego/session
    - go-session/session

## 自分で実装する場合の流れ
- セッションのテーブルを作成する
- セッションをCRUDする
- セッションをCookieに保存する

## セッションの構造体
```go
type Session struct {
    ID       int
    UUID     string
    email    string
    ExpiresAt time.Time
}
```

## Cookieに保存する
```go
func (s *Session) Save(w http.ResponseWriter) error {
    value := map[string]string{
        "uuid": s.UUID,
    }
    encoded, err := securecookie.EncodeMulti(s.Name, value, s.Codecs...)
    if err != nil {
        return err
    }
    cookie := &http.Cookie{
        Name:     s.Name,
        Value:    encoded,
        Path:     s.Path,
        Domain:   s.Domain,
        MaxAge:   s.MaxAge,
        Secure:   s.Secure,
        HttpOnly: s.HttpOnly,
    }
    http.SetCookie(w, cookie)
    return nil
}
```

## Cookieから読み込む
```go
func (s *Session) Load(r *http.Request) error {
    cookie, err := r.Cookie(s.Name)
    if err != nil {
        return err
    }
    value := make(map[string]string)
    if err := securecookie.DecodeMulti(s.Name, cookie.Value, &value, s.Codecs...); err != nil {
        return err
    }
    s.UUID = value["uuid"]
    return nil
}
```