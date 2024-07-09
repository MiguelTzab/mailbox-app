FROM golang:1.22.2-alpine AS builder

WORKDIR /build

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -v -o serve ./cmd/api/main.go

FROM scratch

COPY --from=builder /build/serve /api

EXPOSE 3000

CMD ["./api"]