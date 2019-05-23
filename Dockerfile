# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.
FROM golang:1.11-alpine3.7 as builder

# install tools
RUN apk add --no-cache git

WORKDIR /go/src/openpitrix.io/tag
COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN mkdir -p /openpitrix_bin
RUN go build -v -a -installsuffix cgo -ldflags '-w' -o /openpitrix_bin/tag cmd/tag/main.go



FROM alpine:3.7
# modify pod (container) timezone
RUN apk add -U tzdata && ls /usr/share/zoneinfo && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata

COPY --from=builder /openpitrix_bin/tag /usr/local/bin/
EXPOSE 9201
CMD ["/usr/local/bin/tag"]
