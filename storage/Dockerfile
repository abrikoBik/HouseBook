FROM golang:1.25 AS builder
WORKDIR /usr/src/housebook-storage
COPY . .
ENV GOCACHE=/root/.cache/go-build
ENV GOMODCACHE=/root/.cache/go-mod
RUN --mount=type=cache,target="/root/.cache/go-build" --mount=type=cache,target="/root/.cache/go-mod" go build -o /usr/bin/housebook-storage

FROM golang:1.25 AS runner
WORKDIR /usr/bin
COPY --from=builder /usr/src/housebook-storage /usr/src/housebook-storage
COPY --from=builder /usr/bin/housebook-storage /usr/bin/housebook-storage
EXPOSE 3000
CMD ["/usr/bin/housebook-storage"]
