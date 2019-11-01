# BingeBond

BingeBond is an Android application that allows users to connect on multiple networking platforms instantly.

## Backend

### Tech

* [Go]
* [Gin]
* [GraphQL]

### Prerequisites

- BingeBond Go [Go](https://golang.org/) v1.10 and above to run.

### Getting Started

Install the dependencies

```sh
go get -u -v github.com/gin-gonic/gin
go get -u -v github.com/gin-contrib/cors
go get -u -v github.com/gin-gonic/contrib/static
go get -u -v github.com/Unknwon/bra
go get -u -v github.com/fatih/color
```

For development environments...

```sh
cd backend
$ npm install
$ npm run watch
```

For production environments...

```sh
cd backend
$ npm install
$ npm run prod
```

Before committing your changes, format your codebase using...

```sh
$ npm run fmt
```

If project directory is outside $GOPATH, add these lines in ~/.bashrc

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

[Go]: <https://golang.org>
[Gin]: <https://gin-gonic.github.io/gin>
[GraphQL]: <https://graphql.org/learn>

## Reporting Issues

If you think you've found a bug, or something isn't behaving the way you think it should, please raise an [issue](https://github.com/mario-vimal/BingeBond/issues) on GitHub.

