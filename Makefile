help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dbgen: ## sqlboilerによるコード自動生成
	# sqlboilerのインストール
	GO111MODULE=off go get -u github.com/volatiletech/sqlboiler

	# DDL定義を元にコードを自動生成
	sqlboiler mysql -o db/mysql/models --wipe

lint: ## lintの実行
	# golangci-lintのインストール
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

	# pkg配下をチェック。設定は .golangci.yml に記載
	golangci-lint run pkg/...

fmt: ## fmtの実行
	# goimportsのインストール
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

	# fmt,goimportsの実行
	gofmt -s -w pkg/
	goimports -w pkg/

fmt-lint: fmt lint ## fmtとlintの実行

wiregen: ## wire_gen.goの生成
	# google/wireのインストール
	GO111MODULE=off go get -u github.com/google/wire

	# wire genの実行
	wire gen

run: ## APIをビルドせずに立ち上げるコマンド
	go run main.go wire_gen.go

build: ## APIをビルドして立ち上げるコマンド
	go build -o binary/todone
	./binary/todone