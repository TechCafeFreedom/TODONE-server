# TODONE-server

## セットアップ
### 【sqlboiler】
1. sqlboiler_example.tomlをコピーして、同じ階層にsqlboiler.tomlを作成する。
2. 自分の開発環境(localhost)に合わせて `user`と`pass`を編集する
例）user: `root` / pass: `password`の場合

```toml
[mysql]
  dbname = "todone"
  host   = "localhost"
  port   = 3306
  user   = "root"
  pass   = "password"
  sslmode= "false"
```

### 【環境変数】
1. .env_exampleをコピーして、同じ階層に.envを作成する
2. 自分の開発環境(localhost)に合わせて`MYSQL_USER`と`MYSQL_PASSWORD`を修正

```.env
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_HOST=localhost
```

## CloudSQLへの接続
unixドメインソケットを使って接続する必要がある。値は全てSecretManagerで管理。詳しくはHackMDのドキュメントを参考にしてください。
```
// dataSourceNameは以下のように指定
{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_PROTOCOL}({MYSQL_INSTANCE})/{MYSQL_DB}
```

## CI/CDのフローについて
- StaticCheck(fmt, lint)/Test/Build
    - 全プルリク
- Deploy
    - masterブランチへのpush

