SOURCE_FILE := $(notdir $(source))
SOURCE_DIR := $(dir $(source))
MOCK_FILE := mock_${SOURCE_FILE}
MOCK_DIR := ${SOURCE_DIR}mock_$(lastword $(subst /, ,${SOURCE_DIR}))/
MOCK_TARGET := $(lastword $(subst /, ,${SOURCE_DIR}))


help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

mockgen: # mockgenの実行
	# Usege: make mockgen source=<インターフェースの定義しているファイル>

	# mockgenのインストール
	go install github.com/golang/mock/mockgen

	# mockgenの実行
	mockgen -source ${SOURCE_DIR}${SOURCE_FILE} -destination ${MOCK_DIR}${MOCK_FILE}

dbgen: ## sqlboilerによるコード自動生成
	# sqlboilerのインストール
	GO111MODULE=off go get -u github.com/volatiletech/sqlboiler

	# DDL定義を元にコードを自動生成
	sqlboiler mysql -o db/mysql/model -p model --wipe db/mysql/model

wiregen: ## wire_gen.goの生成
	# google/wireのインストール
	GO111MODULE=off go get -u github.com/google/wire

	# wire genの実行
	wire gen cmd/wire.go

test: ## testの実行
	go test -v ./...

lint: ## lintの実行
	# golangci-lintのインストール
	GO111MODULE=on go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

	# pkg配下をチェック。設定は .golangci.yml に記載
	golangci-lint run

fmt: ## fmtの実行
	# goimportsのインストール
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

	# fmt,goimportsの実行
	gofmt -s -w pkg/
	goimports -w pkg/

fmt-lint: fmt lint ## fmtとlintの実行

run: ## APIをビルドせずに立ち上げるコマンド
	go run ./cmd

build: ## APIをビルドして立ち上げるコマンド
	go build -o binary/todone ./cmd
	./binary/todone