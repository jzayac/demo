FROM golang:alpine as builder
COPY ./ /go/src/demo
WORKDIR /go/src/demo
RUN CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=builder /go/src/demo/main .
EXPOSE 7000
CMD ["./main"]
