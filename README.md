# onlyOurs-api



## Stack
-[x] Docker (Maybe, RDS Docker On Test)  
-[ ] ORM: entgo  
-[x] Framework: [Gorilla](https://github.com/gorilla/mux)  
-[ ] Doc: [Swaggo](https://github.com/swaggo/swag)  
-[ ] TestFramework: [testify](https://github.com/stretchr/testify) (If Need)


### Docker
```
docker build --tag golang-docker-tutorial:test .
docker run -p 9999:9999 -v $(pwd):/app golang-docker-tutorial:test
```
