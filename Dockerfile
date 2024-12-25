FROM golang:bullseye as builder

WORKDIR /app

COPY . .

RUN ["go", "build"]

CMD "ls"

FROM golang:bullseye AS runner

WORKDIR /app

COPY --from=builder /app/backend .

CMD ["/app/backend"]
