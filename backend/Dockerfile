FROM golang:1.25 AS builder
WORKDIR /app
RUN go install github.com/jackc/tern/v2@v2.3.4
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN --mount=type=cache,id=backend-cache,target="/root/.cache/go-build" go build -o housebook
RUN sha256sum housebook

FROM debian:trixie-slim AS runner
WORKDIR /app
COPY --from=builder /go/bin/tern /usr/local/bin/
COPY --from=builder /app/housebook .
COPY ./database/migrations ./migrations
EXPOSE 3000
RUN sha256sum housebook
CMD ["./housebook"]
