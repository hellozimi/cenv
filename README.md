# cenv

Cenv is a tool to execute commands with an env-file without having to
export/source the environment variables to your current tty session.


### Usage

```
cenv <env-file> -- <command> [<args>...]
  -v    prints current cenv version
```

#### Example usage

```
cenv .env -- go test ./...
cenv .env.development -- make build
```

The envfiles works with both `export` prefix or not. All lines starting with
`#` will be discarded.

```
# example .env file
POSTGRES_USER=user
POSTGRES_PW=secret
CLIENT_ID=some-id
```

### Install

macos:

```
curl https://github.com/hellozimi/cenv/releases/latest/download/cenv-v0.1.0-darwin-amd64 -o cenv
chmod +x cenv
mv cenv /usr/local/bin
```

linux:

```
wget -O cenv https://github.com/hellozimi/cenv/releases/latest/download/cenv-v0.1.0-darwin-amd64
chmod +x cenv
```

from source:

```
$ go get github.com/hellozimi/cenv
$ go install github.com/hellozimi/cenv
```
