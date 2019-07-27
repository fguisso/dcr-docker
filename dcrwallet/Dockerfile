# Build image
FROM golang:1.12-alpine3.10

ENV DCRWALLET_SRC=github.com/decred/dcrwallet

RUN apk add --no-cache git

RUN git clone https://${DCRWALLET_SRC} /src/${DCRWALLET_SRC} \
    && cd /src/$DCRWALLET_SRC \
    && go install .

COPY . ../config

RUN dcrwallet --simnet --noinitialload & sleep 5 \
    && cd ../config \
    && go run main.go

# Final image
FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=0 /go/bin/* /bin/
COPY --from=0 /root/.dcrwallet/simnet ../root/data/simnet

# Simnet default port (server, rpc)
EXPOSE 19556 19557

CMD ["dcrwallet", "--configfile=/root/.dcrwallet/dcrwallet.conf"]
