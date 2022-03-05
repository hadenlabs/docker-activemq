#
# Dockerfile for 0.0.0
#

FROM ubuntu:20.04

ARG VERSION=0.0.0

LABEL maintainer="Team Hadenlabs <hello@hadenlabs.com>" \
      org.label-schema.vcs-url="https://github.com/hadenlabs/docker-activemq" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

ENV LIBRARY_DEPS \
  libproj-dev

ENV RUNTIME_DEPS \
    curl \
    git \
    gcc

RUN apt-get update -y \
  && apt-get install -y --no-install-recommends $LIBRARY_DEPS $RUNTIME_DEPS \
  && apt-get clean \
  && apt-get purge -y \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /data

ENTRYPOINT ["curl"]

CMD ["--help"]
