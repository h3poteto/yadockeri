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

FROM ghcr.io/h3poteto/yadockeri-base:202010031549

EXPOSE 9090

USER go

RUN set -ex && mkdir -p /go/src/github.com/h3poteto

ADD --chown=go:go . /go/src/github.com/h3poteto/yadockeri

COPY --from=assets --chown=go:go /var/opt/app/assets /go/src/github.com/h3poteto/yadockeri/assets
COPY --from=assets --chown=go:go /var/opt/app/app/templates/index.html /go/src/github.com/h3poteto/yadockeri/app/templates/index.html

WORKDIR /go/src/github.com/h3poteto/yadockeri

RUN set -ex && \
    go mod download && \
    go generate && \
    go build

ENV ECHO_ENV production

ENTRYPOINT ["/go/src/github.com/h3poteto/yadockeri/docker/entrypoint.sh"]

CMD ["/go/src/github.com/h3poteto/yadockeri/yadockeri"]
