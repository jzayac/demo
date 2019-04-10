FROM golang:alpine as builder
COPY ./ /go/src/demo
WORKDIR /go/src/demo
RUN apk add curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=builder /go/src/demo/main .
EXPOSE 7000
CMD ["./main"]
