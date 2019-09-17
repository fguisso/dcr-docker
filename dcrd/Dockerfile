# Build image
FROM golang:1.12-alpine3.10

ENV GOPATH=/go \
    DCRD_SRC=github.com/decred/dcrd

RUN apk add --no-cache git\
    && mkdir -p $GOPATH/src && cd $GOPATH/src

ARG DCRD_VERSION
RUN git clone https://${DCRD_SRC} ${DCRD_SRC} \
    && cd $DCRD_SRC \
    && if [ -n "$DCRD_VERSION" ]; then git checkout ${DCRD_VERSION}; fi \
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

ENV DCRD_NETWORK=mainnet
CMD ["sh", "-c", "dcrd -C ~/.dcrd/dcrd-${DCRD_NETWORK}.conf"]
