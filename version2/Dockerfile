FROM golang:1.12-alpine AS build
RUN set -ex; \
    apk update; \
    apk add --no-cache git
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN go build -v -o go-web
FROM alpine:3.4
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# ENV DB db
COPY --from=build /src/go-web /usr/local/bin/go-web
COPY --from=build /src/template.html /usr/local/bin/template.html
RUN chmod +x /usr/local/bin/*
EXPOSE 8080
CMD ["go-web"]
