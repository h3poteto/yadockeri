## Development
Development guid for Yadockeri in your local machine.

### Server Side

At first, you have to register GitHub OAuth Application.
And you please export environment variables:

```bash
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
/go/src/github.com/h3poteto/yadockeri # go generate // generate asset
/go/src/github.com/h3poteto/yadockeri # go run main.go
```

### Frontend

```cmd
$ docker-compose run --rm frontend
```

After that, you can access `http://localhost:9090`.
