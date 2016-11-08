FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/pes /opt/bin/pes
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/pes"]
