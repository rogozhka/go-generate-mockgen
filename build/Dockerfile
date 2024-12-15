FROM golang:1.23.4-alpine3.21 AS builder

# ca-certificates   - for https connections
# git               - to download go packages
RUN apk add --no-cache git ca-certificates

RUN go install go.uber.org/mock/mockgen@v0.5.0

FROM golang:1.23.4-alpine3.20

RUN apk add --no-cache bash

COPY --from=builder /go/bin/mockgen /go/bin/mockgen
COPY ./scripts/wrap-mockgen.sh /bin/wrap-mockgen.sh
RUN chmod +x /bin/wrap-mockgen.sh
WORKDIR /w
ENTRYPOINT ["/bin/wrap-mockgen.sh"]
