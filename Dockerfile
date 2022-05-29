FROM golang:alpine as builder
RUN apk add make
ADD . /image-api
RUN cd /image-api && make build

FROM alpine:3.14
COPY --from=builder /image-api/bin/* /usr/local/bin/
COPY --from=builder /image-api/config/config-local.yml ./config/config-local.yml
EXPOSE 5000
