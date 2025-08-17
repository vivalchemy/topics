FROM golang:alpine AS rabbitmq-producer-builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
  -ldflags="-s -w -extldflags '-static'" \
  -gcflags="-l=4" \
  -a -installsuffix cgo \
  -o /go/bin/app .


FROM alpine:latest AS rabbitmq-producer-production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=rabbitmq-producer-builder /go/bin/app .
EXPOSE 3000
CMD ["./app"]

FROM golang:alpine AS rabbitmq-producer-development
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
CMD ["go", "run", "main.go"]
