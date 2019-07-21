FROM golang:stretch as builder
WORKDIR src/webserver/
COPY . .
#RUN go build -ldflags="-s -w"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v

FROM scratch
COPY --from=builder go/src/webserver/webserver .

EXPOSE 8080
ENTRYPOINT ["./webserver"]