FROM golang:alpine AS intermediate

RUN apk update && \
    apk add --no-cache git make

RUN adduser -D -g '' corridgen

WORKDIR $GOPATH/src/

COPY . .

RUN go mod download
RUN go mod verify
RUN make build

FROM scratch

ENV PORT=7891

COPY --from=intermediate /go/src/bin/corridgen /go/bin/corridgen
COPY --from=intermediate /etc/passwd /etc/passwd

USER corridgen

WORKDIR /go/bin

CMD ["/go/bin/corridgen"]