# Build image
FROM golang:1.12-alpine3.10

ENV GOPATH=/go \
    NETWORK=mainnet \
    DCRD_RPC= \
    NODE_NAME=

ARG DCRLND_SRC

RUN apk add --no-cache git make\
    && mkdir -p $GOPATH/src && cd $GOPATH/src

ARG DCRLND_VERSION
RUN git clone https://${DCRLND_SRC} ${DCRLND_SRC} \
    && cd $DCRLND_SRC \
    && if [ -n "$DCRLND_VERSION" ]; then git checkout ${DCRLND_VERSION}; fi \
    && make install

# Final image
FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=0 /go/bin/* /bin/

# Default ports (RPC, REST, PEER)
EXPOSE 10009 8080 9735

CMD ["sh", "-c", "sleep 5; dcrlnd --configfile=~/.dcrlnd/dcrlnd-${NETWORK}.conf --decred.${NETWORK} --dcrd.rpchost=${DCRD_RPC} --lnddir=~/.dcrlnd/${NODE_NAME}"]
