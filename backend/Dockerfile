FROM golang:1.25 AS builder
WORKDIR /usr/src/housebook
COPY . .
ENV GOCACHE=/root/.cache/go-build
ENV GOMODCACHE=/root/.cache/go-mod
RUN --mount=type=cache,target="/root/.cache/go-build" --mount=type=cache,target="/root/.cache/go-mod" go build -o /usr/bin/housebook

FROM golang:1.25 AS runner
WORKDIR /usr/bin
RUN go install github.com/jackc/tern/v2@v2.3.4
COPY --from=builder /usr/src/housebook /usr/src/housebook
COPY --from=builder /usr/bin/housebook /usr/bin/housebook
EXPOSE 3000
CMD ["/usr/bin/housebook"]
