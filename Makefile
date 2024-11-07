.PHONY tidy run

tidy:
	go mod tidy
run:
	go run cmd/main.go