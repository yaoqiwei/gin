
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /build
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOPROXY=https://goproxy.cn,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN go build -o gin .

FROM alpine
EXPOSE 20152
ENV TZ=Asia/Shanghai
COPY --from=builder /build/gin /
ENTRYPOINT ["/gin"]