# 開発環境セットアップガイド

## 概要

このプロジェクトは**mise**を使ってツールバージョンを管理しています。

## 必要なツール

以下は`.tool-versions`で自動管理されます：

| ツール | バージョン | 用途 |
|--------|-----------|------|
| Go | 1.21.1 | アプリケーション開発 |
| protoc | 29.4 | Protocol Buffers コンパイラ |

**利用可能なprotocバージョン**:
```bash
mise ls-remote protoc  # 29.4, 29.5, 30.x, 31.x, 32.x, 33.x など
```

## セットアップ手順

### 1. mise のインストール

#### macOS (Homebrew使用)
```bash
brew install mise
```

#### macOS/Linux (curl使用)
```bash
curl https://mise.run | sh
```

#### またはスクリプト使用
```bash
./scripts/install-mise.sh
```

### 2. Shell設定

mise を有効化するために、シェルの設定ファイルに以下を追加：

**Zsh (.zshrc)**:
```bash
eval "$(mise activate zsh)"
```

**Bash (.bashrc)**:
```bash
eval "$(mise activate bash)"
```

設定後、シェルを再起動：
```bash
source ~/.zshrc  # または source ~/.bashrc
```

### 3. プロジェクトセットアップ

**自動セットアップ（推奨）**:
```bash
./scripts/setup.sh
```

このスクリプトは以下を自動実行します：
1. mise でツールをインストール (`go`, `protoc`)
2. Go protoc プラグインをインストール
3. Go依存関係をダウンロード
4. Protocol Buffers コード生成

**手動セットアップ**:
```bash
# ツールインストール
mise install

# protoc プラグイン
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 依存関係
go mod download
go mod tidy

# Protocol Buffers生成
make proto
```

## 確認

セットアップが成功したか確認：

```bash
# バージョン確認
go version       # go version go1.21.1
protoc --version # libprotoc 25.1

# テスト実行
make test-unit
```

## トラブルシューティング

### `protoc: command not found`

```bash
# mise プラグインをインストール
mise plugins install protoc

# mise を使ってインストール
mise install protoc

# または Docker を使用（ローカルインストール不要）
make proto-docker
```

### `protoc-gen-go: program not found`

```bash
# プラグインを再インストール
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# PATHに追加（必要な場合）
export PATH="$PATH:$(go env GOPATH)/bin"
```

### `package golang-gin/grpc/proto is not in std`

```bash
# Protocol Buffersコードを生成
make proto
```

### `missing go.sum entry`

```bash
# 依存関係を再解決
go mod tidy
```

## 環境変数

開発に便利な環境変数を設定できます：

```bash
# .envrc.example をコピー
cp .envrc.example .envrc

# direnv を使う場合（オプション）
direnv allow
```

## 次のステップ

セットアップ完了後：

1. **モックサーバー起動**: `docker-compose up -d`
2. **テスト実行**: `make test-unit`
3. **アプリ起動**: `make run`

詳細は [README.md](../README.md) を参照してください。
