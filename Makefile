make run:
	go run ./cmd/main/.

test-service:
	go test ./internal/service -cover -v

test-rest:
	go test ./internal/rest -cover -v

gen-mock:
	mockgen -source=internal/interfaces/storage.go -destination=internal/storage/mocks/service.go

test-cover:
	@go test -coverprofile cover.out ./... -covermode atomic
	@go tool cover -html=cover.out