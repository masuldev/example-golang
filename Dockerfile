FROM golang:1.16-alpine as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /usr/src/app
COPY . .

RUN go mod download

RUN go build -o main .

FROM scratch
COPY --from=builder /usr/src/app /
CMD ["/main"]