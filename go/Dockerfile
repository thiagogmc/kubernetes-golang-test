FROM golang:stretch as builder
WORKDIR src/webserver/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -ldflags="-s -w"

FROM scratch
COPY --from=builder go/src/webserver/webserver .

EXPOSE 8000
ENTRYPOINT ["./webserver"]