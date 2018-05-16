ARG DEP_VER=latest
FROM local/dependency:$DEP_VER

ARG VERSION
ENV VERSION $VERSION

ARG OUT_DIR=/go/out
ENV OUT_DIR $OUT_DIR
ENV WORKDIR $GOPATH/src/github.com/mcquackers/local

LABEL maintainer="Brian McQueen <brian.d.mcqueen@gmail.com>"
LABEL version=$VERSION

WORKDIR ${WORKDIR}
COPY . $WORKDIR

RUN mkdir -p $OUT_DIR
VOLUME $OUT_DIR

RUN go install

CMD ["cp", "/go/bin/local", "/go/out"]
