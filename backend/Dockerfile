FROM golang:1.25 AS builder
WORKDIR /usr/src/housebook
COPY . .
RUN go mod download
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o /usr/bin/housebook

# TODO: replace with scratch container in production
FROM golang:1.25 AS runner
WORKDIR /usr/bin
RUN go install github.com/jackc/tern/v2@v2.3.4
COPY --from=builder /usr/src/housebook /usr/src/housebook
COPY --from=builder /usr/bin/housebook /usr/bin/housebook
EXPOSE 3000
CMD ["/usr/bin/housebook"]
