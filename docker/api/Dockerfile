FROM golang:1.16 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go get -v ./...
COPY . .
RUN go build -o metrogalaxy-api

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/metrogalaxy-api /
COPY --from=builder /app/env/*.json /env/
EXPOSE 9000
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/metrogalaxy-api", "serve"]