FROM golang:1.11.0
WORKDIR /go/src/gowebdemo/
#RUN go get -d -v golang.org/x/net/html
COPY jenkins/app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gowebdemo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/gowebdemo .
EXPOSE 8088
CMD ["./gowebdemo"]