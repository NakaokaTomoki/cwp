FROM golang:1.19-bullseye as builder

ADD . /go/cwp
WORKDIR /go/cwp
RUN make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/cwp/cwp /usr/bin/cwp
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/cwp" ]
