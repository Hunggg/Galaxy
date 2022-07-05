
FROM golang:alpine as builder

RUN apk --no-cache add git curl

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    mkdir -p $GOPATH/src/github.com/improbable-eng/grpc-web && \
    cd $GOPATH/src/github.com/improbable-eng && \
    git clone https://github.com/improbable-eng/grpc-web && \
    cd grpc-web && \
    go get -v ./... && go install ./go/grpcwebproxy

# RUN go get -u github.com/improbable-eng/grpc-web/go/grpcwebproxy

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /go/bin/grpcwebproxy .
RUN mkdir certs
COPY server-cert.pem /certs/server-cert.pem 
COPY server-key.pem /certs/server-key.pem

EXPOSE 8081 8443

CMD ["/grpcwebproxy"]