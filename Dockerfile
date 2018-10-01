
FROM golang:alpine

ADD ./airvisual.go /go/src/app
WORKDIR /go/src/app
ENV PORT=5000
CMD ["go", "run", "main.go"]
