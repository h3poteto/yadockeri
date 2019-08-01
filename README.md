# yadockeri

## Development
### Server Side

At first, you have to export environment variables:

```bash
export AWS_ACCESS_KEY_ID=hogehoge
export AWS_SECRET_ACCESS_KEY=fugafuga
export GITHUB_CLIENT_ID=hoge
export GITHUB_CLIENT_SECRET=fuga
export SESSION_SECRET=your_secret
export ALLOW_GITHUB_ORG=your_github_org_name
```

And run server.

```cmd
$ docker-compose run --rm --service-ports app sh

/go/src/github.com/h3poteto/yadockeri # goose up
...
/go/src/github.com/h3poteto/yadockeri # glide install -v
/go/src/github.com/h3poteto/yadockeri # go run main.go
```

### Frontend

```cmd
$ docker-compose run --rm frontend
```
