#
# Dockerfile for 5.9.1
#

FROM bellsoft/liberica-openjdk-alpine:13

ENV ACTIVEMQ_VERSION 5.9.1
ENV ACTIVEMQ apache-activemq-$ACTIVEMQ_VERSION

ENV ACTIVEMQ_HOME /opt/activemq

ARG VERSION=5.9.1

LABEL maintainer="Team Hadenlabs <hello@hadenlabs.com>" \
      org.label-schema.vcs-url="https://github.com/hadenlabs/docker-activemq" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

ENV RUNTIME_DEPS \
    curl \
    git \
    gzip \
    bash \
    gcc

RUN \
    apk add --update --no-cache $RUNTIME_DEPS \
    && mkdir -p /opt \
    && curl -s -S https://archive.apache.org/dist/activemq/$ACTIVEMQ_VERSION/$ACTIVEMQ-bin.tar.gz | tar -xvz -C /opt \
    && ln -s /opt/$ACTIVEMQ $ACTIVEMQ_HOME \
    && rm -rf /var/cache/apk/*

EXPOSE 1883 5672 8161 61613 61614 61616

# Expose some folders
VOLUME ["/opt/activemq/data"]
VOLUME ["/var/log/activemq"]
VOLUME ["/opt/activemq/conf"]

WORKDIR $ACTIVEMQ_HOME

CMD ["/bin/bash", "-c", "bin/activemq console"]