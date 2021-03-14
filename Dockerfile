#Dockerfile
FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex && \
    go get -u github.com/gorilla/mux && \
    go get github.com/labstack/gommon/log && \
    go get -u github.com/swaggo/swag/cmd/swag && \
    go get -u github.com/swaggo/http-swagger && \
    go get -u github.com/alecthomas/template && \
    go get entgo.io/ent/cmd/ent && \
    go get github.com/lib/pq && \
    go get github.com/stretchr/testify

# 9: HotReload
# 10: Gorilla/mux
# 11: log In labstack/gommon
# 12~14: Swaggo Setting (https://www.soberkoder.com/swagger-go-api-swaggo/)

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]
