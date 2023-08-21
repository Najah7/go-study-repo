# Postgreのインストール
1. `sudo apt-get install postgresql postgresql-contrib`でインストール

# Postgreの設定
1. `sudo -u postgres`でpostgresユーザーに切り替え
1. `createuser -P -s -e <username>`でユーザーを作成（create -interactive）
1. `createdb <username>`でデータベースを作成

※`psql`でPostgreにログイン
