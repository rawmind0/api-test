FROM alpine:3.15.4
RUN apk -u --no-cache add curl ca-certificates
COPY bin/api-test /usr/bin/
CMD ["api-test"]