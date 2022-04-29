ARG ARCH=amd64
ARG OS=linux
FROM alpine:3.15.4
RUN apk -u --no-cache add curl ca-certificates
COPY dist/artifacts/api-test /usr/bin/
ENTRYPOINT ["api-test"]
