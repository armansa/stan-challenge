FROM alpine:latest

MAINTAINER Arman Sarrafi <arman.zrb@gmail.com>

WORKDIR "/opt"

ADD .docker_build/stan-go-test /opt/bin/stan-go-test
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/stan-go-test"]

