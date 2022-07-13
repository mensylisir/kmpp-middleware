FROM golang:1.17 as stage-build
LABEL stage=stage-build
WORKDIR /build/middleware
ARG GOPROXY
ARG GOARCH


ENV GOARCH=$GOARCH
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV GOOS=linux
ENV CGO_ENABLED=1

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's|security.debian.org/debian-security|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list

RUN apt-get update && apt-get install unzip
COPY . .
RUN go mod tidy

RUN wget https://github.com/go-bindata/go-bindata/archive/v3.1.3.zip -O /tmp/go-bindata.zip  \
    && cd /tmp \
    && unzip  /tmp/go-bindata.zip  \
    && cd /tmp/go-bindata-3.1.3 \
    && go build \
    && cd go-bindata \
    && go build \
    && cp go-bindata /go/bin

RUN export PATH=$PATH:$GOPATH/bin
RUN make build_server_linux GOARCH=$GOARCH

FROM ubuntu:20.04
ARG GOARCH
WORKDIR /

COPY --from=stage-build /build/middleware/dist/etc /etc/
COPY --from=stage-build /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip

COPY --from=stage-build /build/middleware/dist/etc /etc/
COPY --from=stage-build /build/middleware/dist/usr /usr/

EXPOSE 8080

CMD ["kmpp-middleware", "server"]
