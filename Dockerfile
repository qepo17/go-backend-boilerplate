FROM golang:1.22-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o web ./cmd/web/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /app/web /app/web

EXPOSE 80

CMD ["/app/web"]