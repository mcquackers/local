FROM golang:1.9

LABEL maintainer="Brian McQueen <brian.d.mcqueen@gmail.com>"
ARG VERSION
ARG BUILD_TIME

LABEL VERSION=$VERSION
ENV VERSION $VERSION
ENV BUILD_TIME ${BUILD_TIME}

COPY build/$VERSION/local /opt/mcquackers/local
RUN chmod +x /opt/mcquackers/local

EXPOSE 8080

ENTRYPOINT ["/opt/mcquackers/local", "-alsologtostderr"]
