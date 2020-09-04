FROM golang:1.13-alpine AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-summer-vacation
COPY ./ ./
RUN go build -o server pkg/server/server.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-summer-vacation/server /usr/local/bin/server
ENV PORT 8080

ENV User miraiketai2020
ENV Pass miraiketai2020
ENV IP db
ENV Port 3306
ENV Name summer

EXPOSE 8080
CMD ["/usr/local/bin/server"]