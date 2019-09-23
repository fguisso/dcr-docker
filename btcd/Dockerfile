# Build image
FROM golang:1.12-alpine3.10

ENV GOPATH=/go \
    NETWORK=mainnet \
    BTCD_SRC=github.com/btcsuite/btcd

RUN apk add --no-cache git\
    && mkdir -p $GOPATH/src

RUN git clone https://${BTCD_SRC} ${BTCD_SRC} \
    && cd $BTCD_SRC \
    && go install . ./cmd/..

# Final image
FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=0 /go/bin/* /bin/

# Simnet default port (server, rpc)
EXPOSE 18555 18556

# Testnet default port (server, rpc)
EXPOSE 18333 18334

# Mainnet default port (server, rpc)
EXPOSE 8333 8334

CMD ["sh", "-c", "btcd -C ~/.btcd/btcd-${NETWORK}.conf"]
