tidy:
	go mod tidy

run:
	go run main.go

run-prod:
	export GIN_MODE=release && make run