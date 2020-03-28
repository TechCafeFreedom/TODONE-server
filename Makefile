help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## APIをビルドせずに立ち上げるコマンド
	go run main.go wire_gen.go

build: ## APIをビルドして立ち上げるコマンド
	go build -o binary/todone
	./binary/todone