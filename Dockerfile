FROM golang:1.17.0-alpine3.13 AS build-env
RUN apk --no-cache add build-base git mercurial gcc openssh-client ca-certificates tzdata
RUN addgroup -g 32000 grp && adduser -D -H -u 32001 -G grp appuser
ADD . /src

RUN cd /src && make production_bin

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-env /src/signer /app/
COPY --from=build-env /etc/passwd /etc/passwd
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-env /usr/share/zoneinfo /usr/share/zoneinfo
USER appuser

EXPOSE 8080
ENTRYPOINT ["./signer"]
