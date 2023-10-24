DB_URL=mysql://root:ZXCasdqwe123!@tcp(localhost:3306)/mandoor_db

migrate_up:
	migrate -database "$(DB_URL)" -path db/migration -verbose up

migrate_down:
	migrate -database "$(DB_URL)" -path db/migration -verbose down

proto:
	rm -f api/pb/*.go
	rm -f cmd/doc/swagger/*.swagger.json

	protoc --proto_path=api/proto \
	--go_out ./api/pb --go_opt=paths=source_relative \
	--go-grpc_out ./api/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ./api/pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=cmd/doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=mandoor \
	-I api/proto  api/proto/*.proto

	statik -src=./cmd/doc/swagger -dest=./cmd/doc

run:
	go run cmd/main.go


gen:
	rm -f api/pb/*.go
	rm -f cmd/doc/swagger/*.swagger.json
	buf generate