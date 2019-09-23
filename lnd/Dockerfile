# Build image
FROM golang:1.13

ENV GOPATH=/go \
    NETWORK=mainnet \
    LND_SRC=github.com/lightningnetwork/lnd \
    BTCD_RPC= \
    NODE_NAME=

RUN git clone https://${LND_SRC} ${LND_SRC} \
    && cd $LND_SRC \
    && make \
    && make install

# Final image
FROM alpine:3.10

RUN apk add --no-cache ca-certificates \
        libc6-compat

COPY --from=0 /go/bin/* /bin/

# Default ports (RPC, REST, PEER)
EXPOSE 10009 8080 9735

CMD ["sh", "-c", "sleep 5; lnd --configfile=~/.lnd/lnd-${NETWORK}.conf --bitcoin.${NETWORK} --btcd.rpchost=${BTCD_RPC} --datadir=~/.lnd/${NODE_NAME}"]
