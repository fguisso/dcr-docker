# Build image
FROM golang:1.12-alpine3.10

ENV GOPATH=/go \
    DCRD_SRC=github.com/decred/dcrd

RUN apk add --no-cache git\
    && mkdir -p $GOPATH/src && cd $GOPATH/src

RUN git clone https://${DCRD_SRC} ${DCRD_SRC} \
    && cd $DCRD_SRC \
    && go install . ./cmd/..

# Final image
FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=0 /go/bin/* /bin/

# Simnet default port (server, rpc)
EXPOSE 18555 19556

# Testnet default port (server, rpc)
EXPOSE 19108 19109

# Mainnet default port (server, rpc)
EXPOSE 9108 9109

CMD ["dcrd"]
