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
There is a helm chart to install yadockeri.

```
$ helm repo add h3poteto-stable https://h3poteto.github.io/charts/stable
$ helm install h3poteto-stable/yadockeri --name yadockeri --set config.github.organization="your_organization_name" --set config.github.client_id="id_of_your_application" --set config.github.client_secret="secret_of_your_application"
```

Please see [chart document](https://github.com/h3poteto/charts/tree/master/stable/yadockeri) for more information.

## Usage
TODO.

## Example
TODO.

## Development
Please read [development guid](DEVELOPMENT.md).

# License
The software is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

