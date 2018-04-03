FROM golang:alpine as builder
WORKDIR /go/src/github.com/kasuboski/test-go-api/
RUN apk add --no-cache git
RUN go get -d -v github.com/go-chi/chi \
	&& go get -d -v github.com/go-chi/render  
COPY *.go    .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/kasuboski/test-go-api/app .
ENV PORT 8080
EXPOSE $PORT 
CMD ["./app"]
