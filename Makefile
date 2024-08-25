.PHONY:
up:
	docker compose up
down:
	docker compose down
run:
	go build ./cmd/carda/main.go
	./main
test:
	go clean -testcache
	go test -v ./internal/carda/repo/