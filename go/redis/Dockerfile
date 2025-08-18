FROM golang:alpine AS app-builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
  -ldflags="-s -w -extldflags '-static'" \
  -gcflags="-l=4" \
  -a -installsuffix cgo \
  -o /go/bin/app .


FROM alpine:latest AS app-production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=app-builder /go/bin/app .
CMD ["./app"]

FROM golang:alpine AS app-development
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
CMD ["go", "run", "main.go"]
