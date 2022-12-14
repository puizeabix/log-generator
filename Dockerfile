# builder image
FROM golang:1.19.0-alpine3.15 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a cmd/main.go


# generate clean, final image for end users
FROM alpine:3.15
COPY --from=builder /build/main ./generator

# executable
ENTRYPOINT [ "./generator"]