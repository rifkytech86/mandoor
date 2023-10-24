TODO 
### Mandoor System Exploration

- [x] Generate GRPC 
- [x] Generate GRPC Gateway
- [x] Generate GRPC validate
- [x] Generate migrate
- [x] Generate Ecosystem GRPC using Buf
- [x] Up Swagger Static
- [ ] Logger
- [ ] Docker
- [ ] KONG API Gateway
- [ ] Implement Consul




# Note set golang migrate
```
go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#create table 
migrate create -ext sql -dir db/migration create_table_user

# after input scema up and down 
migrate -database 'mysql://root:ZXCasdqwe123!@tcp(localhost:3306)/mandoor_db' -path db/migration up
migrate -database 'mysql://root:ZXCasdqwe123!@tcp(localhost:3306)/mandoor_db' -path db/migration up

```



