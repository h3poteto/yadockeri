FROM node:12.7.0-alpine AS assets

WORKDIR /var/opt/app

COPY ./package.json /var/opt/app/
COPY ./yarn.lock /var/opt/app/

RUN set -ex && \
    yarn install

COPY .browserslistrc /var/opt/app/
COPY .eslintrc.js /var/opt/app/
COPY .prettierrc /var/opt/app/
COPY ./vue.config.js /var/opt/app/
COPY ./tsconfig.json /var/opt/app/

COPY ./public /var/opt/app/public
COPY ./frontend /var/opt/app/frontend

RUN set -ex && \
    yarn run build && \
    rm -rf node_modules

FROM yadockeri-base:latest

EXPOSE 9090

ENV GO111MODULE on

WORKDIR /go/src/github.com/h3poteto/yadockeri

USER root

ADD . /go/src/github.com/h3poteto/yadockeri

COPY --from=assets /var/opt/app/assets /go/src/github.com/h3poteto/yadockeri/assets
COPY --from=assets /var/opt/app/app/templates/index.html /go/src/github.com/h3poteto/yadockeri/app/templates/index.html

RUN set -ex && \
    apk add --no-cache gettext

RUN set -ex && \
    go mod download && \
    go generate && \
    go build && \
    rm -rf /go/pkg/mod && \
    rm -rf /root/.cache && \
    chown go:go -R /go

USER go

CMD ["/go/src/github.com/h3poteto/yadockeri/yadockeri"]
