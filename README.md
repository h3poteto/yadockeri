# yadockeri
[![CircleCI](https://circleci.com/gh/h3poteto/yadockeri.svg?style=svg)](https://circleci.com/gh/h3poteto/yadockeri)
[![Docker Pulls](https://img.shields.io/docker/pulls/h3poteto/yadockeri)](https://hub.docker.com/r/h3poteto/yadockeri)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/h3poteto/yadockeri)](https://github.com/h3poteto/yadockeri/releases)

Yadockeri is a web application to install helm charts in your kubernetes cluster.
**The aim of Yadockeri is to deploy your web application as a separate environment for each branch of GitHub.**


## Introduction
Usually, you are deploying your application in production, staging and development, aren't you?
But do you want to deploy the branch which you are developing to a dedicated environment? In this case, Yadockeri can help you.

At first, you have to build Docker image which contains your application, and push it like `docker_image_name:commit_hash`.

Yadockeri is using helm chart to deploy your application, so you have to create helm chart for your application.

After Yadocker start, you can specify your application repository and the helm chart on Yadockeri.
And you can specify some values to override your helm chart. This values are provided as `--set` option when helm install.

Finally, you can specify a branch in your application repository, and start deploy. In deploy, Yadockeri find specified GitHub branch and get latest commit hash from it, and override docker image tag in your helm chart. After that Yadockeri runs helm install, and you can confirm helm logs on Yadockeri.

If you want to know more information, please read examples.

## Install
There is a docker image, so please pull it.

```
$ docker pulll h3poteto/yadockeri:latest
```

Yadockeri uses PostgreSQL as detabase, so please prepare PostgreSQL and create `yadockeri` database.

```sql
> create database yadockeri;
```

Yadockeri needs a GitHub Application for OAuth login, so please create OAuth Application in GitHub.

And start Yadockeri:

```
$ docker run --rm --service-ports \
  -v $HOME/.kube:/root/.kube \
  -e CLIENT_ID=github_application_id \
  -e CLIENT_SECRET=github_application_secret \
  -e ORGANIZATION=github_organization \
  -e POSTGRES_HOST=database_hostname \
  -e POSTGRES_USER=database_username \
  -e POSTGRES_PASSWORD=database_password \
  -e KUBECONFIG=/root/.kube/config
```

Yadockeri uses helm command, so KUBECONFIG is required.

## Usage
TODO.

## Example
TODO.

## Development
Please read [development guid](DEVELOPMENT.md).

# License
The software is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

