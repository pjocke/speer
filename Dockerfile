FROM golang:alpine

ARG pkg=github.com/pjocke/speer

RUN apk add --no-cache ca-certificates

COPY . $GOPATH/src/$pkg

RUN set -ex \
      && apk add --no-cache --virtual .build-deps \
              git \
      && go get -v $pkg \
      && apk del .build-deps

RUN go install $pkg

WORKDIR $GOPATH/src/$pkg

CMD ["speer"]
