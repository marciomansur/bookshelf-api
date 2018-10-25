# Build
FROM golang:1.11 as builder

ENV HOMEDIR $GOPATH/src/github.com/marciomansur/bookshelf-api
ENV GLIDE github.com/Masterminds/glide

WORKDIR $HOMEDIR
COPY . ./

RUN go get -v ${GLIDE} && glide install

RUN ["CGO_ENABLED=0", "GOOS=linux", "go build -a -installsuffix nocgo -o /app ."]

# Run
FROM alpine:3.8
RUN apk --no-cache add ca-certificates

COPY --from=builder /app ./
ENTRYPOINT ["./app"]