#Dockerfile
FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]
