# Build
FROM golang:1.11 as builder

ENV GOPATH /go
ENV HOMEDIR $GOPATH/src/github.com/marciomansur/bookshelf-api
ENV GLIDE github.com/Masterminds/glide

WORKDIR $HOMEDIR
COPY . ./

RUN go get -v ${GLIDE} && glide install
WORKDIR $HOMEDIR/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api . 

# Run
FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /bin/api .
ENTRYPOINT ["./api"]