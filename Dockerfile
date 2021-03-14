#Dockerfile
FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex && \
    go get -u github.com/gorilla/mux

# 9: HotReload
# 10: Gorilla/mux

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]
