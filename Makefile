.PHONY: build test fmt

build:
	go build ./...

test:
	go test ./...

aider:
	aider \
  --model openrouter/deepseek/deepseek-chat \
  --weak-model openrouter/deepseek/deepseek-chat \
  --map-tokens 1024
fmt:
	goimports -w .

check:
	go build ./...
	go test ./...

lint: