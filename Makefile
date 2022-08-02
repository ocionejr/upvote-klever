test:
	go test -v -cover ./...

server:
	cd server && go run .

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: test server proto evans
