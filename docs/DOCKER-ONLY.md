# Dockerのみを使った開発（ローカル環境を汚染しない）

## 概要

ローカル環境に一切ツールをインストールせず、すべてDockerで開発する方法です。

## 前提条件

- Docker
- Docker Compose

のみ必要です。Go、protoc、その他のツールは一切不要です。

## セットアップ

### 1. Protocol Buffers コード生成

```bash
# Dockerを使ってProtocol Buffersコードを生成
make proto-docker
```

これで `grpc/proto/*.pb.go` が生成されます。

### 2. すべてのサービスを起動

```bash
# すべてのサービスをDockerで起動
docker-compose up -d
```

## 開発フロー

### テストを実行

```bash
# アプリケーションコンテナに入る
docker-compose exec app sh

# コンテナ内でテスト実行
go test -v ./...
```

### アプリケーションを起動

```bash
# docker-compose.ymlで自動起動
docker-compose up -d app

# ログ確認
docker-compose logs -f app
```

### Protocol Buffers を再生成

```bash
# ローカルで実行（protoc不要）
make proto-docker

# または、コンテナ内で実行
docker-compose exec app make proto
```

## コマンド一覧

| コマンド | 説明 |
|---------|------|
| `make proto-docker` | Protocol Buffers生成（ローカル環境不要） |
| `docker-compose up -d` | すべてのサービス起動 |
| `docker-compose exec app sh` | アプリコンテナに入る |
| `docker-compose logs -f app` | ログ確認 |
| `docker-compose down` | すべて停止 |

## メリット

✅ **ローカル環境がクリーン**: Go、protoc、その他のツールをインストール不要
✅ **再現性**: 全員が同じDocker環境で開発
✅ **簡単なクリーンアップ**: `docker-compose down -v` で完全削除

## デメリット

❌ **速度**: ローカルより若干遅い
❌ **エディタ統合**: IDEの補完が効きにくい場合がある

## 推奨する使い方

**ハイブリッド**:
- **Protocol Buffers生成**: Docker (`make proto-docker`)
- **日常開発**: mise（ローカル環境）
- **CI/CD**: Docker

これにより、ローカル環境はmiseで管理し、protocだけはDockerを使うことで、最小限の汚染で済みます。
