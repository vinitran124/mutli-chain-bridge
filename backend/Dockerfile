FROM golang:1.19-alpine as builder
RUN apk add build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN go build -o content     ./cmd/content/*go
RUN go build -o crawler     ./cmd/crawler/*go
RUN go build -o blockchain  ./cmd/blockchain/*go
RUN go build -o migrate  ./cmd/migrate/*go

FROM alpine:latest
RUN apk add ca-certificates multirun
WORKDIR /app
COPY --from=builder /app/. ./
RUN go run "./migrate up"
EXPOSE 3030
CMD ["multirun",,"./crawler","./blockchain","./content server"]

