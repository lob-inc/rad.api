FROM golang:1.12.4-alpine3.9 AS build
WORKDIR /go/src/github.rakops.com/gatd/rad.api
ARG SSH_PRIVATE_KEY
ENV CGO_ENABLED 0
ENV GO111MODULE on
ARG SSH_PRIVATE_KEY
RUN apk add --no-cache git openssh
RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
  && git config --global --add url."git@github.rakops.com:".insteadOf "https://github.rakops.com/" \
	&& ssh-keyscan github.rakops.com >> ~/.ssh/known_hosts
RUN apk --update upgrade && \
    apk --no-cache add curl git make
COPY . .
WORKDIR /go/src/github.rakops.com/gatd/rad.api/libs/rad
RUN git submodule update
WORKDIR /go/src/github.rakops.com/gatd/rad.api/libs/rad/server
RUN go mod vendor
RUN GO111MODULE=off make bindata
WORKDIR /go/src/github.rakops.com/gatd/rad.api/libs/rad/server/api
RUN make clean
RUN RSSP_API_PACKAGE=/go/src/github.rakops.com/gatd/rad.api/libs/rad/server/api make generate_gohandler
RUN go get -d -v ./...
WORKDIR /go/src/github.rakops.com/gatd/rad.api
RUN go mod vendor
WORKDIR /go/src/github.rakops.com/gatd/rad.api/cmd/api
RUN go build .

FROM alpine:3.9
ARG RSSP_BUILD_ID
RUN mkdir -p /var/log/rssp /scripts
RUN echo $RSSP_BUILD_ID > /etc/rssp-release && mkdir -p /config
RUN apk --update upgrade && \
    apk --no-cache add ca-certificates curl tzdata mysql-client && \
    update-ca-certificates
COPY ./config/example.yml /config/
COPY --from=build /go/src/github.rakops.com/gatd/rad.api/cmd/api/api /go/bin/
ENV PATH /go/bin:$PATH
ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
EXPOSE 8008