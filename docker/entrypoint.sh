#!/bin/sh

goose -env production up

exec "$@"
