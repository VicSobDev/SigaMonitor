FROM golang:1.21-alpine AS builder

WORKDIR /src

RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/myapp ./cmd

FROM alpine:latest

COPY --from=builder /bin/myapp /bin/myapp

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["/bin/myapp"]
