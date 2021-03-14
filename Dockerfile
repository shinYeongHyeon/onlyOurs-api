#Dockerfile
FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex && \
    go get -u github.com/gorilla/mux && \
    go get github.com/labstack/gommon/log

# 9: HotReload
# 10: Gorilla/mux
# 11: log In labstack/gommon

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]
