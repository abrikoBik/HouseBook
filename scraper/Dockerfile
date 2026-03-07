FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN --mount=type=cache,id=backend-cache,target="/root/.cache/go-build" go build -o housebook-scraper
RUN --mount=type=cache,id=backend-cache,target="/root/.cache/go-build" go test ./... -v

FROM golang:1.25 AS runner
WORKDIR /app
COPY --from=builder /app/housebook-scraper .
EXPOSE 3002
CMD ["./housebook-scraper"]
