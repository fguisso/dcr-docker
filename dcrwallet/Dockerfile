# Build image
FROM golang:1.12-alpine3.10

ENV DCRWALLET_SRC=github.com/decred/dcrwallet

RUN apk add --no-cache git

RUN git clone https://${DCRWALLET_SRC} /src/${DCRWALLET_SRC} \
    && cd /src/$DCRWALLET_SRC \
    && go install .

COPY . ../config

RUN mkdir "/rpc" \
        && git clone https://github.com/decred/dcrd /src/github.com/decred/dcrd \
        && cd /src/github.com/decred/dcrd \
        && go install -v ./cmd/gencerts ./cmd/dcrctl \
        && gencerts --host="localhost" --directory="/rpc" --force

RUN dcrwallet \ 
        --noinitialload \
        --configfile=/config/dcrwallet-simnet.conf \
        --rpccert=/rpc/rpc.cert \
        --rpckey=/rpc/rpc.key \
        & sleep 5 \
    && cd /config \
    && go run main.go

# Final image
FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=0 /go/bin/* /bin/
COPY --from=0 /data/simnet ../data/simnet
COPY dcrwallet-simnet.conf /root/.dcrwallet/dcrwallet-simnet.conf

RUN mkdir "/rpc"
VOLUME ["/rpc"]

# Simnet default port (server, rpc)
EXPOSE 19556 19557

CMD ["sh", "-c", "dcrwallet -C ~/.dcrwallet/dcrwallet-${NETWORK}.conf"]
