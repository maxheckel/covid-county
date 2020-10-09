FROM golang:1.15.0-alpine

MAINTAINER Max Heckel

RUN apk update && apk upgrade
RUN apk add --update bash
RUN apk add --update curl
RUN apk --no-cache add ca-certificates

RUN mkdir /server
ADD ./infra/bin /server/bin

ADD ./infra/docker-run.sh /server/docker-run.sh
RUN chmod +x /server/docker-run.sh

EXPOSE 8000
WORKDIR /server

ENTRYPOINT ["/server/docker-run.sh"]
