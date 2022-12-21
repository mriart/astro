FROM alpine:3.17
#FROM golang:1.19
#FROM golang:alpine

# Alpine needs some libs to run go executable, so this link. And TZ data not included in alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk --no-cache add tzdata

# ENV RAPID_API_KEY ... //if running from docker
ENV RAPID_API_KEY "d290331181msh2e6a23ba8aff482p1cd1b9jsn491faa927e0b"

COPY . ./

EXPOSE 8080
CMD ["./main"]
