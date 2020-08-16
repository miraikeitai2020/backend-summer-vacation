FROM golang:1.13-alpine AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-summer-vacation
COPY ./ ./
RUN go build -o server pkg/server/server.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-summer-vacation/server /usr/local/bin/server
ENV PORT 8080

EXPOSE 8080
CMD ["/usr/local/bin/server"]