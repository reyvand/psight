FROM golang:1.15.3-alpine3.12 AS builder

# Install OS level dependencies
RUN apk add --update alpine-sdk git && \
	git config --global http.https://gopkg.in.followRedirects true

WORKDIR /go/src/github.com/reyvand/psight/
COPY ./ .

RUN go build -o psight

FROM alpine:latest
WORKDIR /go/src/github.com/reyvand/psight/
COPY --from=builder /go/src/github.com/reyvand/psight /go/src/github.com/reyvand/psight
COPY --from=builder /go/src/github.com/reyvand/psight /bin/psight

ENTRYPOINT ["/bin/psight"]
