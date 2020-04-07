# TODONE-server

## CloudSQLへの接続
unixドメインソケットを使って接続する必要がある。値は全てSecretManagerで管理。詳しくはHackMDのドキュメントを参考にしてください。
```
// dataSourceNameは以下のように指定
{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_PROTOCOL}({MYSQL_INSTANCE})/{MYSQL_DB}
```
