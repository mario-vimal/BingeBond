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
go get -u -v github.com/jinzhu/gorm
go get -u -v github.com/graph-gophers/graphql-go
```

Setting up the database
 * Create a mysql database
 * Create a new file called config.json in the config directory
 * Copy the contents of config.json.example and paste it in the newly created config.json file
 * Set the correct db name, user and password in config.json file
`

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

# Contributing

Read our [Contribution Guidelines](https://github.com/mario-vimal/BingeBond/blob/master/CONTRIBUTING.md) for information on how you can help out BingeBond.