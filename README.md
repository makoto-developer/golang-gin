# Golang/Gin

Ginの練習

## Feature

|checked|feature|note|
|:---|:---|:---|
|🔲|テーマを決める||
|🔲|ディレクトリ構成を決定||
|🔲|grpc|
|🔲|データベースと接続||
|🔲|GORM||
|🔲|logger||
|🔲|kubernetes||
|🔲|cobraでサービス起動||
|🔲|マイグレーション機能|使うならgolang-migrate辺りか|
|🔲|テストコード||
|🔲|フォーマッター||
|🔲|||
|🔲|||

## Stack

TODO miseで管理したい

- go v1.21.1
- gin v1.19.1

## Starting

```zsh
go run .
```

全てのアルバムリスト取得


```zsh
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

アルバムを追加

```zsh
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "only my railgun","artist": "FripSide","price": 30.2, "tax": 0.1}'
```

アルバムIDからアルバムの詳細を取得

```zsh
curl http://localhost:8080/albums/1 \
    --header "Content-Type: application/json" \
    --request "GET"
```

## Notes

`go.mod`のGoのバージョンを上げる

```shell
go mod tidy -go=1.21.1
```
