FROM alpine:3.17

# Alpine needs some libs to run go executable. So this link
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY . ./

EXPOSE 8080
CMD ["./main"]
