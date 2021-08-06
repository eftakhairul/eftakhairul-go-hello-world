ARG GO_VERSION=1.16.5
FROM golang:${GO_VERSION} as builder


WORKDIR /app
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor main.go


FROM scratch as production

COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["./main"]
CMD ""
