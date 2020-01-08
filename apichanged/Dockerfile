
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/senura96/apichanged
COPY . .

RUN go get -d -v

RUN go build -o /bin/apichanged

FROM scratch  

COPY --from=builder /bin/apichanged /bin/apichanged

ENTRYPOINT ["/bin/go-gin-app"]