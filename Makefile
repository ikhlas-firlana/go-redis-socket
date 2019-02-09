
export PORT=6969
export PORT_REDIS=6379
export HOST_REDIS=localhost
local:
	go run main.go
build: 
	go build -v -o go-redis-socket
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./go-redis-socket
	docker build -t test-go-redis-socket .
