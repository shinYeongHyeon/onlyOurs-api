# onlyOurs-api



## Stack
-[x] Docker (Maybe, RDS Docker On Test)  
-[ ] ORM: entgo  
-[x] Router Library: [Gorilla](https://github.com/gorilla/mux)  
-[x] Doc: [Swaggo](https://github.com/swaggo/swag)  
-[ ] TestFramework: [testify](https://github.com/stretchr/testify) (If Need)


### Docker
```shell
docker-compose up -d
```

deprecated
```
docker build --tag golang-docker-tutorial:test .
docker run -p 9999:9999 -v $(pwd):/app golang-docker-tutorial:test
```
