# ─── ビルドステージ ─────────────────────────────────────────────
FROM golang:1.24-alpine AS builder

# 必要なパッケージ
RUN apk add --no-cache git

WORKDIR /app

# 依存関係だけ先に取得（キャッシュ利用のため）
COPY go.mod go.sum ./
RUN go mod download

# ソースコピー＆ビルド
COPY . .
RUN GOOS=linux GOARCH=amd64 \
    go build -o /usr/local/bin/server ./cmd/server/main.go

# ─── ランタイムステージ ────────────────────────────────────────
FROM alpine:latest

# 証明書等必要なら
RUN apk add --no-cache ca-certificates

# ビルド成果物をコピー
COPY --from=builder /usr/local/bin/server /usr/local/bin/server

# 環境変数（必要なら）
# ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["server"]
