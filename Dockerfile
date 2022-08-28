# builder image
FROM golang:1.19.0-alpine3.15 as builder
RUN mkdir /build
ADD src /build/
WORKDIR /build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a cmd/main.go


# generate clean, final image for end users
FROM alpine:3.15
COPY --from=builder /build/main ./generator

# executable
ENTRYPOINT [ "./generator -concurrent 10 -sleepus 500 -msg 'This is a sample log message to simulate application log from Snocko applications'" ]