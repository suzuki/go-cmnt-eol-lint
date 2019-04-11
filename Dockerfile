FROM golang:1.12 AS build-env
ADD . /work
WORKDIR /work
RUN make

FROM busybox
COPY --from=build-env /work/bin/go-cmnt-eol-lint /usr/local/bin/go-cmnt-eol-lint
WORKDIR /app
ENTRYPOINT ["go-cmnt-eol-lint"]
