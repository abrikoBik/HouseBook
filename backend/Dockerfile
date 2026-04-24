FROM golang:1.25 AS builder
WORKDIR /app
RUN go install github.com/jackc/tern/v2@latest && go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN /go/bin/sqlc -f database/sqlc.yml vet && /go/bin/sqlc -f database/sqlc.yml compile && /go/bin/sqlc -f database/sqlc.yml diff
RUN --mount=type=cache,id=backend-cache,target="/root/.cache/go-build" go build -o housebook
RUN --mount=type=cache,id=backend-cache,target="/root/.cache/go-build" go test ./... -v

FROM debian:trixie-slim AS runner
WORKDIR /app
COPY --from=builder /app/docs/swagger.json .
COPY --from=builder /go/bin/tern /usr/local/bin/
COPY --from=builder /app/housebook .
COPY ./database/migrations ./migrations
EXPOSE 3000
CMD ["./housebook"]
