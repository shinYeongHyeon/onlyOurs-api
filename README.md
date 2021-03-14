# onlyOurs-api



## Stack
-[x] Docker (Maybe, RDS Docker On Test)  
-[x] ORM: entgo  
-[x] Router Library: [Gorilla](https://github.com/gorilla/mux)  
-[x] Doc: [Swaggo](https://github.com/swaggo/swag)  
-[x] TestFramework: [testify](https://github.com/stretchr/testify) (If Need)


### Test
```shell
go test ./...
```

### Docker
Before Deploy, Schema Update
```shell
go generate ./ent
```
Then,
```shell
docker-compose up -d
```

deprecated
```
docker build --tag golang-docker-tutorial:test .
docker run -p 9999:9999 -v $(pwd):/app golang-docker-tutorial:test
```
