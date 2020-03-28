# TODONE-server

## セットアップ
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

