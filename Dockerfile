# buildステージ
FROM golang:1.14.4-alpine3.12 as builder

RUN mkdir /app
WORKDIR /app

 # go modに関するファイルをworkspaceにコピー
COPY go.mod .
COPY go.sum .

RUN go mod download

# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

# buildしたイメージを軽量dockerimageにコピー
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /go/bin/app
CMD /go/bin/app $PORT
