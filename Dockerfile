FROM alpine:3.7

COPY ./go-redis-socket /go-redis-socket
RUN apk update
RUN apk --no-cache add \
    tzdata \
    su-exec \
    ca-certificates \
    s6 \
    curl \
    openssh \
    make 
EXPOSE 6969
ENTRYPOINT ["/go-redis-socket"]