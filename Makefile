.PHONY: tidy run clean

tidy:
	go mod tidy

run:
	go run cmd/main.go

# clean the chache to reinstall dependencies, make sure after clean run tidy
clean:
	go clean -modcache
