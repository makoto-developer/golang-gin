# Golang/Gin - シンプルな通信プロトコル網羅プロジェクト

GinフレームワークとgRPCを使った、シンプルなAPIサーバーテンプレートプロジェクトです。

## 特徴

### ✅ 実装済み機能

#### 受信 (Inbound)
- ✅ **HTTP/REST API** - Ginフレームワーク
- ✅ **gRPC** - Protocol Buffers

#### 送信 (Outbound)
- ✅ **HTTP Client** - 外部APIクライアント実装済み
- ✅ **gRPC Client** - gRPCクライアント実装済み
- ✅ **RabbitMQ Publisher** - メッセージキュー送信実装済み
- ✅ **メール送信** - SMTP経由実装済み

#### モックサーバー (Docker Compose)
- ✅ **HTTP Mock Server** - 外部API モック (`:17002`)
- ✅ **gRPC Mock Server** - 外部gRPCサービス モック (`:17003`)
- ✅ **RabbitMQ Consumer Mock** - メッセージキュー受信モック
- ✅ **MailHog** - メール送信テスト用 (`:17008`)

#### Goroutine活用
- ✅ HTTP/gRPC 並行サーバー起動
- ✅ Graceful Shutdown

#### Ginフレームワーク機能
- ✅ カスタムミドルウェア (Logger, CORS, Recovery)
- ✅ グループルーティング (/api/v1)
- ✅ JSONレスポンス
- ✅ ヘルスチェックエンドポイント

## ディレクトリ構成

```
golang-gin/
├── main.go              # エントリーポイント (HTTP + gRPC同時起動)
├── handlers/            # HTTPハンドラー
│   ├── album.go
│   └── health.go
├── grpc/               # gRPC実装
│   ├── server.go       # gRPCサーバー実装
│   ├── client.go       # gRPCクライアント実装
│   └── proto/
│       └── album.proto # Protocol Buffers定義
├── models/             # データモデル
│   └── album.go
├── middleware/         # Ginミドルウェア
│   ├── logger.go
│   └── cors.go
├── clients/            # 外部通信クライアント
│   ├── http.go         # HTTPクライアント
│   ├── rabbitmq.go     # RabbitMQクライアント
│   ├── mail.go         # メールクライアント
│   ├── *_test.go       # 統合テスト
├── mocks/              # モックサーバー
│   ├── http-mock/      # HTTP APIモック
│   ├── grpc-mock/      # gRPC APIモック
│   └── rabbitmq-consumer/  # RabbitMQコンシューマーモック
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## セットアップ

### 前提条件

- **mise** - ツールバージョン管理（推奨）
- Docker & Docker Compose

### クイックスタート（推奨）

#### 方法1: mise（ローカル環境をクリーンに保つ）

```bash
# 1. mise がない場合はインストール
./scripts/install-mise.sh

# 2. 開発環境をセットアップ（自動で全て実行）
./scripts/setup.sh

# これで以下が完了します：
#   - Go 1.25.5 インストール (mise)
#   - protoc 30.2 インストール (mise)
#   - protoc-gen-go, protoc-gen-go-grpc インストール
#   - go mod tidy
#   - Protocol Buffers コード生成
```

#### 方法2: Docker（ローカルに一切インストールしない）

```bash
# Protocol Buffers コード生成のみDockerで実行
make proto-docker

# または、すべてDocker Composeで実行
docker-compose up -d
docker-compose exec app sh
```

### 手動セットアップ

```bash
# 1. mise でツールをインストール
mise install

# 2. Go protoc プラグインをインストール
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 3. 依存関係の解決
go mod download
go mod tidy

# 4. Protocol Buffers コード生成
make proto
```

### ローカル起動

```bash
# アプリケーション起動
go run main.go

# または
make run
```

起動すると以下のサーバーが立ち上がります:
- **HTTP Server**: http://localhost:17000
- **gRPC Server**: localhost:17001

### Docker Composeで起動

```bash
# .env.exampleをコピー（初回のみ）
cp .env.example .env

# .envファイルを編集して環境を設定（オプション）
# ENV=dev  # dev, staging, prod のいずれか

# すべてのサービスを起動
docker-compose up -d

# ログ確認
docker-compose logs -f app

# コンテナ一覧確認
docker-compose ps

# 停止
docker-compose down
```

**環境別の起動**:
```bash
# 開発環境（デフォルト）
ENV=dev docker-compose up -d

# ステージング環境
ENV=staging docker-compose up -d

# 本番環境
ENV=prod docker-compose up -d
```

### コンテナ名とホスト名（環境別）

すべてのサービスには**環境名を含む**明示的なコンテナ名とホスト名が設定されています：

| サービス | コンテナ名 | ホスト名 | 備考 |
|---------|-----------|----------|------|
| app | golang-gin-dev-app | golang-gin-dev-app | 環境名が含まれる |
| postgres | golang-gin-dev-postgres | golang-gin-dev-postgres | ユーザー: golang_gin_dev |
| rabbitmq | golang-gin-dev-rabbitmq | golang-gin-dev-rabbitmq | ユーザー: golang_gin_dev |
| http-mock | golang-gin-dev-http-mock | golang-gin-dev-http-mock | - |
| grpc-mock | golang-gin-dev-grpc-mock | golang-gin-dev-grpc-mock | - |
| rabbitmq-consumer | golang-gin-dev-rabbitmq-consumer | golang-gin-dev-rabbitmq-consumer | - |
| mailhog | golang-gin-dev-mailhog | golang-gin-dev-mailhog | - |

**環境設定**:
- **環境名**: `.env`の`ENV`変数で設定（dev, staging, prod）
- **プロジェクト名**: `golang-gin-dev` （`.env`の`COMPOSE_PROJECT_NAME`）
- **DBユーザー名**: `golang_gin_dev` （環境名が含まれる）
- **RabbitMQユーザー名**: `golang_gin_dev` （環境名が含まれる）

**命名規則**: `{プロジェクト名}-{環境名}-{サービス名}`

これにより、以下のメリットがあります：
- 開発・ステージング・本番環境を同時に実行可能
- 環境ごとにユーザー名とデータベースが分離される
- コンテナ名が予測可能で管理しやすい
- コンテナ間通信でホスト名を使用できる

**環境切り替え例**:
```bash
# 開発環境（デフォルト）
ENV=dev docker-compose up -d

# ステージング環境
ENV=staging docker-compose up -d

# 本番環境
ENV=prod docker-compose up -d
```

## Docker Compose サービス一覧

| サービス | コンテナ名 (dev環境) | ポート | 説明 |
|---------|---------------------|--------|------|
| app | golang-gin-dev-app | 17000, 17001 | メインアプリケーション (HTTP + gRPC) |
| http-mock | golang-gin-dev-http-mock | 17002 | 外部API モックサーバー |
| grpc-mock | golang-gin-dev-grpc-mock | 17003 | 外部gRPCサービス モックサーバー |
| mailhog | golang-gin-dev-mailhog | 17007 (SMTP), 17008 (Web UI) | メール送信テスト用 |
| rabbitmq | golang-gin-dev-rabbitmq | 17005 (AMQP), 17006 (Management) | メッセージキュー |
| rabbitmq-consumer | golang-gin-dev-rabbitmq-consumer | - | RabbitMQコンシューマーモック |
| postgres | golang-gin-dev-postgres | 17004 | PostgreSQL データベース |

**注**: 上記はdev環境のコンテナ名です。staging/prod環境では`dev`の部分が変わります。

## API仕様

### HTTP REST API

#### ヘルスチェック
```bash
curl http://localhost:17000/health
```

#### 全アルバム取得
```bash
curl http://localhost:17000/api/v1/albums
```

#### アルバムID指定取得
```bash
curl http://localhost:17000/api/v1/albums/1
```

#### アルバム作成
```bash
curl http://localhost:17000/api/v1/albums \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"id": "4","title": "only my railgun","artist": "FripSide","price": 30.2, "tax": 0.1}'
```

### gRPC API

gRPCクライアントの使用例は `grpc/client.go` を参照してください。

```go
import "golang-gin/grpc"

// クライアント作成
client, err := grpc.NewClient("localhost:17001")
defer client.Close()

// 全アルバム取得
albums, err := client.GetAlbums()

// ID指定取得
album, err := client.GetAlbumByID("1")

// 新規作成
newAlbum, err := client.CreateAlbum("4", "Title", "Artist", 29.99, 0.1)
```

## モックサーバーの使い方

### HTTP Mock Server

外部APIのモックサーバーが `:17002` で起動します。

```bash
# モックAPIにリクエスト
curl http://localhost:17002/api/v1/users

# アプリからモックサーバーを使用
import "golang-gin/clients"

httpClient := clients.NewHTTPClient("http://localhost:17002")
data, err := httpClient.Get("/api/v1/users")
```

**モックエンドポイント**:
- `GET /api/v1/users` - ユーザー一覧
- `GET /api/v1/users/:id` - ユーザー詳細
- `POST /api/v1/users` - ユーザー作成
- `GET /api/v1/products` - 商品一覧
- `GET /api/v1/error` - エラーレスポンス（テスト用）

### gRPC Mock Server

外部gRPCサービスのモックサーバーが `:17003` で起動します。

```bash
# アプリからgRPCモックサーバーを使用
import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

conn, _ := grpc.Dial("localhost:17003", grpc.WithTransportCredentials(insecure.NewCredentials()))
// ProductServiceClient を使用してモックサービスにアクセス
```

**モックサービス**:
- `GetProducts()` - 商品一覧取得（5件のモックデータ）
- `GetProductByID(id)` - 商品詳細取得
- `CreateProduct(...)` - 商品作成

**テスト**:
```bash
# gRPCモックサーバーへの接続テスト
go test ./grpc -run TestGRPCMockServer -v
```

### RabbitMQ

RabbitMQにメッセージを送信すると、コンシューマーモックがログ出力します。

```bash
# コンシューマーのログを確認
docker-compose logs -f rabbitmq-consumer

# アプリからメッセージ送信
import "golang-gin/clients"

client, _ := clients.NewRabbitMQClient("amqp://guest:guest@localhost:17005/")
client.Publish("test_queue", []byte(`{"message":"test"}`))
```

**RabbitMQ Management UI**: http://localhost:17006 (guest/guest)

### MailHog

メール送信をテストできます。送信したメールはWeb UIで確認できます。

```bash
# アプリからメール送信
import "golang-gin/clients"

mailClient := clients.NewMailClient("localhost", "17007", "", "", "noreply@test.com")
mailClient.SendMail([]string{"to@example.com"}, "Test", "Body")
```

**MailHog Web UI**: http://localhost:17008

## テスト

### ユニットテスト（外部サービス不要）

```bash
# すべてのユニットテストを実行
make test-unit

# または個別に実行
go test ./handlers -v      # HTTPハンドラーのテスト
go test ./grpc -v          # gRPCサーバーのテスト
go test ./middleware -v    # ミドルウェアのテスト
go test ./models -v        # モデルのテスト
```

### 統合テスト（モックサーバー必要）

```bash
# 1. モックサーバーを起動
docker-compose up -d

# 2. 統合テストを実行
make test-integration

# または個別に実行
go test ./clients -v              # クライアントのテスト
go test ./integration_test.go -v  # E2Eテスト
```

### すべてのテストを実行

```bash
# モックサーバー起動 + 全テスト実行
docker-compose up -d
make test

# またはシンプルに
go test -v ./...
```

### カバレッジレポート

```bash
# カバレッジ計測 + HTMLレポート生成
make test-coverage

# ブラウザで coverage.html を開く
open coverage.html
```

### テストの構成

| ディレクトリ | テストタイプ | 説明 |
|------------|------------|------|
| handlers/ | ユニット | HTTPハンドラーのテスト |
| grpc/ | ユニット | gRPCサーバーのテスト |
| middleware/ | ユニット | ミドルウェアのテスト |
| models/ | ユニット | データモデルのテスト |
| clients/ | 統合 | 外部通信クライアントのテスト（モック必要） |
| integration_test.go | E2E | フルワークフローテスト |

## 開発コマンド

```bash
# Protocol Buffersからコード生成
make proto

# または Docker を使用（ローカルにprotocインストール不要）
make proto-docker

# アプリケーション起動
make run

# テスト実行
make test

# 生成ファイル削除
make clean

# 依存関係インストール
make deps
```

## Stack

- **Go** v1.25.5 - プログラミング言語
- **Gin** v1.11.0 - HTTPフレームワーク
- **gRPC** v1.76.0 - RPCフレームワーク
- **Protocol Buffers** v1.36.3 - シリアライゼーション
- **RabbitMQ Client** v1.10.0 - AMQPクライアント
- **PostgreSQL** 18 (Docker Compose) - データベース
- **RabbitMQ** 4.2 (Docker Compose) - メッセージブローカー
- **MailHog** - メールテスト
- **protoc** v30.2 - Protocol Buffersコンパイラ

## 今後の実装予定

### 受信プロトコル追加
- [ ] GraphQL (github.com/99designs/gqlgen)

### Gin機能拡張
- [ ] バリデーション
- [ ] ファイルアップロード
- [ ] 静的ファイル配信
- [ ] HTMLテンプレート
- [ ] マルチフォーマットレスポンス (XML, YAML)

### Goroutine活用
- [ ] バックグラウンドワーカー
- [ ] 非同期処理パターン
- [ ] ワーカープール

### その他
- [ ] データベース接続 (GORM)
- [ ] マイグレーション (golang-migrate)
- [ ] ロギング強化
- [ ] Kubernetes対応

## セキュリティ注意事項

### ⚠️ 開発環境用の設定

このプロジェクトは**開発・学習目的**のテンプレートです。本番環境で使用する前に、必ず以下のセキュリティ対策を実施してください。

### 🔒 本番環境への移行前チェックリスト

#### 1. 認証情報の変更

**絶対に変更が必要**:
- ✅ PostgreSQLのパスワード（デフォルト: `postgres`）
- ✅ RabbitMQのユーザー名とパスワード（デフォルト: `guest/guest`）
- ✅ すべてのデフォルト認証情報

#### 2. 環境変数の管理

```bash
# .env.exampleをコピーして.envを作成
cp .env.example .env

# .envファイルを編集して本番用の値を設定
# NEVER commit .env to version control!
```

**.gitignoreで保護されているファイル**:
- `.env`
- `.env.local`
- `.envrc`

これらのファイルは**絶対にGitにコミットしないでください**。

#### 3. Docker Composeの設定変更

`docker-compose.yml`は環境変数を使用するように設定されています:

```yaml
# ✅ 環境変数から読み込む（デフォルト値あり）
environment:
  POSTGRES_USER: ${POSTGRES_USER:-golang_gin_dev}
  POSTGRES_PASSWORD: ${POSTGRES_PASSWORD:-postgres_dev}
  POSTGRES_DB: ${POSTGRES_DB:-golang_gin_dev}
```

**本番環境では`.env`ファイルで以下を変更**:
```bash
ENV=prod
COMPOSE_PROJECT_NAME=golang-gin-prod
POSTGRES_USER=golang_gin_prod
POSTGRES_PASSWORD=your_secure_password_here
POSTGRES_DB=golang_gin_prod
RABBITMQ_USER=golang_gin_prod
RABBITMQ_PASSWORD=your_secure_password_here
```

#### 4. ネットワークセキュリティ

- ✅ 必要なポートのみを公開
- ✅ ファイアウォールの設定
- ✅ HTTPS/TLSの有効化
- ✅ gRPCの認証・暗号化

#### 5. その他のセキュリティ対策

- ✅ CORS設定の見直し（現在は `Access-Control-Allow-Origin: *`）
- ✅ レート制限の実装
- ✅ 入力値のバリデーション強化
- ✅ SQLインジェクション対策（ORMの使用）
- ✅ ログに機密情報を含めない
- ✅ 定期的な依存関係の更新（`go mod tidy`, セキュリティパッチ）

### 📋 セキュリティチェックコマンド

```bash
# 機密情報が誤ってコミットされていないかチェック
git log --all --full-history -- .env
git log --all --full-history -- .envrc

# .gitignoreが正しく機能しているか確認
git check-ignore .env .envrc

# 依存関係の脆弱性スキャン（推奨）
go list -json -m all | docker run --rm -i sonatypecommunity/nancy:latest sleuth
```

### 🚨 本番環境では絶対に避けること

- ❌ デフォルトの認証情報を使用
- ❌ `.env`ファイルをGitにコミット
- ❌ `GIN_MODE=debug`のまま運用
- ❌ すべてのポートを公開
- ❌ CORS設定で`*`を許可

## 参考資料

- [Go公式チュートリアル](https://golang.org/doc/tutorial/web-service-gin)
- [Ginドキュメント](https://gin-gonic.com/docs/)
- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [RabbitMQ Tutorials](https://www.rabbitmq.com/getstarted.html)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
