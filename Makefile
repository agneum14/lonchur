build:
	templ generate
	go build -o tmp cmd/lonchur.go

run:
	go run cmd/lonchur.go
