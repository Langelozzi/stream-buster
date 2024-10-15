.PHONEY: run

run:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/swag/cmd/swag
	swag init
	go mod tidy
	go run main.go