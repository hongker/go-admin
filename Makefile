startup:
	sh scripts/startup.sh

test: startup
	go test -timeout 9000s -cover -a -v ./...
