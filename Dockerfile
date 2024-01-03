# Build stage
FROM golang:1.21-alpine AS builder
RUN apk --no-cache add gcc musl-dev
WORKDIR /app
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

#ENV GOPROXY=direct
RUN go mod download
RUN CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o ./ ./app/...

CMD ["./app"]

FROM golang:1.21-alpine

WORKDIR /app

# ENV DB_PATH=/sqlite/lensman.db
ENV CONFIG_PATH=/.goimg/
COPY --from=builder /app/cmd .

CMD ["./cmd"]