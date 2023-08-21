# Goの環境構築

## Goのインストール
```bash
sudo apt install golang-go
```

or

```bash
wget https://golang.org/dl/go1.18.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
```

## アンインストール
```bash
sudo apt remove golang-go
```

## 確認
```bash
go version
```

## パスを通す
```bash
export PATH=$PATH:/usr/local/go/bin
```

## パスを通す(永続化)
```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
```

## 設定すべき環境変数
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export GOROOT=/usr/local/go
```


