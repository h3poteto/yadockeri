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

FROM h3poteto/yadockeri-base:20190801

EXPOSE 9090

WORKDIR /go/src/github.com/h3poteto/yadockeri

ADD . /go/src/github.com/h3poteto/yadockeri

RUN set -ex && \
    apk add --no-cache gettext && \
    glide install -v && \
    go generate && \
    go build && \
    rm -rf vendor && \
    rm -rf /root/.glide && \
    rm -rf /root/.cache

COPY --from=assets /var/opt/app/assets /go/src/github.com/h3poteto/yadockeri/assets
COPY --from=assets /var/opt/app/app/templates/index.html /go/src/github.com/h3poteto/yadockeri/app/templates/index.html

CMD ["/go/src/github.com/h3poteto/yadockeri/yadockeri"]
