protoroot:
	rm -R pb/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/user/*.proto
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/role/*.proto
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
evans:
	evans --host localhost --port 9001 -r repl
server:
	go run main.go
migrateup:
	migrate -path config/migrate -database "mysql://denies:Semangat45!@tcp(127.0.0.1:3306)/jhapi-staging" --verbose up
migratedown:
	migrate -path config/migrate -database "mysql://denies:Semangat45!@tcp(127.0.0.1:3306)/jhapi-staging" --verbose down